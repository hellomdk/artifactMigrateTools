package migrate

import (
	"artifactMigrateTools/internal/common"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
)

type NexusGeneric struct {
}

// 生成project.yaml
func (ng NexusGeneric) GenericNexusProjectYaml(context *config.Context) {
	config, er := config.NewConfig()
	if er != nil {
		context.Loggers.SendLoggerError("读取config.yaml文件失败", er)
	}
	orm := NewNexusRepoMigrate(config)
	resultList := orm.GetNexusRepositories(context)
	projectList := common.GetProjectKey(resultList)
	common.WriteNexusProject(projectList)
}

// 生成repo.yaml
func (ng NexusGeneric) GenericNexusRepoYaml(context *config.Context) {
	config, er := config.NewConfig()
	if er != nil {
		context.Loggers.SendLoggerError("读取config.yaml文件失败: ", er)
	}
	orm := NewNexusRepoMigrate(config)
	resultList := orm.GetNexusRepositories(context)
	common.WriteNexusRepo(resultList)
}

// 生成Node.yaml
func (ng NexusGeneric) GenericNexusArtiYaml(context *config.Context) {
	config, er := config.NewConfig()
	if er != nil {
		context.Loggers.SendLoggerError("读取config.yaml文件失败: ", er)
	}
	orm := NewNexusRepoMigrate(config)
	var resultList []model.Arti
	repoList := orm.GetNexusRepositories(context)
	for _, repoItem := range repoList {
		artiList := orm.GetNexusArtifact(context, repoItem.RepoKey)
		resultList = append(resultList, artiList...)
	}
	common.WriteNexusArti(resultList)
}
