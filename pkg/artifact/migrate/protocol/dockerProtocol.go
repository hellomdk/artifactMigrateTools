package protocol

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"jfrogToArtifact/pkg/artifact/config"
	"jfrogToArtifact/pkg/artifact/httpclient"
	"jfrogToArtifact/pkg/artifact/migrate/common"
	"jfrogToArtifact/pkg/artifact/nexus"
	"net/http"
	"path"
	"strings"
)

// 上传docker
func DeployDocker(arti config.Arti) {
	config, er := config.NewConfig()
	if er != nil {
		fmt.Println(config)
	}
	nx := &nexus.Nexus{
		httpclient.HttpClient{
			BaseURL:  config.SourceRepo.URL,
			Username: config.SourceRepo.Username,
			Password: config.SourceRepo.Password,
			Header:   http.Header{},
		},
	}
	dockerManifests, str := nx.ReadManifests(arti.RepoOrg, arti.Path)

	repoKey := arti.RepoParent
	arr := strings.Split(dockerManifests.Name, repoKey+"/")
	repoName := arr[1]

	arti.Path = path.Join(repoName, dockerManifests.Tag)
	layerList := AssembleLaysList(arti, dockerManifests)
	CreateDockerLayer(layerList)
	CreateDockerManifest(dockerManifests, arti, str)

}

func CreateDockerLayer(artiList []config.Arti) {
	for _, arti := range artiList {
		fmt.Println("正在创建制品: ", arti.Path)
		common.CreateArtifact(arti)
	}
}
func CreateDockerManifest(dockerManifests nexus.DockerManifests, artiConf config.Arti, str string) {

	repoKey := artiConf.RepoParent
	arr := strings.Split(dockerManifests.Name, repoKey+"/")
	repoName := arr[1]
	version := dockerManifests.Tag
	sha256 := artiConf.Sha256
	// 获取属性
	prop := GetProp(repoName, version, sha256)
	var arti config.Arti
	path := path.Join(repoName, version, "manifest.json")
	arti.Name = "manifest.json"
	arti.Repo = artiConf.Repo
	arti.ProtocolType = artiConf.ProtocolType
	arti.Sha256 = sha256
	arti.Path = path
	fmt.Println("正在创建制品: ", arti.Path)

	common.CreateArtifact(arti)
	common.UpdateArtifactProp(artiConf.Repo, path, prop)
}

func AssembleLaysList(artiConfig config.Arti, manifestsConf nexus.DockerManifests) []config.Arti {
	var artiList []config.Arti
	for _, layer := range manifestsConf.FsLayers {
		var arti config.Arti
		sha256 := layer.BlobSum
		arti.Name = strings.Replace(sha256, "sha256:", "sha256__", 1)
		arti.Repo = artiConfig.Repo
		arti.ProtocolType = artiConfig.ProtocolType
		arti.Path = path.Join(artiConfig.Path, arti.Name)
		arti.Sha256 = strings.Replace(sha256, "sha256:", "", 1)
		artiList = append(artiList, arti)
	}
	return artiList
}

// 更新属性
func GetProp(repoName, version, sha256 string) string {
	var propertieMap map[string]string
	propertieMap = make(map[string]string)

	//arti.Path = "hao-docker/test/nginx/manifests/v1.0"
	//arti.Path = "hao-docker/test/nginx/manifests/sha256\\:bfb112db4075460ec042ce13e0b9c3ebd982f93ae0be155496d050bb70006750"

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

//SHA256生成哈希值
func GetSHA256HashCode(message []byte) string {
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode

	//方法二：
	//bytes2:=sha256.Sum256(message)//计算哈希值，返回一个长度为32的数组
	//hashcode2:=hex.EncodeToString(bytes2[:])//将数组转换成切片，转换成16进制，返回字符串
	//return hashcode2
}
