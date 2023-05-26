package migrate

import (
	"fmt"
	"jfrogToArtifact/pkg/artifact/config"
	"strings"
	"testing"
)

func TestMigrateRepo_MigrateRepo(t *testing.T) {
	repositories, er := config.NewRepositories()
	if er != nil {
		fmt.Println(repositories)
	}
	_ = MigrateRepo(repositories)
}

func TestMigrateRepo_MigrateArti(t *testing.T) {
	arti, er := config.NewArti()
	if er != nil {
		fmt.Println(arti)
	}
	_ = MigrateArti(arti)
}

func TestUpdate_Arti_Checksum(t *testing.T) {
	arti, er := config.NewArti()
	if er != nil {
		fmt.Println(arti)
	}
	_ = Update_Arti_Checksum(arti)
}

func TestUpdate_Arti_Checksum2(t *testing.T) {
	var propertieMap map[string]string
	propertieMap = make(map[string]string)

	propertieMap["docker.repoName"] = "test/nginx"
	propertieMap["docker.manifest"] = "v1.0"
	propertieMap["docker.manifest.digest"] = "sha256:bfb112db4075460ec042ce13e0b9c3ebd982f93ae0be155496d050bb70006750"
	propertieMap["docker.manifest.type"] = "application/vnd.docker.distribution.manifest.v2+json"
	propertieMap["artifactory.content-type"] = "artifactory.content-type"
	propertieMap["sha256"] = "bfb112db4075460ec042ce13e0b9c3ebd982f93ae0be155496d050bb70006750"
	var keyValuePairs []string
	for key, value := range propertieMap {
		keyValuePairs = append(keyValuePairs, fmt.Sprintf("%s=%s", key, value))
	}
	propertiesString := strings.Join(keyValuePairs, ";")
	fmt.Println(propertiesString)
}
