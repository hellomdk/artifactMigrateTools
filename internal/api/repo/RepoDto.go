package repo

import (
	"artifactMigrateTools/internal/util"
)

type Repo struct {
	HttpClient util.HttpClient
}

type RepoRepository struct {
	RepoKey               string                     `json:"repoKey"`
	Description           string                     `json:"description"`
	RepoType              string                     `json:"repoType"`
	Layout                string                     `json:"layout"`
	Browse                bool                       `json:"browse"`
	ProtocolSpecific      RepoProtocolType           `json:"protocolSpecific"`
	Url                   string                     `json:"url"`
	Username              string                     `json:"username"`
	Password              string                     `json:"password"`
	ProjectKey            string                     `json:"projectKey"`
	SelectedRepositories  []RepoSelectedRepositories `json:"selectedRepositories"`
	ResolvedRepositories  []RepoSelectedRepositories `json:"resolvedRepositories"`
	DefaultDeploymentRepo string                     `json:"defaultDeploymentRepo"`
}

type RepoSelectedRepositories struct {
	RepoKey    string `json:"repoKey"yaml:"repoKey"`
	ProjectKey string `json:"projectKey"yaml:"projectKey"`
	Key        string `json:"key"yaml:"key"`
	Type       string `json:"type"yaml:"type"`
}

type RepoProtocolType struct {
	ProtocolType     string `json:"protocolType"`
	DockerApiVersion string `json:"dockerApiVersion"`
}
type RepoArtifact struct {
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

type RepoProp struct {
	Properties struct {
		Age  []string `json:"age"`
		Name []string `json:"name"`
	} `json:"properties"`
	URI string `json:"uri"`
}

type RepoProject struct {
	ProjectKey  string `json:"projectKey"`
	DisplayName string `json:"displayName"`
}
