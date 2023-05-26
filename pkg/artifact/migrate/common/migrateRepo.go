package common

import (
	"fmt"
	"jfrogToArtifact/pkg/artifact/config"
	"jfrogToArtifact/pkg/artifact/httpclient"
	"jfrogToArtifact/pkg/artifact/repo"
	"net/http"
)

func CreateRepository(repositories config.Repositories) {
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

	repoData := repo.Repository{
		RepoKey:     repositories.RepoKey,
		RepoType:    repositories.RepoType,
		Description: repositories.Description,
		Layout:      repositories.Layout,
		Browse:      true,
		Url:         repositories.Url,
		Username:    repositories.Username,
		Password:    repositories.Password,
		ProtocolSpecific: repo.ProtocolType{
			ProtocolType:     repositories.ProtocolType,
			DockerApiVersion: "V2",
		},
		ProjectKey:            repositories.ProjectKey,
		SelectedRepositories:  repositories.SelectedRepositories,
		ResolvedRepositories:  repositories.SelectedRepositories,
		DefaultDeploymentRepo: repositories.DefaultDeploymentRepo,
	}

	got := jf.CreateRepository(repoData)
	fmt.Println(got)
}
