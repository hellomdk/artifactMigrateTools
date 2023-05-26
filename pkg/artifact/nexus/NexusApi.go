package nexus

import (
	"jfrogToArtifact/pkg/artifact/httpclient"
	"log"
	"strings"
	"time"
)

type Nexus struct {
	HttpClient httpclient.HttpClient
}

type DockerManifests struct {
	SchemaVersion int        `json:"schemaVersion"`
	Name          string     `json:"name"`
	Tag           string     `json:"tag"`
	Architecture  string     `json:"architecture"`
	FsLayers      []FsLayers `json:"fsLayers"`
	History       []History  `json:"history"`
}
type FsLayers struct {
	BlobSum string `json:"blobSum"`
}
type History struct {
	V1Compatibility string `json:"v1Compatibility"`
}

type FileItem struct {
	Items             []Items     `json:"items"`
	ContinuationToken interface{} `json:"continuationToken"`
}

// item数据结构
type Items struct {
	ID         string      `json:"id"`
	Repository string      `json:"repository"`
	Format     string      `json:"format"`
	Group      interface{} `json:"group"`
	Name       string      `json:"name"`
	Version    string      `json:"version"`
	Assets     []Assets    `json:"assets"`
}

type Assets struct {
	DownloadURL    string    `json:"downloadUrl"`
	Path           string    `json:"path"`
	ID             string    `json:"id"`
	Repository     string    `json:"repository"`
	Format         string    `json:"format"`
	Checksum       Checksum  `json:"checksum"`
	ContentType    string    `json:"contentType"`
	LastModified   time.Time `json:"lastModified"`
	BlobCreated    time.Time `json:"blobCreated"`
	LastDownloaded time.Time `json:"lastDownloaded"`
}

type Checksum struct {
	Sha1   string `json:"sha1"`
	Sha256 string `json:"sha256"`
}

/**
获取bucket 文件列表
client.CreateArtifact("Gradle_Test")
*/
func (n *Nexus) GetItemList(repoKey string) FileItem {
	url := "/service/rest/v1/components?repository={repoKey}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)

	var cacheConfig = &FileItem{}
	_, err := httpclient.Client.Get(n.HttpClient.BaseURL+url, n.HttpClient.Header, nil, cacheConfig)

	if err != nil {
		log.Fatal(err)
	}

	return *cacheConfig
}

/**
读取manifests
client.CreateArtifact("Gradle_Test")
*/
func (n *Nexus) ReadManifests(repoKey, repoPath string) (DockerManifests, string) {
	url := "/repository/{repoKey}/{repoPath}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	url = strings.Replace(url, "{repoPath}", repoPath, 1)

	var cacheConfig = &DockerManifests{}
	_, err := httpclient.Client.Get(n.HttpClient.BaseURL+url, n.HttpClient.Header, nil, cacheConfig)
	_, str, err := httpclient.Client.GetString(n.HttpClient.BaseURL+url, n.HttpClient.Header, nil)

	if err != nil {
		log.Fatal(err)
	}

	return *cacheConfig, str
}
