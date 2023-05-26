package util

import (
	"encoding/csv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type RepoData struct {
	RepoKey          string
	RepositoriesList []RepositoriesYaml
}

type RepoCsv struct {
	RepoKey       string
	Description   string
	Layout        string
	ProtocolType  string
	Project       string
	ProxyRepoName string
	ProxyUrl      string
	ProxyUser     string
	ProxyPassword string
}

type RepoYaml struct {
	Repositories []RepositoriesYaml `yaml:"repositories"`
}

type RepositoriesYaml struct {
	OriginKey             string                     `yaml:"-"`
	RepoKey               string                     `yaml:"repoKey"`
	RepoType              string                     `yaml:"repoType"`
	Description           string                     `yaml:"description"`
	Layout                string                     `yaml:"layout"`
	ProtocolType          string                     `yaml:"protocolType"`
	Project               string                     `yaml:"project"`
	Url                   string                     `yaml:"url,omitempty"`
	Username              string                     `yaml:"username,omitempty"`
	Password              string                     `yaml:"password,omitempty"`
	DefaultDeploymentRepo string                     `yaml:"defaultDeploymentRepo,omitempty"`
	SelectedRepositories  []SelectedRepositoriesYaml `yaml:"selectedRepositories,omitempty"`
}

type SelectedRepositoriesYaml struct {
	RepoKey    string `yaml:"repoKey"`
	ProjectKey string `yaml:"projectKey"`
	Key        string `yaml:"key"`
	Type       string `yaml:"type"`
}

