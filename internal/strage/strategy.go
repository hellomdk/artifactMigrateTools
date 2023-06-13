package strage

import (
	"artifactMigrateTools/internal/common"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
)

type Strategy struct {
	Strategy Migrate
}

func NewStrategySelf() Strategy {
	instance := new(Strategy)
	return *instance
}

// 协议类型&&制品库类型（oldRepo、jfrog、nexus）
func NewStrategy(protocolType, migrateType string) Strategy {
	c := new(Strategy)
	switch protocolType {
	case "docker":
		c.Strategy = NewMigrateDocker(migrateType)
		break
	case "maven":
		c.Strategy = NewMigrateMaven()
		break
	default:
		c.Strategy = NewMigrateGeneric()
		break
	}
	return *c
}

// 执行入口
func (s Strategy) MigrateRepo(context *config.Context, repo model.Repositories) bool {
	return s.Strategy.MigrateRepo(context, repo)
}

// 执行入口
func (s Strategy) MigrateArti(context *config.Context, arti model.Arti) bool {
	return s.Strategy.MigrateArti(context, arti)
}

// 同步空间信息
func (s Strategy) MigrateProject(context *config.Context, project model.Project) bool {
	return common.CreateProject(context, project.ProjectKey)
}
