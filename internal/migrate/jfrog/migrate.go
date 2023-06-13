package migrate

import (
	"artifactMigrateTools/internal/api/jfrog"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
	"artifactMigrateTools/internal/util"
	"fmt"
	"net/http"
	"strings"
)

type JfrogMigrate struct {
	Jfrog jfrog.Jfrog
}

func NewJfrogMigrate(config *model.Config) JfrogMigrate {
	jf := &jfrog.Jfrog{
		util.HttpClient{
			BaseURL:  config.SourceRepo.URL,
			Username: config.SourceRepo.Username,
			Password: config.SourceRepo.Password,
			Header:   http.Header{},
		},
	}
	util.Auth(jf.HttpClient)

	instance := new(JfrogMigrate)
	instance.Jfrog = *jf
	return *instance
}

func (jm JfrogMigrate) VerifyConfig(context *config.Context) bool {
	body, _ := jm.Jfrog.Ping(context)
	if body != "" {
		context.Loggers.SendLoggerInfo(fmt.Sprint("验证连接成功: ", jm.Jfrog.HttpClient.BaseURL))
		return true
	}
	return false
}

// 获取仓库列表
func (jm JfrogMigrate) GetJfrogRepositories(context *config.Context) []model.Repositories {
	jfrogData := jm.Jfrog.GetRepositories(context)

	var repoList []model.Repositories
	for _, item := range jfrogData {
		var repoItem model.Repositories
		repositoryDetails := jm.Jfrog.GetRepository(context, item.Key)

		repoItem.RepoKey = item.Key
		repoItem.RepoKeyMapping = item.Key
		repoItem.ProtocolType = repositoryDetails.PackageType
		repoItem.Layout = repositoryDetails.RepoLayoutRef
		repoItem.RepoType = strings.ToLower(item.Type)

		if repoItem.RepoType == "virtual" {
			var selectedRepoList []model.SelectedRepositories
			group := repositoryDetails.Repositories
			for _, repoKey := range group {
				var selectedRepo model.SelectedRepositories
				selectedRepo.RepoKey = repoKey
				selectedRepoList = append(selectedRepoList, selectedRepo)
			}
			repoItem.SelectedRepositories = selectedRepoList
		}
		repoItem.DefaultDeploymentRepo = repositoryDetails.DefaultDeploymentRepo
		repoItem.Url = repositoryDetails.Url
		repoItem.Username = repositoryDetails.Username
		repoItem.Password = repositoryDetails.Password
		repoItem.ProjectKey = repositoryDetails.ProjectKey
		repoItem.Migrated = false
		repoList = append(repoList, repoItem)
	}
	return repoList
}

func (jm JfrogMigrate) GetJfrogRepositoriesByProjectKey(context *config.Context, projectKey string) []model.Repositories {
	allJfrogRepositories := jm.GetJfrogRepositories(context)
	if projectKey == "" || projectKey == "all" {
		return allJfrogRepositories
	}

	var repoList []model.Repositories
	for _, item := range allJfrogRepositories {
		if item.ProjectKey == projectKey {
			repoList = append(repoList, item)
		}
	}
	return repoList
}

func (jm JfrogMigrate) GetJfrogArtifact(context *config.Context, repoKey string) []model.Arti {
	fileItemList := jm.Jfrog.GetArtifacts(context, repoKey)
	var artiList []model.Arti
	for _, item := range fileItemList.Files {
		var artiItem model.Arti
		artiItem.Name = item.Uri
		artiItem.Repo = repoKey
		artiItem.RepoMapping = repoKey
		artiItem.Path = item.Uri
		artiItem.Sha1 = item.Sha1
		artiItem.Sha256 = item.Sha2
		artiItem.Md5 = item.Md5
		artiList = append(artiList, artiItem)
	}
	return artiList
}
