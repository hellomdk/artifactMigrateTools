package jfrog

import (
	"artifactMigrateTools/internal/util"
)

type Jfrog struct {
	HttpClient util.HttpClient
}
type JfrogRepository struct {
	Key         string
	Description string
	Url         string
	PackageType string
	Type        string
}

type JfrogRepositoryDetail struct {
	Key                   string
	PackageType           string
	Description           string
	RepoLayoutRef         string   // layout列表
	Repositories          []string // 虚拟仓库列表
	DefaultDeploymentRepo string   // 虚拟仓库默认仓库
	Rclass                string   // 仓库类型（local、remote、virtual）
	Url                   string   // 远程仓库地址
	Username              string   // 远程仓库用户名
	Password              string   // 远程仓库密码
	ProjectKey            string   // 所属空间
}

type JfrogArtifacts struct {
	Uri     string
	Created string
	Files   []FileItem
}

type FileItem struct {
	Uri          string
	Size         int
	LastModified string
	Folder       bool
	Sha1         string
	Sha2         string
	Md5          string
}
type JfrogArtifact struct {
	Repo         string `json:"repo"`
	Path         string `json:"path"`
	Created      string `json:"created"`
	CreatedBy    string `json:"createdBy"`
	LastModified string `json:"lastModified"`
	ModifiedBy   string `json:"modifiedBy"`
	LastUpdated  string `json:"lastUpdated"`
	DownloadUri  string `json:"downloadUri"`
	MimeType     string `json:"mimeType"`
	Size         string `json:"size"`
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

type JfrogProp struct {
	Properties struct {
		Age  []string `json:"age"`
		Name []string `json:"name"`
	} `json:"properties"`
	URI string `json:"uri"`
}
