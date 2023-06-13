package model

type NodeYaml struct {
	Artifact []Arti `yaml:"artifact"`
}

type Arti struct {
	Name         string `yaml:"name"`
	Repo         string `yaml:"repo"`
	OriginRepo   string `yaml:"originRepo"`
	RepoMapping  string `yaml:"repoMapping"`
	VirtualRepo  string `yaml:"virtualRepo"` // 虚仓repoKey
	ProtocolType string `yaml:"protocolType"`
	Path         string `yaml:"path"`
	Sha1         string `yaml:"sha1"`
	Sha256       string `yaml:"sha256"`
	Md5          string `yaml:"md5"`
	Migrated     bool   `yaml:"migrated"`
}
