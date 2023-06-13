package common

import (
	"artifactMigrateTools/internal/api/repo"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/util"
	"net/http"
)

func CreateProject(context *config.Context, project string) bool {
	config, er := config.NewConfig()
	if er != nil {
		context.Loggers.SendLoggerError("获取配置文件失败: ", er)
	}

	jf := &repo.Repo{
		util.HttpClient{
			BaseURL:  config.TargetRepo.URL,
			Username: config.TargetRepo.Username,
			Password: config.TargetRepo.Password,
			Header:   http.Header{},
		},
	}
	util.Auth(jf.HttpClient)

	got := jf.CreateProject(context, project)
	return got
}
