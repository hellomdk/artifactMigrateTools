package config

import (
	"artifactMigrateTools/internal/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	_ "os"
)

func NewConfig() (*model.Config, error) {
	c := new(Context)
	filePath := c.GetConfigPath()

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("读取config.yaml文件失败: ", err)
	}

	config, err := parseConfig(data)
	if err != nil {
		log.Println("解析config.yaml文件失败: ", err)
	}
	return config, nil
}

func parseConfig(yamlData []byte) (*model.Config, error) {
	var cfg struct {
		Config model.Config `yaml:"config"`
	}

	err := yaml.Unmarshal(yamlData, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg.Config, nil
}
