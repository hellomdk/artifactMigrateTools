package jfrog

import (
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/util"
	"fmt"
	"log"
	"strings"
)

/**
探活
client.Ping()
*/
func (jf *Jfrog) Ping(context *config.Context) (string, error) {
	url := "/api/system/ping"
	statusCode, body, _ := util.Client.GetString(jf.HttpClient.BaseURL+url, jf.HttpClient.Header, nil)
	if statusCode != 200 {
		context.Loggers.SendLoggerError(fmt.Sprint("测试连接失败: ", statusCode), nil)
	}
	return body, nil
}

/**
获取仓库列表
client.GetRepositories()
*/
func (jf *Jfrog) GetRepositories(context *config.Context) []JfrogRepository {
	url := "/api/repositories"
	var cacheConfig = &[]JfrogRepository{}
	statusCode, _ := util.Client.Get(jf.HttpClient.BaseURL+url, jf.HttpClient.Header, nil, cacheConfig)
	if statusCode != 200 {
		context.Loggers.SendLoggerError(fmt.Sprint("API获取仓库列表失败: ", statusCode), nil)
	}
	return *cacheConfig
}

/**
获取仓库详情
client.GetRepository("Gradle_Test")
*/
func (jf *Jfrog) GetRepository(context *config.Context, repoKey string) JfrogRepositoryDetail {
	url := "/api/repositories/{repoKey}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)

	var cacheConfig = &JfrogRepositoryDetail{}
	statusCode, _ := util.Client.Get(jf.HttpClient.BaseURL+url, jf.HttpClient.Header, nil, cacheConfig)
	if statusCode != 200 {
		context.Loggers.SendLoggerError(fmt.Sprint("API获取仓库详情: ", statusCode), nil)
	}
	return *cacheConfig
}

/**
获取制品列表
client.GetArtifact("Gradle_Test", "maven-metadata-local.xml")
*/
func (jf *Jfrog) GetArtifacts(context *config.Context, repoKey string) JfrogArtifacts {
	url := "/api/storage/{repoKey}?list&deep=1&listFolders=0&mdTimestamps=1"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)

	var cacheConfig = &JfrogArtifacts{}
	statusCode, _ := util.Client.Get(jf.HttpClient.BaseURL+url, jf.HttpClient.Header, nil, cacheConfig)
	if statusCode != 200 {
		context.Loggers.SendLoggerError(fmt.Sprintf("API获取制品列表【%s】: %d", url, statusCode), nil)
	}
	return *cacheConfig
}

/**
获取制品详情
client.GetArtifact("Gradle_Test", "maven-metadata-local.xml")
*/
func (jf *Jfrog) GetArtifact(context *config.Context, repoKey, repoPath string) JfrogArtifact {
	url := "/api/storage/{repoKey}/{repoPath}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	url = strings.Replace(url, "{repoPath}", repoPath, 1)

	var cacheConfig = &JfrogArtifact{}
	statusCode, _ := util.Client.Get(jf.HttpClient.BaseURL+url, jf.HttpClient.Header, nil, cacheConfig)
	if statusCode != 200 {
		context.Loggers.SendLoggerError(fmt.Sprint("API获取制品详情: ", statusCode), nil)
	}
	return *cacheConfig
}

/**
获取制品属性
client.GetArtifactProperties("Gradle_Test", "maven-metadata-local.xml")
*/
func (jf *Jfrog) GetArtifactProperties(repoKey, repoPath string) (JfrogProp, error) {
	url := "/api/storage/{repoKey}/{repoPath}?properties"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	url = strings.Replace(url, "{repoPath}", repoPath, 1)

	var cacheConfig = &JfrogProp{}
	_, err := util.Client.Get(jf.HttpClient.BaseURL+url, jf.HttpClient.Header, nil, cacheConfig)
	if err != nil {
		log.Println("API获取制品属性: ", err)
	}
	return *cacheConfig, nil
}

/**
制品下载
client.DownloadFile("http://192.168.80.97:8082/artifactory/api/storage/Gradle_Test/maven-metadata-local.xml", "./maven-metadata-local.xml")
*/
//func (jf *Jfrog) DownloadFile(url string, filepath string) error {
//	// 创建一个HTTP客户端
//	client := &http.Client{}
//
//	// 创建一个HTTP请求
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		return fmt.Errorf("创建HTTP请求失败：%s", err)
//	}
//
//	// 发送HTTP请求并获取响应
//	resp, err := client.Do(req)
//	if err != nil {
//		return fmt.Errorf("发送HTTP请求失败：%s", err)
//	}
//	defer resp.Body.Close()
//
//	// 创建一个文件用于保存下载的文件
//	file, err := os.Create(filepath)
//	if err != nil {
//		return fmt.Errorf("创建文件失败：%s", err)
//	}
//	defer file.Close()
//
//	// 将响应体写入文件
//	_, err = io.Copy(file, resp.Body)
//	if err != nil {
//		return fmt.Errorf("写入文件失败：%s", err)
//	}
//
//	return nil
//}
