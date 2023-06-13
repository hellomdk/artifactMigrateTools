package test

import (
	"fmt"
	config2 "jfrogToArtifact/internal/service/config"
	"jfrogToArtifact/pkg/artifact/httpclient"
	"net/http"
	"testing"
)

func TestClient_DownloadFile(t *testing.T) {
	config, er := config2.NewConfig()
	if er != nil {
		fmt.Println(config)
	}
	jf := &Jfrog{
		httpclient.HttpClient{
			BaseURL:  config.SourceRepo.URL,
			Username: config.SourceRepo.Username,
			Password: config.SourceRepo.Password,
			Header:   http.Header{},
		},
	}
	err := jf.DownloadFile(jf.HttpClient.BaseURL+"/api/storage/fufaquan-go-local/go.zero/@v/v1.0.0.zip", "./mdk.zip")
	if err != nil {
		fmt.Println(err)
		return
	}

}

func TestClient_GetArtifact(t *testing.T) {
	config, er := config2.NewConfig()
	if er != nil {
		fmt.Println(config)
	}
	jf := &Jfrog{
		httpclient.HttpClient{
			BaseURL:  config.SourceRepo.URL,
			Username: config.SourceRepo.Username,
			Password: config.SourceRepo.Password,
			Header:   http.Header{},
		},
	}
	got, err := jf.GetArtifact("fufaquan-go-local", "go.zero/@v/v1.0.0.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(got)
}

//todo 结果集解析
func TestClient_GetArtifactProperties(t *testing.T) {
	config, er := config2.NewConfig()
	if er != nil {
		fmt.Println(config)
	}
	jf := &Jfrog{
		httpclient.HttpClient{
			BaseURL:  config.SourceRepo.URL,
			Username: config.SourceRepo.Username,
			Password: config.SourceRepo.Password,
			Header:   http.Header{},
		},
	}
	got, err := jf.GetArtifactProperties("fufaquan-go-local", "go.zero/@v/v1.0.0.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(got)
}

func TestClient_GetRepositories(t *testing.T) {
	config, er := config2.NewConfig()
	if er != nil {
		fmt.Println(config)
	}
	jf := &Jfrog{
		httpclient.HttpClient{
			BaseURL:  config.SourceRepo.URL,
			Username: config.SourceRepo.Username,
			Password: config.SourceRepo.Password,
			Header:   http.Header{},
		},
	}
	got, err := jf.GetRepositories(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(got)
}

func TestClient_GetRepository(t *testing.T) {
	config, er := config2.NewConfig()
	if er != nil {
		fmt.Println(config)
	}
	jf := &Jfrog{
		httpclient.HttpClient{
			BaseURL:  config.SourceRepo.URL,
			Username: config.SourceRepo.Username,
			Password: config.SourceRepo.Password,
			Header:   http.Header{},
		},
	}
	got, err := jf.GetRepository("fufaquan-go-local")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(got)
}

func TestClient_Ping(t *testing.T) {
	config, er := config2.NewConfig()
	if er != nil {
		fmt.Println(config)
	}
	jf := &Jfrog{
		httpclient.HttpClient{
			BaseURL:  config.SourceRepo.URL,
			Username: config.SourceRepo.Username,
			Password: config.SourceRepo.Password,
			Header:   http.Header{},
		},
	}
	got, err := jf.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(got)
}
