package nexus

import (
	"artifactMigrateTools/internal/util"
)

type Nexus struct {
	HttpClient util.HttpClient
}

type NexusDockerManifests struct {
	SchemaVersion int             `json:"schemaVersion"`
	Name          string          `json:"name"`
	Tag           string          `json:"tag"`
	Architecture  string          `json:"architecture"`
	FsLayers      []NexusFsLayers `json:"fsLayers"`
	History       []NexusHistory  `json:"history"`
}
type NexusFsLayers struct {
	BlobSum string `json:"blobSum"`
}
type NexusHistory struct {
	V1Compatibility string `json:"v1Compatibility"`
}

type NexusFileItem struct {
	Items             []NexusItems `json:"items"`
	ContinuationToken interface{}  `json:"continuationToken"`
}

// item数据结构
type NexusItems struct {
	ID         string        `json:"id"`
	Repository string        `json:"repository"`
	Format     string        `json:"format"`
	Group      interface{}   `json:"group"`
	Name       string        `json:"name"`
	Version    string        `json:"version"`
	Assets     []NexusAssets `json:"assets"`
}

type NexusAssets struct {
	DownloadURL    string        `json:"downloadUrl"`
	Path           string        `json:"path"`
	ID             string        `json:"id"`
	Repository     string        `json:"repository"`
	Format         string        `json:"format"`
	Checksum       NexusChecksum `json:"checksum"`
	ContentType    string        `json:"contentType"`
	LastModified   string        `json:"lastModified"`
	BlobCreated    string        `json:"blobCreated"`
	LastDownloaded string        `json:"lastDownloaded"`
}

type NexusChecksum struct {
	Sha1   string `json:"sha1"`
	Sha256 string `json:"sha256"`
	Sha512 string `json:"sha512"`
	Md5    string `json:"md5"`
}
