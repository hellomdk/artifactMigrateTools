package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Arti struct {
	Name         string `yaml:"name"`
	Repo         string `yaml:"repo"`
	RepoOrg         string `yaml:"repoOrg"`
	RepoParent         string `yaml:"repoParent"`
	ProtocolType string `yaml:"protocolType"`
	Path         string `yaml:"path"`
	Sha1         string `yaml:"sha1"`
	Sha256       string `yaml:"sha256"`
	Md5          string `yaml:"md5"`
}

func NewArti() (*[]Arti, error) {
	data, err := ioutil.ReadFile("C:\\Users\\18638\\GolandProjects\\jfrogToArtifact\\conf\\yaml\\node.yaml")
	if err != nil {
		panic(err)
	}

	repoList, err := parseArti(data)
	if err != nil {
		panic(err)
	}
	return repoList, nil
}

func parseArti(yamlData []byte) (*[]Arti, error) {
	var cfg struct {
		Arti []Arti `yaml:"artifact"`
	}

	err := yaml.Unmarshal(yamlData, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg.Arti, nil
}
