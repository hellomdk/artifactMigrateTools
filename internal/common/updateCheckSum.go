package common

import (
	"artifactMigrateTools/internal/api/repo"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
	"artifactMigrateTools/internal/util"
	"log"
	"net/http"
	"strings"
)

func UpdateArtifactSum(context *config.Context, arti model.Arti) bool {
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
	jf.HttpClient.Header.Add("Content-Type", "application/octet-stream")
	jf.HttpClient.Header.Add("X-Checksum-Deploy", "true")
	if arti.Md5 == "" {
		jf.HttpClient.Header.Add("X-Checksum-Md5", "d41d8cd98f00b204e9800998ecf8427e")
	} else {
		jf.HttpClient.Header.Add("X-Checksum-Md5", arti.Md5)
	}

	if arti.Sha1 == "" {
		jf.HttpClient.Header.Add("X-Checksum-Sha1", "da39a3ee5e6b4b0d3255bfef95601890afd80709")
	} else {
		jf.HttpClient.Header.Add("X-Checksum-Sha1", arti.Sha1)
	}

	if arti.Sha256 == "" {
		jf.HttpClient.Header.Add("X-Checksum-Sha256", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	} else {
		jf.HttpClient.Header.Add("X-Checksum-Sha256", arti.Sha256)
	}

	got := jf.UpdateArtifactCheckSum(context, arti.Repo, arti.Path, arti.Sha1, arti.Sha256, arti.Md5)
	return got
}

func UpdateArtifactProp(context *config.Context, repoKey, repoPath, properties string) bool {
	config, er := config.NewConfig()
	if er != nil {
		log.Println("获取配置文件失败: ", er)
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
	repoPath = strings.TrimPrefix(repoPath, "/")

	got := jf.UpdateArtifactProperties(context, repoKey, repoPath, properties)
	return got
}
