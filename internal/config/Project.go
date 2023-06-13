package config

import (
	"artifactMigrateTools/internal/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func NewProject() ([]model.Project, error) {
	c := new(Context)
	filePath := c.GetProjectPath()
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("读取project.yaml文件失败: ", err)
	}

	projectList, err := parseProject(data)
	if err != nil {
		log.Println("解析project.yaml文件失败: ", err)
	}
	return *projectList, nil
}

func parseProject(yamlData []byte) (*[]model.Project, error) {
	var cfg struct {
		Project []model.Project `yaml:"project"`
	}

	err := yaml.Unmarshal(yamlData, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg.Project, nil
}
