package migrate

import (
	"artifactMigrateTools/internal/common"
	"artifactMigrateTools/internal/config"
	migrateJfrog "artifactMigrateTools/internal/migrate/jfrog"
	migrateNexus "artifactMigrateTools/internal/migrate/nexus"
	"artifactMigrateTools/internal/model"
	"fmt"
)

type MigrateExcute struct {
	Context  *config.Context // 上下文
	ExitCode int             //  脚本执行最终退出状态
}

func NewMigrateExcute(context *config.Context) *MigrateExcute {
	return &MigrateExcute{
		Context: context,
	}
}

func (me MigrateExcute) Excute() {
	context := me.Context.Pool
	// 调用辅助函数Println打印日志到标准
	me.ExcuteInit(context["init"])
	me.ExcuteVerifyConfig(context["verifyConfig"])
	me.ExcuteGenericProject(context["genericProject"])
	me.ExcuteMigrateProject(context["migrateProject"])
	me.ExcuteGenericRepo(context["genericRepo"])
	me.ExcuteMigrateRepo(context["migrateRepo"])
	me.ExcuteGenericArtifacts(context["genericArtifacts"])
	me.ExcuteMigrateArtifacts(context["migrateArtifacts"])
}

func (me MigrateExcute) ExcuteInit(content string) {
	if content != "" {
		filePath := me.Context.GetConfigPath()
		me.Context.Loggers.SendLoggerInfo(fmt.Sprint("生成配置文件：", filePath))
		common.GenericConfig()
	}
}

func (me MigrateExcute) ExcuteVerifyConfig(content string) {
	config, er := config.NewConfig()
	if er != nil {
		me.Context.Loggers.SendLoggerError("读取config.yaml文件失败", er)
	}
	if content != "" {
		switch config.SourceRepo.Type {
		case "oldRepo":
			orm := migrateNexus.NewNexusRepoMigrate(config)
			orm.VerifyConfig(me.Context)
			return
		case "nexus":
			orm := migrateNexus.NewNexusRepoMigrate(config)
			orm.VerifyConfig(me.Context)
			return
		case "jfrog":
			orm := migrateJfrog.NewJfrogMigrate(config)
			orm.VerifyConfig(me.Context)
			return
		default:
			ormNexus := migrateNexus.NewNexusRepoMigrate(config)
			ormNexus.VerifyConfig(me.Context)
			orm := NewRepoMigrate(config)
			orm.VerifyConfig(me.Context)
			return
		}
	}
}

func (me MigrateExcute) ExcuteGenericProject(content string) {
	config, _ := config.NewConfig()

	if content != "" {
		filePath := me.Context.GetProjectPath()
		me.Context.Loggers.SendLoggerInfo(fmt.Sprint("生成配置文件：", filePath))

		switch config.SourceRepo.Type {
		case "oldRepo":
			ng := new(migrateNexus.NexusGeneric)
			ng.GenericNexusProjectYaml(me.Context)
			return
		case "nexus":
			ng := new(migrateNexus.NexusGeneric)
			ng.GenericNexusProjectYaml(me.Context)
			return
		case "jfrog":
			jg := new(migrateJfrog.JfrogGeneric)
			jg.GenericJfrogProjectYaml(me.Context)
			return
		default:
			ng := new(migrateNexus.NexusGeneric)
			ng.GenericNexusProjectYaml(me.Context)
			return
		}
		return
	}
}
func (me MigrateExcute) ExcuteMigrateProject(content string) {
	if content != "" {
		rms := new(RepoMigrateService)
		orm := new(RepoMigrate)
		if content == "all" {
			_ = orm.MigrateProject(me.Context, rms.GetProjectListByAll())
		} else {
			allProjectList := rms.GetProjectListByAll()
			resultList := orm.MigrateProject(me.Context, rms.GetProjectListByProjectKey(content))
			if resultList != nil {
				mergeList := common.MergeMigratedProject(allProjectList, resultList)
				common.UpdateProjectYaml(mergeList)
			}
		}
	}
}

