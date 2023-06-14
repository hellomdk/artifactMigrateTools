package migrate

import (
	"artifactMigrateTools/internal/api/repo"
	"artifactMigrateTools/internal/common"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
	"artifactMigrateTools/internal/strage"
	"artifactMigrateTools/internal/util"
	"fmt"
	"net/http"
)

type RepoMigrate struct {
	Repo repo.Repo
}

func NewRepoMigrate(config *model.Config) RepoMigrate {
	r := &repo.Repo{
		util.HttpClient{
			BaseURL:  config.TargetRepo.URL,
			Username: config.TargetRepo.Username,
			Password: config.TargetRepo.Password,
			Header:   http.Header{},
		},
	}
	util.Auth(r.HttpClient)

	instance := new(RepoMigrate)
	instance.Repo = *r
	return *instance
}

func (orm RepoMigrate) VerifyConfig(context *config.Context) bool {
	body, _ := orm.Repo.Ping(context)
	if body != "" {
		context.Loggers.SendLoggerInfo(fmt.Sprint("验证连接成功: ", orm.Repo.HttpClient.BaseURL))
		return true
	}
	return false
}

func (orm RepoMigrate) MigrateProject(context *config.Context, projectList []model.Project) []model.Project {
	// 读取 value
	var resultList []model.Project
	for _, projectItem := range projectList {
		if !projectItem.Migrated {
			context.Loggers.SendLoggerInfo(fmt.Sprint("正在迁移空间：", projectItem.ProjectKey))
			strategy := strage.NewStrategySelf()
			projectItem.Migrated = strategy.MigrateProject(context, projectItem)
			if projectItem.Migrated {
				resultList = append(resultList, projectItem)
				context.Loggers.SendLoggerInfo(fmt.Sprint("迁移空间: ", projectItem.ProjectKey, "成功"))
			} else {
				context.Loggers.SendLoggerError(fmt.Sprint("迁移空间: ", projectItem.ProjectKey, "失败"), nil)
			}
		} else {
			context.Loggers.SendLoggerInfo(fmt.Sprint("空间: ", projectItem.ProjectKey, "已存在"))
		}
	}

	return resultList
}

func (orm RepoMigrate) MigrateRepo(context *config.Context, repoList []model.Repositories, migrateType string) []model.Repositories {
	// 读取 value
	var resultList []model.Repositories
	for _, repo := range repoList {
		if !repo.Migrated {
			context.Loggers.SendLoggerInfo("正在迁移仓库: ", repo.RepoKey)
			strategy := strage.NewStrategy(repo.ProtocolType, migrateType)
			repo.Migrated = strategy.MigrateRepo(context, repo)
			if repo.Migrated {
				context.Loggers.SendLoggerInfo("迁移仓库: ", repo.RepoKey, "成功")
			} else {
				context.Loggers.SendLoggerInfo("迁移仓库: ", repo.RepoKey, "失败")
			}
		}
		resultList = append(resultList, repo)
	}

	return resultList
}

func (orm RepoMigrate) MigrateArti(context *config.Context, artiList []model.Arti, repoKey, migrateType string) []model.Arti {
	var artiResultList []model.Arti
	context.Loggers.SendLoggerInfo(fmt.Sprint("正在迁移仓库: ", repoKey, "   制品总数量: ", len(artiList)))

	// 开启携程迁移制品文件
	context.ThreadPool.Run()
	for _, arti := range artiList {
		//context.ThreadPool.AddTask(
		//	worker.Task{
		//		Handler: func(args interface{}) {
		// 执行任务的代码
		if !arti.Migrated {
			//log.Println("正在同步制品: ", arti.Path)
			strategy := strage.NewStrategy(arti.ProtocolType, migrateType)
			arti.Migrated = strategy.MigrateArti(context, arti)
			if arti.Migrated {
				context.Loggers.SendLoggerInfo(fmt.Sprintf("goroutine ID: %d", util.GetGoroutineID()),
					"迁移制品: ", arti.Path, "成功")
			} else {
				context.Loggers.SendLoggerError(fmt.Sprintf("goroutine ID: %d, 迁移制品: %s 失败",
					util.GetGoroutineID(), arti.Path), nil)
			}
		} else {
			context.Loggers.SendLoggerInfo(fmt.Sprintf("goroutine ID: %d", util.GetGoroutineID()),
				"制品: ", arti.Path, "已存在")
		}
		artiResultList = append(artiResultList, arti)
		//	},
		//	Args: arti,
		//})
	}
	success := common.CountMigrateArtiSuccess(artiResultList)
	context.Loggers.SendLoggerInfo(fmt.Sprintf("仓库: %s 制品迁移完成，迁移成功数量: %d 迁移失败数量: %d ", repoKey, success, len(artiList)-success))

	return artiResultList
}
