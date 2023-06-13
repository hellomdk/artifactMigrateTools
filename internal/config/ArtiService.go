package config

import (
	"artifactMigrateTools/internal/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func NewArti() (*[]model.Arti, error) {
	c := new(Context)
	filePath := c.GetArtifactPath()
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("读取artifact.yaml文件失败: ", err)
	}

	repoList, err := parseArti(data)
	if err != nil {
		log.Println("解析artifact.yaml文件失败: ", err)
	}
	return repoList, nil
}

func parseArti(yamlData []byte) (*[]model.Arti, error) {
	var cfg struct {
		Arti []model.Arti `yaml:"artifact"`
	}

	err := yaml.Unmarshal(yamlData, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg.Arti, nil
}
