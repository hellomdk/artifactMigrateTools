package migrate

import (
	"testing"
)

// 写入project.yaml
func TestNexusRepoMigrate_GenericNexusProjectYaml(t *testing.T) {
	ng := new(NexusGeneric)
	ng.GenericNexusProjectYaml()
}

// 写入repo.yaml
func TestNexusRepoMigrate_GenericNexusRepoYaml(t *testing.T) {
	ng := new(NexusGeneric)
	ng.GenericNexusRepoYaml()
}

// 写入Arti.yaml
func TestNexusRepoMigrate_GenericNexusArtiYaml(t *testing.T) {
	ng := new(NexusGeneric)
	ng.GenericNexusArtiYaml()
}
