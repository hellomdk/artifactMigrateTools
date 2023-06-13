package migrate

import (
	"artifactMigrateTools/internal/api/nexus"
	"artifactMigrateTools/internal/common"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
	"artifactMigrateTools/internal/util"
	"fmt"
	"net/http"
)

type NexusRepoMigrate struct {
	Nexus nexus.Nexus
}

func NewNexusRepoMigrate(config *model.Config) NexusRepoMigrate {
	nx := &nexus.Nexus{
		util.HttpClient{
			BaseURL:  config.SourceRepo.URL,
			Username: config.SourceRepo.Username,
			Password: config.SourceRepo.Password,
			Header:   http.Header{},
		},
	}
	util.Auth(nx.HttpClient)

	instance := new(NexusRepoMigrate)
	instance.Nexus = *nx
	return *instance
}

func (orm NexusRepoMigrate) VerifyConfig(context *config.Context) bool {
	body, _ := orm.Nexus.Ping(context)
	if body != "" {
		context.Loggers.SendLoggerInfo(fmt.Sprint("验证连接成功: ", orm.Nexus.HttpClient.BaseURL))
		return true
	}
	return false
}

// 获取仓库列表
func (orm NexusRepoMigrate) GetNexusRepositories(context *config.Context) []model.Repositories {
	nexusData := orm.Nexus.GetNexusData(context, "artifactorymigrator")
	var repoList []model.Repositories
	for _, item := range nexusData.Repos {
		var repoItem model.Repositories
		repoItem.RepoKey = item.Name
		repoItem.RepoKeyMapping = item.Name
		protocolType, layout := common.GetPro(item.Format)
		repoItem.ProtocolType = protocolType
		repoItem.Layout = layout

		if item.Type == "hosted" {
			repoItem.RepoType = "local"
		} else if item.Type == "proxy" {
			repoItem.RepoType = "remote"
		} else {
			repoItem.RepoType = "virtual"
			var selectedRepoList []model.SelectedRepositories
			group := item.Config.Attributes.Group
			for _, repoKey := range group.MemberNames {
				var selectedRepo model.SelectedRepositories
				selectedRepo.RepoKey = repoKey
				selectedRepoList = append(selectedRepoList, selectedRepo)
			}
			repoItem.SelectedRepositories = selectedRepoList
		}
		repoItem.Url = item.Config.Attributes.Proxy.RemoteUrl
		repoItem.Username = item.Config.Attributes.Httpclient.Authentication.Username
		repoItem.Password = item.Config.Attributes.Httpclient.Authentication.Password
		repoItem.ProjectKey = item.Config.Attributes.Storage.BlobStoreName
		repoItem.Migrated = false
		repoList = append(repoList, repoItem)
	}
	return repoList
}

func (orm NexusRepoMigrate) GetNexusArtifact(context *config.Context, repoKey string) []model.Arti {
	fileItemList := orm.Nexus.GetItemList(context, repoKey)
	var artiList []model.Arti
	for _, item := range fileItemList.Items {
		for _, fileItem := range item.Assets {
			var artiItem model.Arti
			artiItem.Name = fileItem.Path
			artiItem.Repo = fileItem.Repository
			artiItem.RepoMapping = fileItem.Repository
			protocolType, _ := common.GetPro(item.Format)
			artiItem.ProtocolType = protocolType
			artiItem.Path = fileItem.Path
			artiItem.Sha1 = fileItem.Checksum.Sha1
			artiItem.Sha256 = fileItem.Checksum.Sha256
			artiItem.Md5 = fileItem.Checksum.Md5
			artiList = append(artiList, artiItem)
		}
	}
	return artiList
}
