package main

import (
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/migrate"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:    "Jfrog migrate to Artifact CLI",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "init", Aliases: []string{"i"}, Usage: "初始化配置, value[config]"},
			&cli.StringFlag{Name: "verifyConfig", Aliases: []string{"vc"}, Usage: "配置校验"},
			&cli.StringFlag{Name: "genericProject", Aliases: []string{"gp"}, Usage: "生成空间project.yaml, value[all]"},
			&cli.StringFlag{Name: "migrateProject", Aliases: []string{"mp"}, Usage: "迁移空间，value[all,projectKey]"},
			&cli.StringFlag{Name: "genericRepo", Aliases: []string{"gr"}, Usage: "生成仓库repo.yaml, value[all]"},
			&cli.StringFlag{Name: "migrateRepo", Aliases: []string{"mr"}, Usage: "同步仓库，value[all,projectKey]"},
			&cli.StringFlag{Name: "genericArtifacts", Aliases: []string{"ga"}, Usage: "生成制品artifact.yaml, [all,projectKey]"},
			&cli.StringFlag{Name: "migrateArtifacts", Aliases: []string{"ma"}, Usage: "同步制品，value[all,projectKey]"},
		},
		Action: func(context *cli.Context) error {
			var args = make(map[string]string)
			args["init"] = GetArgsValue(context, "init", "i")
			args["verifyConfig"] = GetArgsValue(context, "verifyConfig", "vc")
			args["genericProject"] = GetArgsValue(context, "genericProject", "gp")
			args["migrateProject"] = GetArgsValue(context, "migrateProject", "mp")
			args["genericRepo"] = GetArgsValue(context, "genericRepo", "gr")
			args["migrateRepo"] = GetArgsValue(context, "migrateRepo", "mr")
			args["genericArtifacts"] = GetArgsValue(context, "genericArtifacts", "ga")
			args["migrateArtifacts"] = GetArgsValue(context, "migrateArtifacts", "ma")

			contextConfig := config.NewContext(args)
			contextConfig.InitLogger(contextConfig)

			newCommand := migrate.NewMigrateExcute(contextConfig)
			newCommand.Excute()
			return nil
		},
	}
	contextConfig := config.NewContextNoArgs()
	contextConfig.InitLogger(contextConfig)
	err := app.Run(os.Args)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func GetArgsValue(context *cli.Context, name, aliases string) string {
	if context.String(name) != "" {
		return context.String(name)
	} else if context.String(aliases) != "" {
		return context.String(aliases)
	}
	return ""
}
