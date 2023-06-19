package repo

import (
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/util"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

/**
探活
client.Ping()
*/
func (r *Repo) Ping(context *config.Context) (string, error) {
	url := "/api/monitor/healthy"
	_, body, err := util.Client.GetString(r.HttpClient.BaseURL+url, r.HttpClient.Header, nil)
	if err != nil {
		context.Loggers.SendLoggerError("API测试连接失败: ", err)
	}
	return body, nil
}

/**
获取仓库详情，判断仓库是否存在
client.ExistRepository("Gradle_Test")
*/
func (r *Repo) ExistRepository(context *config.Context, repoKey string) bool {
	url := "/open/api/v1/repositories/{repoKey}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	var cacheConfig = &RepoRepository{}

	statusCode, err := util.Client.Get(r.HttpClient.BaseURL+url, r.HttpClient.Header, nil, cacheConfig)
	if statusCode == 404 {
		return false
	}
	if err != nil {
		context.Loggers.SendLoggerError("API判断仓库是否存在失败: ", err)
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
func (r *Repo) ExistArtifact(context *config.Context, repoKey, repoPath string) bool {
	url := "/open/api/v1/storage/file/{repoKey}/{repoPath}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	url = strings.Replace(url, "{repoPath}", repoPath, 1)

	var cacheConfig = &RepoArtifact{}
	_, err := util.Client.Get(r.HttpClient.BaseURL+url, r.HttpClient.Header, nil, cacheConfig)
	if err != nil {
		context.Loggers.SendLoggerError("API判断制品是否存在失败: ", err)
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
func (r *Repo) CreateRepository(context *config.Context, repo RepoRepository) bool {
	url := "/open/api/v1/repositories"

	var cacheConfig = &RepoRepository{}
	//将切片进行序列化
	data, err := json.Marshal(repo)
	if err != nil {
		context.Loggers.SendLoggerError("API创建仓库参数序列化失败: ", err)
	}
	//输出序列化之后的结果
	fmt.Printf("序列化后=%v\n", string(data))

	statusCode, err := util.Client.Post(r.HttpClient.BaseURL+url, r.HttpClient.Header, repo, cacheConfig)
	if err != nil {
		context.Loggers.SendLoggerError("API创建仓库失败: ", err)
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
func (r *Repo) CreateArtifact(context *config.Context, repoKey, repoPath string) bool {
	//url := "/repository/{repoKey}/{repoPath}"
	url := "/{repoKey}/{repoPath}"
	//url := "/open/api/v1/storage/deploy/{repoKey}/{repoPath}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	url = strings.Replace(url, "{repoPath}", repoPath, 1)

	var cacheConfig = &RepoArtifact{}
	statusCode, err := util.Client.Put(r.HttpClient.BaseURL+url, r.HttpClient.Header, "application/octet-stream", nil, cacheConfig)

	if err != nil {
		context.Loggers.SendLoggerError("API发布制品失败: ", err)
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
func (r *Repo) UpdateArtifactCheckSum(context *config.Context, repoKey, repoPath, sha1, sha256, md5 string) bool {
	url := "/open/api/v1/storage/checksum/{repoKey}/{repoPath}?sha1={sha1}&sha256={sha256}&md5={md5}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	url = strings.Replace(url, "{repoPath}", repoPath, 1)
	url = strings.Replace(url, "{sha1}", sha1, 1)
	url = strings.Replace(url, "{sha256}", sha256, 1)
	url = strings.Replace(url, "{md5}", md5, 1)

	statusCode, err := util.Client.Post(r.HttpClient.BaseURL+url, r.HttpClient.Header, nil, nil)
	if err != nil {
		context.Loggers.SendLoggerError("API更新制品checksum失败: ", err)
	}
	if statusCode == 200 {
		return true
	}
	return false
}

/**
更新制品属性
*/
func (r *Repo) UpdateArtifactProperties(context *config.Context, repoKey, repoPath, properties string) bool {
	urlstr := "/open/api/v1/storage/properties/{repoKey}/{repoPath}?propertySourceType=${propertySourceType}&&properties={properties}"
	urlstr = strings.Replace(urlstr, "{repoKey}", repoKey, 1)
	urlstr = strings.Replace(urlstr, "{repoPath}", repoPath, 1)
	urlstr = strings.Replace(urlstr, "{propertySourceType}", "oldRepo", 1)
	urlstr = strings.Replace(urlstr, "{properties}", url.QueryEscape(properties), 1)

	//urlstr = url.QueryEscape(urlstr)

	statusCode, err := util.Client.Put(r.HttpClient.BaseURL+urlstr, r.HttpClient.Header, "application/json", nil, nil)
	if err != nil {
		context.Loggers.SendLoggerError("API更新制品属性失败: ", err)
	}
	if statusCode == 200 {
		return true
	}
	return false
}

/**
创建空间
*/
func (r *Repo) CreateProject(context *config.Context, project string) bool {
	url := "/ui/project"
	body := RepoProject{
		ProjectKey:  project,
		DisplayName: project,
	}
	statusCode, err := util.Client.Post(r.HttpClient.BaseURL+url, r.HttpClient.Header, body, nil)
	if err != nil {
		context.Loggers.SendLoggerError("API创建空间失败: ", err)
	}
	if statusCode == 200 {
		return true
	}
	return false
}

/**
创建空间
*/
func (r *Repo) ExistProject(context *config.Context, projectKey string) bool {
	url := "/ui/project/{projectKey}"
	url = strings.Replace(url, "{projectKey}", projectKey, 1)

	statusCode, err := util.Client.Get(r.HttpClient.BaseURL+url, r.HttpClient.Header, nil, nil)
	if statusCode == 404 {
		return false
	}
	if err != nil {
		context.Loggers.SendLoggerError("API创建空间失败: ", err)
	}
	if statusCode == 200 {
		return true
	}
	return false
}

/**
索引刷新
*/
func (r *Repo) ReIndex(context *config.Context, repoKey string) bool {
	url := "/ui/tree-browser/artifact/action/calculateIndex/{repoKey}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)

	statusCode, err := util.Client.Post(r.HttpClient.BaseURL+url, r.HttpClient.Header, nil, nil)
	if statusCode == 404 {
		return false
	}
	if err != nil {
		context.Loggers.SendLoggerError("API创建刷新索引失败: ", err)
	}
	if statusCode == 200 {
		return true
	}
	return false
}
