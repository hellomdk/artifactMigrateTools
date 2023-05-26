package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"jfrogToArtifact/pkg/artifact/repo"
)

type Repositories struct {
	RepoKey               string                 `yaml:"repoKey"`
	RepoType              string                 `yaml:"repoType"`
	Description           string                 `yaml:"description"`
	Layout                string                 `yaml:"layout"`
	ProtocolType          string                 `yaml:"protocolType"`
	Url                   string                 `yaml:"url"`
	Username              string                 `yaml:"username"`
	Password              string                 `yaml:"password"`
	ProjectKey            string                 `yaml:"project"`
	SelectedRepositories  []repo.SelectedRepositories `yaml:"selectedRepositories"`
	DefaultDeploymentRepo string                 `yaml:"defaultDeploymentRepo"`
}

//type SelectedRepositories struct {
//	RepoKey    string `yaml:"repoKey"`
//	ProjectKey string `yaml:"projectKey"`
//	Key        string `yaml:"key"`
//	Type       string `yaml:"type"`
//}

func NewRepositories() (*[]Repositories, error) {
	data, err := ioutil.ReadFile("C:\\Users\\18638\\GolandProjects\\jfrogToArtifact\\conf\\yaml\\repo.yaml")
	if err != nil {
		panic(err)
	}

	repoList, err := parseRepositories(data)
	if err != nil {
		panic(err)
	}
	return repoList, nil
}

func parseRepositories(yamlData []byte) (*[]Repositories, error) {
	var cfg struct {
		Repositories []Repositories `yaml:"repositories"`
	}

	err := yaml.Unmarshal(yamlData, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg.Repositories, nil
}
