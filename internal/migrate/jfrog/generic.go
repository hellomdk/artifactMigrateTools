package migrate

import (
	"artifactMigrateTools/internal/common"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/migrate/jfrog/chain"
	"artifactMigrateTools/internal/model"
)

type JfrogGeneric struct {
}

// 生成project.yaml
func (jg JfrogGeneric) GenericJfrogProjectYaml(context *config.Context) {
	config, er := config.NewConfig()
	if er != nil {
		context.Loggers.SendLoggerError("读取config.yaml文件失败", er)
	}
	jm := NewJfrogMigrate(config)
	resultList := jm.GetJfrogRepositories(context)
	projectList := common.GetProjectKey(resultList)
	common.WriteNexusProject(projectList)
}

// 生成repo.yaml
func (jg JfrogGeneric) GenericJfrogRepoYaml(context *config.Context) {
	config, er := config.NewConfig()
	if er != nil {
		context.Loggers.SendLoggerError("读取config.yaml文件失败: ", er)
	}
	jm := NewJfrogMigrate(config)
	resultList := jm.GetJfrogRepositories(context)
	common.WriteNexusRepo(resultList)
}

// 生成Node.yaml
func (jg JfrogGeneric) GenericJfrogArtiYaml(context *config.Context, projectKey string) {
	config, er := config.NewConfig()
	if er != nil {
		context.Loggers.SendLoggerError("读取config.yaml文件失败: ", er)
	}
	jm := NewJfrogMigrate(config)
	var resultList []model.Arti

	repoList := jm.GetJfrogRepositoriesByProjectKey(context, projectKey)
	for _, repoItem := range repoList {
		if repoItem.RepoType == "virtual" {
			continue
		}
		artiList := jm.GetJfrogArtifact(context, repoItem)
		filterList := FilterIndex(artiList)
		resultList = append(resultList, filterList...)
	}
	common.WriteNexusArti(resultList)
}

// 过滤npm索引
func FilterIndex(artiList []model.Arti) []model.Arti {
	var resultList []model.Arti
	center := new(chain.ChainCenter)
	handle := center.AssembleChainHandle()
	for _, item := range artiList {
		if handle.Handle(item.Name) {
			continue
		}
		resultList = append(resultList, item)
	}

	return resultList
}
