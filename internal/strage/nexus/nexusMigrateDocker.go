package strage

import (
	"artifactMigrateTools/internal/api/nexus"
	"artifactMigrateTools/internal/common"
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/model"
	"artifactMigrateTools/internal/util"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

type NexusMigrateDocker struct {
}

func NewNexusMigrateDocker() NexusMigrateDocker {
	instance := new(NexusMigrateDocker)
	return *instance
}

func (md NexusMigrateDocker) MigrateRepo(context *config.Context, repo model.Repositories) bool {
	return common.CreateRepository(context, repo)
}

func (md NexusMigrateDocker) MigrateArti(context *config.Context, arti model.Arti) bool {
	config, er := config.NewConfig()
	if er != nil {
		log.Println("读取配置文件失败: ", er)
	}
	nx := &nexus.Nexus{
		util.HttpClient{
			BaseURL:  config.SourceRepo.URL,
			Username: config.SourceRepo.Username,
			Password: config.SourceRepo.Password,
			Header:   http.Header{},
		},
	}
	repoK := util.If(arti.OriginRepo == "", arti.Repo, arti.OriginRepo).(string)
	dockerManifests := nx.ReadManifests(context, repoK, arti.Path)

	repoKey := util.If(arti.VirtualRepo == "", repoK, arti.VirtualRepo).(string)
	arr := strings.Split(dockerManifests.Name, repoKey+"/")
	repoName := arr[1]

	arti.Path = path.Join(repoName, dockerManifests.Tag)
	layerList := AssembleLaysList(arti, dockerManifests)
	layerFlag := CreateDockerLayer(context, layerList)
	dockerFlag := CreateDockerManifest(context, dockerManifests, arti)

	if layerFlag && dockerFlag {
		return true
	}
	return false
}

func CreateDockerLayer(context *config.Context, artiList []model.Arti) bool {
	var flag bool
	for _, arti := range artiList {
		log.Println("正在创建制品: ", arti.Path)
		flag = common.CreateArtifact(context, arti)
		flagSum := common.UpdateArtifactSum(context, arti)

		if !flag || !flagSum {
			break
		}
	}
	return flag
}

func CreateDockerManifest(context *config.Context, dockerManifests nexus.NexusDockerManifests, artiConf model.Arti) bool {
	repoName := util.If(artiConf.VirtualRepo == "", artiConf.RepoMapping, artiConf.VirtualRepo).(string)

	version := dockerManifests.Tag
	sha256 := artiConf.Sha256
	// 获取属性
	prop := GetProp(repoName, version, sha256)
	var arti model.Arti
	path := path.Join(artiConf.Path, "manifest.json")
	arti.Name = "manifest.json"
	arti.Repo = artiConf.Repo
	arti.ProtocolType = artiConf.ProtocolType
	arti.Sha256 = sha256
	arti.Path = path
	log.Println("正在创建制品: ", arti.Path)

	flagArti := common.CreateArtifact(context, arti)
	flagProp := common.UpdateArtifactProp(context, artiConf.Repo, path, prop)
	flagSum := common.UpdateArtifactSum(context, arti)
	return false
	if flagArti && flagProp && flagSum {
		return true
	}
	return false
}

func AssembleLaysList(artiConfig model.Arti, manifestsConf nexus.NexusDockerManifests) []model.Arti {
	var artiList []model.Arti
	for _, layer := range manifestsConf.FsLayers {
		var arti model.Arti
		sha256 := layer.BlobSum
		arti.Name = strings.Replace(sha256, "sha256:", "sha256__", 1)
		arti.Repo = artiConfig.Repo
		arti.ProtocolType = artiConfig.ProtocolType
		arti.Path = path.Join(artiConfig.Path, arti.Name)
		// todo
		arti.Sha256 = strings.Replace(sha256, "sha256:", "", 1)
		artiList = append(artiList, arti)
	}
	return artiList
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
