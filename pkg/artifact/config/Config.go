package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	_ "os"
)

type Config struct {
	WorkDir      string `yaml:"workDir"`
	Threads      int    `yaml:"threads"`
	Logging      string `yaml:"logging"`
	ChecksumCalc bool   `yaml:"checksumCalc"`
	SourceRepo   Repo   `yaml:"sourceRepo"`
	TargetRepo   Repo   `yaml:"targetRepo"`
}

type Repo struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func NewConfig() (*Config, error) {
	data, err := ioutil.ReadFile("C:\\Users\\18638\\GolandProjects\\jfrogToArtifact\\conf\\yaml\\config.yaml")
	if err != nil {
		panic(err)
	}

	config, err := parseConfig(data)
	if err != nil {
		panic(err)
	}
	return config, nil
}

func parseConfig(yamlData []byte) (*Config, error) {
	var cfg struct {
		Config Config `yaml:"config"`
	}

	err := yaml.Unmarshal(yamlData, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg.Config, nil
}