func ReadCsv(fileName string) []RepoCsv {
	csvFile, err := os.Open(fileName)
	if err != nil {
		log.Println("Open CSV fail:", err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	data, err := csvReader.ReadAll()

	csvList := CreateRepoCsvList(data)
	return csvList
}

func CreateRepoCsvList(data [][]string) []RepoCsv {
	var repoCsvList []RepoCsv
	for i, line := range data {
		if i >= 0 { // omit header line
			var rec RepoCsv
			for j, field := range line {
				if j == 0 {
					rec.RepoKey = field
				} else if j == 1 {
					rec.Description = field
				} else if j == 2 {
					rec.Layout = field
				} else if j == 3 {
					rec.ProtocolType = field
				} else if j == 4 {
					rec.Project = field
				} else if j == 5 {
					rec.ProxyRepoName = field
				} else if j == 6 {
					rec.ProxyUrl = field
				} else if j == 7 {
					rec.ProxyUser = field
				} else if j == 8 {
					rec.ProxyPassword = field
				}
			}
			repoCsvList = append(repoCsvList, rec)
		}
	}
	return repoCsvList
}

func ConvertCsvDataToYaml(repoCsvData []RepoCsv) RepoYaml {
	var repositoriesYamlList []RepositoriesYaml
	var repoDataList []RepoData

	// 初始化集合，判断是否存在
	var set map[string]struct{}
	set = make(map[string]struct{})

	// 设置本地仓库&远程仓库
	for _, repoCsv := range repoCsvData {
		var repoData RepoData
		// 虚拟仓库专用
		var repositoriesList []RepositoriesYaml

		localRepoList := GenerateLocalRepo(repoCsv, set)
		remoteRepoList := GenerateRemoteRepo(repoCsv, set)

		// 本地仓库存在，说明有多远程仓库，需要合并虚仓数据
		if localRepoList.RepoKey != "" {
			repositoriesYamlList = append(repositoriesYamlList, localRepoList)
			repositoriesList = append(repositoriesList, localRepoList)
		}

		if remoteRepoList.RepoKey != "" {
			repositoriesYamlList = append(repositoriesYamlList, remoteRepoList)
			repositoriesList = append(repositoriesList, remoteRepoList)
		}

		repoData.RepoKey = repoCsv.RepoKey
		repoData.RepositoriesList = repositoriesList
		repoDataList = append(repoDataList, repoData)
	}

	// 设置虚拟仓库
	//resMap := ListToMap(repositoriesYamlList, "OriginKey")
	// merger虚拟仓库数据
	repoDataList = MergeRepoDataList(repoDataList)
	for _, v := range repoDataList {
		//repoDataList := v.([]RepositoriesYaml)
		virtualRepoList := GenerateVirtualRepo(v, set)
		repositoriesYamlList = append(repositoriesYamlList, virtualRepoList)
	}

	var repoYaml RepoYaml
	repoYaml.Repositories = repositoriesYamlList
	return repoYaml
}

// 合并结果集
func MergeRepoDataList(repoDataList []RepoData) []RepoData {
	var newRepoDataList []RepoData
	// 初始化集合，判断是否存在
	var myMap map[string][]RepositoriesYaml
	myMap = make(map[string][]RepositoriesYaml)

	for i, repoData := range repoDataList {
		var newList []RepositoriesYaml
		if _, ok := myMap[repoData.RepoKey]; ok {
			tmpList := myMap[repoData.RepoKey]
			for _, repository := range repoData.RepositoriesList {
				tmpList = append(tmpList, repository)
			}
			newList = tmpList
			myMap[repoData.RepoKey] = newList
			repoData.RepositoriesList = newList
			newRepoDataList = append(newRepoDataList[i:], repoData)
		} else {
			newList = repoData.RepositoriesList
			myMap[repoData.RepoKey] = newList
			repoData.RepositoriesList = newList
			newRepoDataList = append(newRepoDataList, repoData)
		}

	}
	return newRepoDataList
}
func GenerateLocalRepo(repoCsv RepoCsv, set map[string]struct{}) RepositoriesYaml {
	var repositoriesYaml RepositoriesYaml
	key := repoCsv.RepoKey + "-local"
	if _, ok := set[key]; ok {
		return repositoriesYaml
	}
	set[key] = struct{}{}
	repositoriesYaml.OriginKey = repoCsv.RepoKey
	repositoriesYaml.RepoKey = key
	repositoriesYaml.RepoType = "local"
	repositoriesYaml.Description = repoCsv.Description
	repositoriesYaml.Layout = repoCsv.Layout
	repositoriesYaml.ProtocolType = repoCsv.ProtocolType
	repositoriesYaml.Project = repoCsv.Project
	return repositoriesYaml
}

func GenerateRemoteRepo(repoCsv RepoCsv, set map[string]struct{}) RepositoriesYaml {
	var repositoriesYaml RepositoriesYaml
	if repoCsv.ProxyUrl == "" {
		return repositoriesYaml
	}
	key := repoCsv.ProxyRepoName
	if _, ok := set[key]; ok {
		return repositoriesYaml
	}
	set[key] = struct{}{}

	repositoriesYaml.OriginKey = repoCsv.RepoKey
	repositoriesYaml.RepoKey = key
	repositoriesYaml.RepoType = "remote"
	repositoriesYaml.Description = repoCsv.Description
	repositoriesYaml.Layout = repoCsv.Layout
	repositoriesYaml.ProtocolType = repoCsv.ProtocolType
	repositoriesYaml.Project = repoCsv.Project

	repositoriesYaml.Url = repoCsv.ProxyUrl
	repositoriesYaml.Username = repoCsv.ProxyUser
	repositoriesYaml.Password = repoCsv.ProxyPassword
	return repositoriesYaml
}

func GenerateVirtualRepo(repoData RepoData, set map[string]struct{}) RepositoriesYaml {
	var repositoriesYaml RepositoriesYaml
	key := repoData.RepoKey
	if _, ok := set[key]; ok {
		return repositoriesYaml
	}

	firstData := repoData.RepositoriesList[0]

	repositoriesYaml.RepoKey = repoData.RepoKey
	repositoriesYaml.RepoType = "virtual"
	repositoriesYaml.Description = firstData.Description
	repositoriesYaml.Layout = firstData.Layout
	repositoriesYaml.ProtocolType = firstData.ProtocolType
	repositoriesYaml.Project = firstData.Project

	repositoriesYaml.DefaultDeploymentRepo = repoData.RepoKey + "-local"
	var selectedList []SelectedRepositoriesYaml
	for _, data := range repoData.RepositoriesList {
		var selected SelectedRepositoriesYaml
		selected.RepoKey = data.RepoKey
		selected.ProjectKey = data.Project
		selected.Key = data.RepoKey
		selected.Type = data.RepoType
		selectedList = append(selectedList, selected)
	}

	repositoriesYaml.SelectedRepositories = selectedList
	return repositoriesYaml
}

func WriteCSV(name string, contents RepoYaml) error {

	data, err := yaml.Marshal(contents) // 第二个表示每行的前缀，这里不用，第三个是缩进符号，这里用tab
	checkError(err)
	err = ioutil.WriteFile(name, data, 0777)
	checkError(err)
	return nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
