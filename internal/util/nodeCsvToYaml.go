package util

import (
	"encoding/csv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type NodeCsv struct {
	Name         string
	Repo         string
	OriginRepo   string
	ProtocolType string
	Path         string
	Sha1         string
	Sha256       string
	Md5          string
}

type NodeYaml struct {
	Artifact []ArtifactYaml `yaml:"artifact"`
}

type ArtifactYaml struct {
	Name         string `yaml:"name"`
	Repo         string `yaml:"repo"`
	OriginRepo   string `yaml:"originRepo"`
	ProtocolType string `yaml:"protocolType"`
	Path         string `yaml:"path"`
	Sha1         string `yaml:"sha1"`
	Sha256       string `yaml:"sha256"`
	Md5          string `yaml:"md5"`
	Migrated     bool   `yaml:"migrated"`
}

func ReadNodeCsv(fileName string) []NodeCsv {
	csvFile, err := os.Open(fileName)
	if err != nil {
		log.Println("Open CSV fail:", err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	data, err := csvReader.ReadAll()

	csvList := CreateNodeCsvList(data)
	return csvList
}

func CreateNodeCsvList(data [][]string) []NodeCsv {
	var nodeCsvList []NodeCsv
	for i, line := range data {
		if i >= 0 { // omit header line
			var rec NodeCsv
			for j, field := range line {
				if j == 0 {
					rec.Name = field
				} else if j == 1 {
					rec.Repo = field
				} else if j == 2 {
					rec.OriginRepo = field
				} else if j == 3 {
					rec.ProtocolType = field
				} else if j == 4 {
					rec.Path = field
				} else if j == 5 {
					rec.Sha1 = field
				} else if j == 6 {
					rec.Sha256 = field
				} else if j == 7 {
					rec.Md5 = field
				}
			}
			nodeCsvList = append(nodeCsvList, rec)
		}
	}
	return nodeCsvList
}

func ConvertNodeCsvDataToYaml(nodeCsvData []NodeCsv) NodeYaml {
	var nodeYaml NodeYaml
	var artifactYamlList []ArtifactYaml

	// 设置本地仓库&远程仓库
	for _, nodeCsv := range nodeCsvData {
		artifact := GenerateArtifact(nodeCsv)
		if artifact.Path != "" {
			artifactYamlList = append(artifactYamlList, artifact)
		}
	}

	nodeYaml.Artifact = artifactYamlList
	return nodeYaml
}

func GenerateArtifact(nodeCsvData NodeCsv) ArtifactYaml {
	var artifactYaml ArtifactYaml
	artifactYaml.Name = nodeCsvData.Name
	artifactYaml.Repo = nodeCsvData.Repo
	artifactYaml.OriginRepo = nodeCsvData.OriginRepo
	artifactYaml.ProtocolType = nodeCsvData.ProtocolType
	artifactYaml.Path = nodeCsvData.Path
	artifactYaml.Sha1 = nodeCsvData.Sha1
	artifactYaml.Sha256 = nodeCsvData.Sha256
	artifactYaml.Md5 = nodeCsvData.Md5
	artifactYaml.Migrated = false
	return artifactYaml
}

func WriteNodeCSV(name string, contents NodeYaml) error {

	data, err := yaml.Marshal(contents) // 第二个表示每行的前缀，这里不用，第三个是缩进符号，这里用tab
	checkError(err)
	err = ioutil.WriteFile(name, data, 0777)
	checkError(err)
	return nil
}
