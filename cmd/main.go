package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"jfrogToArtifact/pkg/artifact/jfrog"
	_ "jfrogToArtifact/pkg/artifact/jfrog"
	"jfrogToArtifact/pkg/artifact/config"
	"os"
)

func test() {
	config, er := config.NewConfig()
	if er != nil {
		fmt.Println(config)
	}
	client := &jfrog.Client{
		Username: config.SourceRepo.Username,
		Password: config.SourceRepo.Password,
	}
	users, err := client.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(users))
}

func main() {
	app := &cli.App{
		Name:    "Jfrog migrate to Artifact CLI",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "workspace", Aliases: []string{"w"}, Usage: "工作空间"},
			&cli.StringFlag{Name: "setConfig", Aliases: []string{"sc"}, Usage: "系统配置"},
			&cli.StringFlag{Name: "verifyConfig", Aliases: []string{"vc"}, Usage: "配置校验"},
			&cli.StringFlag{Name: "getRepositories", Aliases: []string{"gr"}, Usage: "获取仓库列表"},
			&cli.StringFlag{Name: "migrateRepo", Aliases: []string{"mr"}, Usage: "同步仓库"},
			&cli.StringFlag{Name: "migrateArtifact", Aliases: []string{"ma"}, Usage: "同步制品"},
		},
		Action: func(context *cli.Context) error {
			var args = make(map[string]string)
			args["SYSTEM_WORKSPACE"] = context.String("workspace")
			setConfig := context.String("setConfig")
			verifyConfig := context.String("verifyConfig")
			getRepositories := context.String("getRepositories")
			migrateRepo := context.String("migrateRepo")
			migrateArtifact := context.String("migrateArtifact")

			config, er := config.NewConfig()
			if er != nil {
				fmt.Println(config)
			}
			//newOut, er := Gohttp.NewGohttp(&jobcenter.Context{
			//	Config: newConfig,
			//})
			//if newOut != nil {
			//
			//}
			//if er != nil {
			//	return er
			//}
			var err error

			if setConfig != "" {

			}

			if verifyConfig == "" {
				client := &jfrog.Client{
					Username: config.SourceRepo.Username,
					Password: config.SourceRepo.Password,
				}
				users, err := client.GetRepositories(config)
				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println(users)
			}

			if getRepositories != "" {

			}

			if migrateRepo != "" {

			}
			if migrateArtifact != "" {

			}

			if err != nil {
				return err
			}

			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
