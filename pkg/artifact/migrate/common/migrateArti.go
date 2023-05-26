package common

import (
	"fmt"
	"jfrogToArtifact/pkg/artifact/config"
	"jfrogToArtifact/pkg/artifact/httpclient"
	"jfrogToArtifact/pkg/artifact/repo"
	"net/http"
)

func CreateArtifact(arti config.Arti) {
	config, er := config.NewConfig()
	if er != nil {
		fmt.Println(config)
	}

	jf := &repo.Repo{

		httpclient.HttpClient{
			BaseURL:  config.TargetRepo.URL,
			Username: config.TargetRepo.Username,
			Password: config.TargetRepo.Password,
			Header:   http.Header{},
		},
	}
	httpclient.Auth(jf.HttpClient)
	jf.HttpClient.Header.Add("Content-Type", "application/octet-stream")
	jf.HttpClient.Header.Add("X-Checksum-Deploy", "true")
	jf.HttpClient.Header.Add("X-Checksum-Sha1", "f572d396fae9206628714fb2ce00f72e94f2258f")
	jf.HttpClient.Header.Add("X-Checksum-Sha256", "5891b5b522d5df086d0ff0b110fbd9d21bb4fc7163af34d08286a2e846f6be03")

	got := jf.CreateArtifact(arti.Repo, arti.Path)
	fmt.Println(got)

}
