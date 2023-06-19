package strage

import (
	"artifactMigrateTools/internal/api/jfrog"
	"artifactMigrateTools/internal/common"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
	"artifactMigrateTools/internal/util"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type JfrogMigrateDocker struct {
}

func NewJfrogMigrateDocker() JfrogMigrateDocker {
	instance := new(JfrogMigrateDocker)
	return *instance
}

func (jmd JfrogMigrateDocker) MigrateRepo(context *config.Context, repo model.Repositories) bool {
	return common.CreateRepository(context, repo)
}

func (jmd JfrogMigrateDocker) MigrateArti(context *config.Context, arti model.Arti) bool {
	config, er := config.NewConfig()
	if er != nil {
		log.Println("读取配置文件失败: ", er)
	}
	jf := &jfrog.Jfrog{
		util.HttpClient{
			BaseURL:  config.SourceRepo.URL,
			Username: config.SourceRepo.Username,
			Password: config.SourceRepo.Password,
			Header:   http.Header{},
		},
	}
	util.Auth(jf.HttpClient)

	if strings.Contains(arti.Path, "manifest.json") {
		re := regexp.MustCompile(`\/([^\/]+\/manifest\.json)$`)
		match := re.FindStringSubmatch(arti.Path)
		if len(match) < 2 {
			// 处理错误
		}

		manifestPath := match[1]
		version := strings.Split(manifestPath, "/")[0]

		prop := GetProp(arti.Repo, version, arti.Sha256)
		flag := common.CreateArtifact(context, arti)
		flagProp := common.UpdateArtifactProp(context, arti.Repo, arti.Path, prop)
		sumFlag := common.UpdateArtifactSum(context, arti)

		return flag && flagProp && sumFlag
	}

	flag := common.CreateArtifact(context, arti)
	sumFlag := common.UpdateArtifactSum(context, arti)
	return flag && sumFlag
}

// 更新属性
func GetProp(repoName, version, sha256 string) string {
	var propertieMap map[string]string
	propertieMap = make(map[string]string)
	propertieMap["docker.repoName"] = repoName
	propertieMap["docker.manifest"] = version
	propertieMap["docker.manifest.digest"] = "sha256:" + sha256
	propertieMap["sha256"] = sha256
	propertieMap["docker.manifest.type"] = "application/vnd.docker.distribution.manifest.v2+json"
	propertieMap["artifactory.content-type"] = "artifactory.content-type"
	var keyValuePairs []string
	for key, value := range propertieMap {
		keyValuePairs = append(keyValuePairs, fmt.Sprintf("%s=%s", key, value))
	}
	propertiesString := strings.Join(keyValuePairs, ";")
	return propertiesString
}
