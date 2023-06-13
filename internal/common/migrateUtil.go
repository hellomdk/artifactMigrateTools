package common

import (
	"artifactMigrateTools/internal/model"
	"sort"
	"strings"
)

func GetPro(format string) (string, string) {
	if strings.Contains(format, "maven") {
		return "maven", "maven-2-default"
	} else if strings.Contains(format, "npm") {
		return "npm", "npm-default"
	} else {
		return format, "simple-default"
	}
}

func GetProjectKey(repoList []model.Repositories) []model.Project {
	// 定义map用于记录已经出现过的ID
	idMap := make(map[string]bool)

	// 定义存放去重后的结果的切片
	uniqueStudents := []model.Project{}

	// 遍历对象数组进行去重
	for _, repoItem := range repoList {
		if _, ok := idMap[repoItem.ProjectKey]; !ok {
			idMap[repoItem.ProjectKey] = true
			var project model.Project
			project.ProjectKey = repoItem.ProjectKey
			project.ProjectKeyMapping = repoItem.ProjectKey
			project.Migrated = false
			uniqueStudents = append(uniqueStudents, project)
		}
	}

	return uniqueStudents
}

func MergeProjectAndRepositories(projectList []model.Project, repositoriesList *[]model.Repositories) []model.Repositories {
	var newRepositoriesList []model.Repositories
	projectMap := GetProjectMap(projectList)
	for _, repoItem := range *repositoriesList {
		repoItem.ProjectKeyMapping = projectMap[repoItem.ProjectKey].ProjectKeyMapping
		newRepositoriesList = append(newRepositoriesList, repoItem)
	}
	return newRepositoriesList
}
func MergeRepositoriesAndArti(repositoriesList []model.Repositories, artiList *[]model.Arti) []model.Arti {
	var newArtiList []model.Arti
	projectMap := GetRepoMap(repositoriesList)
	for _, artiItem := range *artiList {
		artiItem.RepoMapping = projectMap[artiItem.Repo].RepoKeyMapping
		newArtiList = append(newArtiList, artiItem)
	}
	return newArtiList
}

func GetProjectMap(projectList []model.Project) map[string]model.Project {
	// 定义map用于记录已经出现过的ID
	projectMap := make(map[string]model.Project)
	for _, projectItem := range projectList {
		projectMap[projectItem.ProjectKey] = projectItem
	}
	return projectMap
}

func GetRepoMap(repositoriesList []model.Repositories) map[string]model.Repositories {
	// 定义map用于记录已经出现过的ID
	repoMap := make(map[string]model.Repositories)
	for _, repoItem := range repositoriesList {
		repoMap[repoItem.RepoKey] = repoItem
	}
	return repoMap
}
func GetArtiMap(artiList []model.Arti) map[string]model.Arti {
	// 定义map用于记录已经出现过的ID
	artiMap := make(map[string]model.Arti)
	for _, artiItem := range artiList {
		artiMap[artiItem.Path] = artiItem
	}
	return artiMap
}

func WriteNexusProject(resultList []model.Project) {
	UpdateProjectYaml(resultList)
}

func WriteNexusRepo(resultList []model.Repositories) {
	UpdateRepoYaml(resultList)
}

func WriteNexusArti(resultList []model.Arti) {
	UpdateNodeYaml(resultList)
}

// 过滤Project
func FilterProjectByProjectKey(projectList []model.Project, projectKey string) []model.Project {
	var newRepoList []model.Project
	for _, projectItem := range projectList {
		if projectItem.ProjectKey == projectKey {
			newRepoList = append(newRepoList, projectItem)
		}
	}

	return newRepoList
}

// 过滤repo仓库
func FilterRepoByWorkspaceKey(repoList []model.Repositories, projectKey string) []model.Repositories {
	var newRepoList []model.Repositories
	for _, repoItem := range repoList {
		if repoItem.ProjectKey == projectKey {
			newRepoList = append(newRepoList, repoItem)
		}
	}

	return newRepoList
}

// 过滤repo仓库
func FilterArtiByRepoKey(artiList []model.Arti, repoKey string) []model.Arti {
	var newRepoList []model.Arti
	for _, artiItem := range artiList {
		if artiItem.Repo == repoKey {
			newRepoList = append(newRepoList, artiItem)
		}
	}
	return newRepoList
}

func CountMigrateArtiSuccess(artiList []model.Arti) int {
	var count int
	for _, artiItem := range artiList {
		if artiItem.Migrated == true {
			count++
		}
	}
	return count
}

func MergeMigratedProject(projectAllList, projectResultList []model.Project) []model.Project {
	projectMap := GetProjectMap(projectResultList)
	var newResultList []model.Project
	for _, projectItem := range projectAllList {
		if _, ok := projectMap[projectItem.ProjectKey]; ok {
			// 存在
			projectItem.Migrated = projectMap[projectItem.ProjectKey].Migrated
		}
		newResultList = append(newResultList, projectItem)
	}
	return newResultList
}
func MergeMigratedRepo(repoAllList, repoResultList []model.Repositories) []model.Repositories {
	repoMap := GetRepoMap(repoResultList)
	var newResultList []model.Repositories
	for _, repoItem := range repoAllList {
		if _, ok := repoMap[repoItem.RepoKey]; ok {
			// 存在
			repoItem.Migrated = repoMap[repoItem.RepoKey].Migrated
		}
		newResultList = append(newResultList, repoItem)
	}
	return newResultList
}

func MergeMigratedArti(artiAllList, artiResultList []model.Arti) []model.Arti {
	artiMap := GetArtiMap(artiResultList)
	var newResultList []model.Arti
	for _, arti := range artiAllList {
		if _, ok := artiMap[arti.Path]; ok {
			// 存在
			arti.Migrated = artiMap[arti.Path].Migrated
		}
		newResultList = append(newResultList, arti)
	}
	return newResultList
}

func FillVirtualRepo(projectList []model.Project, repoList []model.Repositories) []model.Repositories {
	var newRepoList []model.Repositories
	//projectMap := GetProjectMap(projectList)
	repoMap := GetRepoMap(repoList)
	for _, repoItem := range repoList {
		if repoItem.RepoType == "virtual" {
			var selecteds []model.SelectedRepositories
			for _, selectedRepo := range repoItem.SelectedRepositories {
				rp := repoMap[selectedRepo.RepoKey]
				selectedRepo.RepoKey = rp.RepoKeyMapping
				selectedRepo.Key = selectedRepo.RepoKey
				selectedRepo.Type = rp.RepoType
				selectedRepo.ProjectKey = rp.ProjectKeyMapping
				selecteds = append(selecteds, selectedRepo)
			}
			repoItem.SelectedRepositories = selecteds
		}
		newRepoList = append(newRepoList, repoItem)
	}

	return newRepoList
}

func OrderRepo(repoList []model.Repositories) []model.Repositories {
	sort.SliceStable(repoList, func(i, j int) bool {
		return GetRepoTypeScope(repoList[i].RepoType) < GetRepoTypeScope(repoList[j].RepoType)
	})
	return repoList
}

func GetRepoTypeScope(repoType string) int {
	if repoType == "local" {
		return 1
	} else if repoType == "remote" {
		return 2
	} else if repoType == "virtual" {
		return 3
	} else {
		return 4
	}
}
