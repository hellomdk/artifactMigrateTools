package common

import (
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
	"artifactMigrateTools/internal/util"
)

func GenericConfig() {
	c := new(config.Context)
	filePath := c.GetConfigPath()
	if !util.PathExists(filePath) {
		util.CreatePathDir(filePath)
	}
	var configYaml model.ConfigYaml
	config := &model.Config{
		WorkDir:      "/home/migrate",
		Threads:      2,
		Logging:      "INFO",
		ChecksumCalc: false,
		SourceRepo: model.Repo{
			URL:      "http://192.168.80.50:6967",
			Type:     "nexus",
			Username: "admin",
			Password: "admin123",
		},
		TargetRepo: model.Repo{
			URL:      "http://test.gitee.work/artifactory",
			Type:     "repo",
			Username: "admin",
			Password: "cq123456",
		},
	}

	configYaml.Config = *config
	util.WriteCSV(filePath, configYaml)
}

// 更新yaml
func UpdateProjectYaml(projectList []model.Project) {
	c := new(config.Context)
	filePath := c.GetProjectPath()
	if !util.PathExists(filePath) {
		util.CreatePathDir(filePath)
	}
	var projectYaml model.ProjectYaml
	projectYaml.Project = projectList
	util.WriteCSV(filePath, projectYaml)
}

// 更新yaml
func UpdateRepoYaml(repoList []model.Repositories) {
	c := new(config.Context)
	filePath := c.GetRepoPath()
	if !util.PathExists(filePath) {
		util.CreatePathDir(filePath)
	}
	var repoYaml model.RepositoriesYaml
	repoYaml.Repositories = repoList
	util.WriteCSV(filePath, repoYaml)
}

// 更新yaml
func UpdateNodeYaml(artiList []model.Arti) {
	c := new(config.Context)
	filePath := c.GetArtifactPath()
	if !util.PathExists(filePath) {
		util.CreatePathDir(filePath)
	}
	var nodeYaml model.NodeYaml
	nodeYaml.Artifact = artiList
	util.WriteCSV(filePath, nodeYaml)
}
