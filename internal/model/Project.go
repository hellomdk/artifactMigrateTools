package model

type ProjectYaml struct {
	Project []Project `yaml:"project"`
}

type Project struct {
	ProjectKey        string `yaml:"project"`
	ProjectKeyMapping string `yaml:"projectMapping"`
	Migrated          bool   `yaml:"migrated"`
}
