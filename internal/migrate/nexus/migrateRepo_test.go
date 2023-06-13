package migrate

import (
	"artifactMigrateTools/internal/common"
	"artifactMigrateTools/internal/migrate"
	"artifactMigrateTools/internal/model"
	"testing"
)

// 迁移所有空间
func TestMigrateRepo_MigrateProjectAll(t *testing.T) {
	rms := new(migrate.RepoMigrateService)
	orm := new(migrate.RepoMigrate)
	_ = orm.MigrateProject(rms.GetProjectListByAll())
}

// 迁移指定空间
func TestMigrateRepo_MigrateProject(t *testing.T) {
	rms := new(migrate.RepoMigrateService)
	orm := new(migrate.RepoMigrate)
	_ = orm.MigrateProject(rms.GetProjectListByProjectKey("osc-mdk"))
}

// 迁移所有仓库
func TestMigrateRepo_MigrateRepoAll(t *testing.T) {
	rms := new(migrate.RepoMigrateService)
	orm := new(migrate.RepoMigrate)
	resultList := orm.MigrateRepo(rms.GetRepoListByAll(), "nexus")
	if resultList != nil {
		common.UpdateRepoYaml(resultList)
	}
}

// 迁移指定空间仓库
func TestMigrateRepo_MigrateRepoByProject(t *testing.T) {
	rms := new(migrate.RepoMigrateService)
	orm := new(migrate.RepoMigrate)
	allRepoList := rms.GetRepoListByAll()
	resultList := orm.MigrateRepo(rms.GetRepoListByProjectKey("osc-mdk-repo"), "nexus")
	if resultList != nil {
		mergeList := common.MergeMigratedRepo(allRepoList, resultList)
		common.UpdateRepoYaml(mergeList)
	}
}

func TestMigrateRepo_MigrateArtiAll(t *testing.T) {
	rms := new(migrate.RepoMigrateService)
	orm := new(migrate.RepoMigrate)
	var artiResultList []model.Arti

	artiAllList := rms.GetArtiListByAll()
	repoList := rms.GetRepoListByProjectKey("osc-mdk-repo")
	for _, repoItem := range repoList {
		repoListByRepoKey := rms.GetArtiListByRepoKey(repoItem.RepoKey)
		artiResults := orm.MigrateArti(repoListByRepoKey, repoItem.RepoKey, "nexus")
		artiResultList = append(artiResultList, artiResults...)
	}
	if artiResultList != nil {
		mergeList := common.MergeMigratedArti(artiAllList, artiResultList)
		common.UpdateNodeYaml(mergeList)
	}

}

func TestMigrateRepo_MigrateArtiByRepoKey(t *testing.T) {
	rms := new(migrate.RepoMigrateService)
	orm := new(migrate.RepoMigrate)
	artiAllList := rms.GetArtiListByAll()
	artiResultList := orm.MigrateArti(rms.GetArtiListByRepoKey("mdk-helm-local"), "mdk-helm-local", "nexus")
	if artiResultList != nil {
		mergeList := common.MergeMigratedArti(artiAllList, artiResultList)
		common.UpdateNodeYaml(mergeList)
	}
}
