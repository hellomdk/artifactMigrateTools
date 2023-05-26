package repo

import (
	"encoding/json"
	"fmt"
	"jfrogToArtifact/pkg/artifact/httpclient"
	"log"
	"strings"
	"time"
)

type Repo struct {
	HttpClient httpclient.HttpClient
}

type Repository struct {
	RepoKey               string                 `json:"repoKey"`
	Description           string                 `json:"description"`
	RepoType              string                 `json:"repoType"`
	Layout                string                 `json:"layout"`
	Browse                bool                   `json:"browse"`
	ProtocolSpecific      ProtocolType           `json:"protocolSpecific"`
	Url                   string                 `json:"url"`
	Username              string                 `json:"username"`
	Password              string                 `json:"password"`
	ProjectKey            string                 `json:"projectKey"`
	SelectedRepositories  []SelectedRepositories `json:"selectedRepositories"`
	ResolvedRepositories  []SelectedRepositories `json:"resolvedRepositories"`
	DefaultDeploymentRepo string                 `json:"defaultDeploymentRepo"`
}

type SelectedRepositories struct {
	RepoKey    string `json:"repoKey"yaml:"repoKey"`
	ProjectKey string `json:"projectKey"yaml:"projectKey"`
	Key        string `json:"key"yaml:"key"`
	Type       string `json:"type"yaml:"type"`
}

type ProtocolType struct {
	ProtocolType string `json:"protocolType"`
	DockerApiVersion string `json:"dockerApiVersion"`
}
type Artifact struct {
	Repo         string    `json:"repo"`
	Path         string    `json:"path"`
	Created      time.Time `json:"created"`
	CreatedBy    string    `json:"createdBy"`
	LastModified time.Time `json:"lastModified"`
	ModifiedBy   string    `json:"modifiedBy"`
	LastUpdated  time.Time `json:"lastUpdated"`
	DownloadUri  string    `json:"downloadUri"`
	MimeType     string    `json:"mimeType"`
	Size         string    `json:"size"`
	Checksums    struct {
		Sha1   string `json:"sha1"`
		Md5    string `json:"md5"`
		Sha256 string `json:"sha256"`
	} `json:"checksums"`
	OriginalChecksums struct {
		Sha1   string `json:"sha1"`
		Md5    string `json:"md5"`
		Sha256 string `json:"sha256"`
	} `json:"originalChecksums"`
	Uri string `json:"uri"`
}

type Prop struct {
	Properties struct {
		Age  []string `json:"age"`
		Name []string `json:"name"`
	} `json:"properties"`
	URI string `json:"uri"`
}

/**
探活
client.Ping()
*/
func (r *Repo) Ping() (string, error) {
	url := "/api/monitor/healthy"
	_, body, err := httpclient.Client.GetString(r.HttpClient.BaseURL+url, r.HttpClient.Header, nil)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

/**
获取仓库详情，判断仓库是否存在
client.ExistRepository("Gradle_Test")
*/
func (r *Repo) ExistRepository(repoKey string) bool {
	url := "/open/api/v1/repositories/{repoKey}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	var cacheConfig = &Repository{}

	_, err := httpclient.Client.Get(r.HttpClient.BaseURL+url, r.HttpClient.Header, nil, cacheConfig)
	if err != nil {
		log.Fatal(err)
	}
	if cacheConfig.RepoKey != "" {
		return true
	}
	return false
}

/**
制品是否存在
client.ExistArtifact("Gradle_Test", "maven-metadata-local.xml")
*/
func (r *Repo) ExistArtifact(repoKey, repoPath string) bool {
	url := "/open/api/v1/storage/file/{repoKey}/{repoPath}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	url = strings.Replace(url, "{repoPath}", repoPath, 1)

	var cacheConfig = &Artifact{}
	_, err := httpclient.Client.Get(r.HttpClient.BaseURL+url, r.HttpClient.Header, nil, cacheConfig)
	if err != nil {
		log.Fatal(err)
	}
	if cacheConfig.Repo != "" {
		return true
	}
	return false
}

/**
创建仓库
client.CreateRepository("Gradle_Test")
*/
func (r *Repo) CreateRepository(repo Repository) bool {
	url := "/open/api/v1/repositories"

	var cacheConfig = &Repository{}
	//将切片进行序列化
	data, err := json.Marshal(repo)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化之后的结果
	fmt.Printf("序列化后=%v\n", string(data))

	statusCode, err := httpclient.Client.Post(r.HttpClient.BaseURL+url, r.HttpClient.Header, repo, cacheConfig)
	if err != nil {
		log.Fatal(err)
	}
	if statusCode == 200 {
		return true
	}
	return false
}

/**
创建制品
client.CreateArtifact("Gradle_Test")
*/
func (r *Repo) CreateArtifact(repoKey, repoPath string) bool {
	//url := "/repository/{repoKey}/{repoPath}"
	url := "/repository/{repoKey}/{repoPath}"
	//url := "/open/api/v1/storage/deploy/{repoKey}/{repoPath}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	url = strings.Replace(url, "{repoPath}", repoPath, 1)

	var cacheConfig = &Artifact{}
	statusCode, err := httpclient.Client.Put(r.HttpClient.BaseURL+url, r.HttpClient.Header, "application/octet-stream", nil, cacheConfig)

	if err != nil {
		log.Fatal(err)
	}
	if statusCode == 201 {
		return true
	}
	return false
}

/**
更新制品checksum
client.UpdateArtifactCheckSum("Gradle_Test")
*/
func (r *Repo) UpdateArtifactCheckSum(repoKey, repoPath, sha1, sha256, md5 string) bool {
	url := "/open/api/v1/storage/checksum/{repoKey}/{repoPath}?sha1={sha1}&sha256={sha256}&md5={md5}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	url = strings.Replace(url, "{repoPath}", repoPath, 1)
	url = strings.Replace(url, "{sha1}", sha1, 1)
	url = strings.Replace(url, "{sha256}", sha256, 1)
	url = strings.Replace(url, "{md5}", md5, 1)

	statusCode, err := httpclient.Client.Post(r.HttpClient.BaseURL+url, r.HttpClient.Header, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if statusCode == 200 {
		return true
	}
	return false
}

/**
更新制品checksum
client.UpdateArtifactCheckSum("Gradle_Test")
*/
func (r *Repo) UpdateArtifactProperties(repoKey, repoPath, properties string) bool {
	url := "/open/api/v1/storage/properties/{repoKey}/{repoPath}?properties={properties}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	url = strings.Replace(url, "{repoPath}", repoPath, 1)
	url = strings.Replace(url, "{properties}", properties, 1)

	statusCode, err := httpclient.Client.Put(r.HttpClient.BaseURL+url, r.HttpClient.Header, "application/json", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if statusCode == 200 {
		return true
	}
	return false
}
