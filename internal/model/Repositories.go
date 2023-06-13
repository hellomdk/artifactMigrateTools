package model

type RepositoriesYaml struct {
	Repositories []Repositories `yaml:"repositories"`
}

type Repositories struct {
	RepoKey               string                 `yaml:"repoKey"`
	RepoKeyMapping        string                 `yaml:"repoKeyMapping"`
	RepoType              string                 `yaml:"repoType"`
	Description           string                 `yaml:"description"`
	Layout                string                 `yaml:"layout"`
	ProtocolType          string                 `yaml:"protocolType"`
	Url                   string                 `yaml:"url"`
	Username              string                 `yaml:"username"`
	Password              string                 `yaml:"password"`
	ProjectKey            string                 `yaml:"project"`
	ProjectKeyMapping     string                 `yaml:"-"`
	SelectedRepositories  []SelectedRepositories `yaml:"selectedRepositories"`
	DefaultDeploymentRepo string                 `yaml:"defaultDeploymentRepo"`
	Migrated              bool                   `yaml:"migrated"`
}

type SelectedRepositories struct {
	RepoKey    string `yaml:"repoKey"`
	ProjectKey string `yaml:"projectKey"`
	Key        string `yaml:"key"`
	Type       string `yaml:"type"`
}
