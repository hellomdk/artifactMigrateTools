package util

import (
	"testing"
)

func TestRepoCsvToYaml(t *testing.T) {
	fileName := "C:\\Users\\18638\\GolandProjects\\artifactMigrateTools\\conf\\csv\\repo.csv"
	distName := "C:\\Users\\18638\\GolandProjects\\artifactMigrateTools\\conf\\yaml\\repo.yaml"
	repoCsvList := ReadCsv(fileName)
	repoYaml := ConvertCsvDataToYaml(repoCsvList)
	WriteCSV(distName, repoYaml)
}

func TestNodeCsvToYaml(t *testing.T) {
	fileName := "C:\\Users\\18638\\GolandProjects\\artifactMigrateTools\\conf\\csv\\node.csv"
	distName := "C:\\Users\\18638\\GolandProjects\\artifactMigrateTools\\conf\\yaml\\artifact.yaml"
	nodeCsvList := ReadNodeCsv(fileName)
	nodeYaml := ConvertNodeCsvDataToYaml(nodeCsvList)
	WriteNodeCSV(distName, nodeYaml)
}