func (me MigrateExcute) ExcuteGenericRepo(content string) {
	config, _ := config.NewConfig()

	if content != "" {
		filePath := me.Context.GetRepoPath()
		me.Context.Loggers.SendLoggerInfo(fmt.Sprint("生成配置文件：", filePath))

		switch config.SourceRepo.Type {
		case "oldRepo":
			ng := new(migrateNexus.NexusGeneric)
			ng.GenericNexusRepoYaml(me.Context)
			return
		case "nexus":
			ng := new(migrateNexus.NexusGeneric)
			ng.GenericNexusRepoYaml(me.Context)
			return
		case "jfrog":
			jg := new(migrateJfrog.JfrogGeneric)
			jg.GenericJfrogRepoYaml(me.Context)
			return
		default:
			ng := new(migrateNexus.NexusGeneric)
			ng.GenericNexusRepoYaml(me.Context)
			return
		}
	}
}

func (me MigrateExcute) ExcuteMigrateRepo(content string) {
	config, _ := config.NewConfig()

	if content != "" {
		rms := new(RepoMigrateService)
		orm := new(RepoMigrate)
		if content == "all" {
			resultList := orm.MigrateRepo(me.Context, rms.GetRepoListByAll(), config.SourceRepo.Type)
			if resultList != nil {
				common.UpdateRepoYaml(resultList)
			}
		} else {
			allRepoList := rms.GetRepoListByAll()
			resultList := orm.MigrateRepo(me.Context, rms.GetRepoListByProjectKey(content), config.SourceRepo.Type)
			if resultList != nil {
				mergeList := common.MergeMigratedRepo(allRepoList, resultList)
				common.UpdateRepoYaml(mergeList)
			}
		}
	}
}

func (me MigrateExcute) ExcuteGenericArtifacts(content string) {
	config, _ := config.NewConfig()

	if content != "" {
		filePath := me.Context.GetArtifactPath()
		me.Context.Loggers.SendLoggerInfo("生成配置文件：", filePath)
		switch config.SourceRepo.Type {
		case "oldRepo":
			ng := new(migrateNexus.NexusGeneric)
			ng.GenericNexusArtiYaml(me.Context)
			return
		case "nexus":
			ng := new(migrateNexus.NexusGeneric)
			ng.GenericNexusArtiYaml(me.Context)
			return
		case "jfrog":
			jg := new(migrateJfrog.JfrogGeneric)
			jg.GenericJfrogArtiYaml(me.Context, content)
			return
		default:
			ng := new(migrateNexus.NexusGeneric)
			ng.GenericNexusArtiYaml(me.Context)
			return
		}
	}
}

func (me MigrateExcute) ExcuteMigrateArtifacts(content string) {
	config, _ := config.NewConfig()

	if content != "" {
		rms := new(RepoMigrateService)
		orm := new(RepoMigrate)
		if content == "all" {
			var artiResultList []model.Arti
			artiAllList := rms.GetArtiListByAll()
			repoList := rms.GetRepoListByAll()
			for _, repoItem := range repoList {
				repoListByRepoKey := rms.GetArtiListByRepoKey(repoItem.RepoKey)
				artiResults := orm.MigrateArti(me.Context, repoListByRepoKey, repoItem.RepoKey, config.SourceRepo.Type)
				artiResultList = append(artiResultList, artiResults...)
			}
			if artiResultList != nil {
				mergeList := common.MergeMigratedArti(artiAllList, artiResultList)
				common.UpdateNodeYaml(mergeList)
			}
		} else {
			artiAllList := rms.GetArtiListByAll()
			var artiResultList []model.Arti
			repoList := rms.GetRepoListByProjectKey(content)
			for _, item := range repoList {
				artiResultLists := orm.MigrateArti(me.Context, rms.GetArtiListByRepoKey(item.RepoKeyMapping), content, config.SourceRepo.Type)
				artiResultList = append(artiResultList, artiResultLists...)
			}
			if artiResultList != nil {
				mergeList := common.MergeMigratedArti(artiAllList, artiResultList)
				common.UpdateNodeYaml(mergeList)
			}
		}
	}
}
