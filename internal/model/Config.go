package model

import (
	_ "os"
)

type ConfigYaml struct {
	Config Config `yaml:"config"`
}
type Config struct {
	WorkDir      string `yaml:"workDir"`
	Threads      int    `yaml:"threads"`
	Logging      string `yaml:"logging"`
	ChecksumCalc bool   `yaml:"checksumCalc"`
	SourceRepo   Repo   `yaml:"sourceRepo"`
	TargetRepo   Repo   `yaml:"targetRepo"`
}

type Repo struct {
	Id       string `yaml:"id"`
	URL      string `yaml:"url"`
	Type     string `yaml:"type"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
