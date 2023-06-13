package strage

import (
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
)

type Migrate interface {
	// 迁移仓库
	MigrateRepo(context *config.Context, repo model.Repositories) bool

	// 迁移制品
	MigrateArti(context *config.Context, arti model.Arti) bool
}
