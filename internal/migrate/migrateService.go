package migrate

import (
	"artifactMigrateTools/internal/common"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
	"log"
)

type RepoMigrateService struct {
}

// 获取所有空间
func (rms RepoMigrateService) GetProjectListByAll() []model.Project {
	projects, er := config.NewProject()
	if er != nil {
		log.Println(projects)
	}
	return projects
}

// 获取所有空间
func (rms RepoMigrateService) GetProjectListByProjectKey(projectKey string) []model.Project {
	allProjectList := rms.GetProjectListByAll()
	resultList := common.FilterProjectByProjectKey(allProjectList, projectKey)
	return resultList
}

func (rms RepoMigrateService) GetRepoListByAll() []model.Repositories {
	projects, er := config.NewProject()
	repositories, er := config.NewRepositories()
	if er != nil {
		log.Println(repositories)
	}
	resultRepos := common.MergeProjectAndRepositories(projects, repositories)
	resultFillRepos := common.FillVirtualRepo(projects, resultRepos)
	orderRepos := common.OrderRepo(resultFillRepos)
	return orderRepos
}

func (rms RepoMigrateService) GetRepoListByProjectKey(projectKey string) []model.Repositories {
	allRepoList := rms.GetRepoListByAll()
	resultRepos := common.FilterRepoByWorkspaceKey(allRepoList, projectKey)
	return resultRepos
}

func (rms RepoMigrateService) GetArtiListByAll() []model.Arti {
	repositories, er := config.NewRepositories()
	artiList, er := config.NewArti()
	if er != nil {
		log.Println(artiList)
	}
	resultRepos := common.MergeRepositoriesAndArti(*repositories, artiList)

	return resultRepos
}

func (rms RepoMigrateService) GetArtiListByRepoKey(repoKey string) []model.Arti {
	allArti := rms.GetArtiListByAll()
	resultList := common.FilterArtiByRepoKey(allArti, repoKey)
	return resultList
}
