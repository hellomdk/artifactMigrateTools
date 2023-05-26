package jfrog

import (
	"fmt"
	"io"
	"jfrogToArtifact/pkg/artifact/config"
	"jfrogToArtifact/pkg/artifact/httpclient"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Jfrog struct {
	HttpClient httpclient.HttpClient
}
type Repository struct {
	Key         string
	Description string
	Url         string
	PackageType string
	Type        string
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
func (jf *Jfrog) Ping() (string, error) {
	url := "/api/system/ping"
	_, body, err := httpclient.Client.GetString(jf.HttpClient.BaseURL+url, jf.HttpClient.Header, nil)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

/**
获取仓库列表
client.GetRepositories()
*/
func (jf *Jfrog) GetRepositories(config *config.Config) ([]Repository, error) {
	url := "/api/repositories"
	var cacheConfig = &[]Repository{}
	_, err := httpclient.Client.Get(jf.HttpClient.BaseURL+url, jf.HttpClient.Header, nil, cacheConfig)
	if err != nil {
		log.Fatal(err)
	}
	return *cacheConfig, nil
}

/**
获取仓库详情
client.GetRepository("Gradle_Test")
*/
func (jf *Jfrog) GetRepository(repoKey string) (Repository, error) {
	url := "/api/repositories/{repoKey}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)

	var cacheConfig = &Repository{}
	_, err := httpclient.Client.Get(jf.HttpClient.BaseURL+url, jf.HttpClient.Header, nil, cacheConfig)
	if err != nil {
		log.Fatal(err)
	}
	return *cacheConfig, nil
}

/**
获取制品详情
client.GetArtifact("Gradle_Test", "maven-metadata-local.xml")
*/
func (jf *Jfrog) GetArtifact(repoKey, repoPath string) (Artifact, error) {
	url := "/api/storage/{repoKey}/{repoPath}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	url = strings.Replace(url, "{repoPath}", repoPath, 1)

	var cacheConfig = &Artifact{}
	_, err := httpclient.Client.Get(jf.HttpClient.BaseURL+url, jf.HttpClient.Header, nil, cacheConfig)
	if err != nil {
		log.Fatal(err)
	}
	return *cacheConfig, nil
}

/**
获取制品属性
client.GetArtifactProperties("Gradle_Test", "maven-metadata-local.xml")
*/
func (jf *Jfrog) GetArtifactProperties(repoKey, repoPath string) (Prop, error) {
	url := "/api/storage/{repoKey}/{repoPath}?properties"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	url = strings.Replace(url, "{repoPath}", repoPath, 1)

	var cacheConfig = &Prop{}
	_, err := httpclient.Client.Get(jf.HttpClient.BaseURL+url, jf.HttpClient.Header, nil, cacheConfig)
	if err != nil {
		log.Fatal(err)
	}
	return *cacheConfig, nil
}

/**
制品下载
client.DownloadFile("http://192.168.80.97:8082/artifactory/api/storage/Gradle_Test/maven-metadata-local.xml", "./maven-metadata-local.xml")
*/
func (jf *Jfrog) DownloadFile(url string, filepath string) error {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 创建一个HTTP请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("创建HTTP请求失败：%s", err)
	}

	// 发送HTTP请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送HTTP请求失败：%s", err)
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存下载的文件
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("创建文件失败：%s", err)
	}
	defer file.Close()

	// 将响应体写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("写入文件失败：%s", err)
	}

	return nil
}
