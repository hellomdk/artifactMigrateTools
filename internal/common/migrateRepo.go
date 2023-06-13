package common

import (
	"artifactMigrateTools/internal/api/repo"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
	"artifactMigrateTools/internal/util"
	"net/http"
)

func CreateRepository(context *config.Context, repositories model.Repositories) bool {
	config, er := config.NewConfig()
	if er != nil {
		context.Loggers.SendLoggerError("获取配置文件失败: ", er)
	}
	repoCli := &repo.Repo{
		util.HttpClient{
			BaseURL:  config.TargetRepo.URL,
			Username: config.TargetRepo.Username,
			Password: config.TargetRepo.Password,
			Header:   http.Header{},
		},
	}
	util.Auth(repoCli.HttpClient)

	repoData := repo.RepoRepository{
		RepoKey:     repositories.RepoKeyMapping,
		RepoType:    repositories.RepoType,
		Description: repositories.Description,
		Layout:      repositories.Layout,
		Browse:      true,
		Url:         repositories.Url,
		Username:    repositories.Username,
		Password:    repositories.Password,
		ProtocolSpecific: repo.RepoProtocolType{
			ProtocolType:     repositories.ProtocolType,
			DockerApiVersion: "V2",
		},
		ProjectKey:            repositories.ProjectKeyMapping,
		SelectedRepositories:  ConvertSelectedRepositories(repositories.SelectedRepositories),
		ResolvedRepositories:  ConvertSelectedRepositories(repositories.SelectedRepositories),
		DefaultDeploymentRepo: repositories.DefaultDeploymentRepo,
	}

	got := repoCli.CreateRepository(context, repoData)
	return got
}

func ConvertSelectedRepositories(selectedRepositories []model.SelectedRepositories) []repo.RepoSelectedRepositories {
	var resultList []repo.RepoSelectedRepositories
	for _, selectRepo := range selectedRepositories {
		var newRepo repo.RepoSelectedRepositories
		newRepo.RepoKey = selectRepo.RepoKey
		newRepo.ProjectKey = selectRepo.ProjectKey
		newRepo.Key = selectRepo.Key
		newRepo.Type = selectRepo.Type

		resultList = append(resultList, newRepo)
	}
	return resultList
}
