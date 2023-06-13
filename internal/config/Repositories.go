package config

import (
	"artifactMigrateTools/internal/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func NewRepositories() (*[]model.Repositories, error) {
	c := new(Context)
	filePath := c.GetRepoPath()
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("读取repo.yaml文件失败: ", err)
	}

	repoList, err := parseRepositories(data)
	if err != nil {
		log.Println("解析repo.yaml文件失败: ", err)
	}
	return repoList, nil
}

func parseRepositories(yamlData []byte) (*[]model.Repositories, error) {
	var cfg struct {
		Repositories []model.Repositories `yaml:"repositories"`
	}

	err := yaml.Unmarshal(yamlData, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg.Repositories, nil
}
