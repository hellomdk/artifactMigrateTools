package strage

import (
	"artifactMigrateTools/internal/common"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
)

type MigrateGeneric struct {
}

func NewMigrateGeneric() MigrateGeneric {
	instance := new(MigrateGeneric)
	return *instance
}

func (mg MigrateGeneric) MigrateRepo(context *config.Context, repo model.Repositories) bool {
	return common.CreateRepository(context, repo)
}

func (mg MigrateGeneric) MigrateArti(context *config.Context, arti model.Arti) bool {
	artiFlag := common.CreateArtifact(context, arti)
	sumFlag := common.UpdateArtifactSum(context, arti)
	if artiFlag && sumFlag {
		return true
	}
	return false
}
