package repo

import (
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/util"
	"fmt"
	"net/http"
	"testing"
)

func TestClient_Ping(t *testing.T) {
	config, er := config.NewConfig()
	if er != nil {
		fmt.Println(config)
	}
	jf := &Repo{
		util.HttpClient{
			BaseURL:  config.TargetRepo.URL,
			Username: config.TargetRepo.Username,
			Password: config.TargetRepo.Password,
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

//func TestClient_ExistRepository(t *testing.T) {
//	config, er := config.NewConfig()
//	if er != nil {
//		fmt.Println(config)
//	}
//	jf := &Repo{
//		util.HttpClient{
//			BaseURL:  config.TargetRepo.URL,
//			Username: config.TargetRepo.Username,
//			Password: config.TargetRepo.Password,
//			Header:   http.Header{},
//		},
//	}
//	got := jf.ExistRepository("mdk_generic")
//	fmt.Println(got)
//}

//func TestClient_ExistArtifact(t *testing.T) {
//	config, er := config.NewConfig()
//	if er != nil {
//		fmt.Println(config)
//	}
//	jf := &Repo{
//		util.HttpClient{
//			BaseURL:  config.TargetRepo.URL,
//			Username: config.TargetRepo.Username,
//			Password: config.TargetRepo.Password,
//			Header:   http.Header{},
//		},
//	}
//	got := jf.ExistArtifact("mdk_generic", "mdk/hellomdk.txt")
//	fmt.Println(got)
//}

//func TestClient_CreateRepository(t *testing.T) {
//	config, er := config.NewConfig()
//	if er != nil {
//		fmt.Println(config)
//	}
//	jf := &Repo{
//		util.HttpClient{
//			BaseURL:  config.TargetRepo.URL,
//			Username: config.TargetRepo.Username,
//			Password: config.TargetRepo.Password,
//			Header:   http.Header{},
//		},
//	}
//
//	repo := Repository{
//		RepoKey:     "hellomdk-remote",
//		RepoType:    "remote",
//		Description: "test create artifact",
//		Layout:      "simple-default",
//		Url:         "https://dev.gitee.work/",
//		Browse:      true,
//		ProtocolSpecific: ProtocolType{
//			ProtocolType: "generic",
//		},
//	}
//	httpclient.Auth(jf.HttpClient)
//
//	got := jf.CreateRepository(repo)
//	fmt.Println(got)
//}
//
//func TestClient_CreateArtifact(t *testing.T) {
//	config, er := config2.NewConfig()
//	if er != nil {
//		fmt.Println(config)
//	}
//
//	jf := &Repo{
//
//		httpclient.HttpClient{
//			BaseURL:  config.TargetRepo.URL,
//			Username: config.TargetRepo.Username,
//			Password: config.TargetRepo.Password,
//			Header:   http.Header{},
//		},
//	}
//	httpclient.Auth(jf.HttpClient)
//	jf.HttpClient.Header.Add("Content-Type", "application/octet-stream")
//	jf.HttpClient.Header.Add("X-Checksum-Deploy", "true")
//	jf.HttpClient.Header.Add("X-Checksum-Sha1", "f572d396fae9206628714fb2ce00f72e94f2258f")
//	jf.HttpClient.Header.Add("X-Checksum-Sha256", "5891b5b522d5df086d0ff0b110fbd9d21bb4fc7163af34d08286a2e846f6be03")
//
//	got := jf.CreateArtifact("mdk_generic", "hello.txt")
//	fmt.Println(got)
//}
//
//func TestClient_UpdateArtifactCheckSum(t *testing.T) {
//	config, er := config2.NewConfig()
//	if er != nil {
//		fmt.Println(config)
//	}
//
//	jf := &Repo{
//
//		httpclient.HttpClient{
//			BaseURL:  config.TargetRepo.URL,
//			Username: config.TargetRepo.Username,
//			Password: config.TargetRepo.Password,
//			Header:   http.Header{},
//		},
//	}
//	httpclient.Auth(jf.HttpClient)
//
//	got := jf.UpdateArtifactCheckSum("mdk_generic", "mdk/hello.txt", "8ec8535a88c27aa4ae1546450792c0f763e22ec5", "", "")
//	fmt.Println(got)
//}
