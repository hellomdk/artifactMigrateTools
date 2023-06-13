package strage

import (
	"artifactMigrateTools/internal/common"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
	"strings"
)

type MigrateMaven struct {
}

func NewMigrateMaven() MigrateMaven {
	instance := new(MigrateMaven)
	return *instance
}

func (mg MigrateMaven) MigrateRepo(context *config.Context, repo model.Repositories) bool {
	return common.CreateRepository(context, repo)
}

func (mg MigrateMaven) MigrateArti(context *config.Context, arti model.Arti) bool {
	if strings.Contains(arti.Path, ".sha1") || strings.Contains(arti.Path, ".sha256") || strings.Contains(arti.Path, ".md5") || strings.Contains(arti.Path, ".sha512") {
		return true
	}
	artiFlag := common.CreateArtifact(context, arti)
	sumFlag := common.UpdateArtifactSum(context, arti)
	if artiFlag && sumFlag {
		return true
	}
	return false
}
