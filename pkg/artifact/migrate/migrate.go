package migrate

import (
	"fmt"
	"jfrogToArtifact/pkg/artifact/config"
	"jfrogToArtifact/pkg/artifact/migrate/common"
	"jfrogToArtifact/pkg/artifact/migrate/protocol"
)

func MigrateArti(artiList *[]config.Arti) error {
	// 读取 value
	for _, arti := range *artiList {
		fmt.Println("正在同步制品: ", arti.Name)
		if arti.ProtocolType == "docker" {
			protocol.DeployDocker(arti)
		} else {
			common.CreateArtifact(arti)
		}
	}
	return nil
}

func MigrateRepo(repoList *[]config.Repositories) error {
	// 读取 value
	for _, repo := range *repoList {
		fmt.Println("正在迁移仓库: ", repo.RepoKey)
		common.CreateRepository(repo)
	}

	return nil
}

func Update_Arti_Checksum(artiList *[]config.Arti) error {
	// 读取 value
	for _, arti := range *artiList {
		fmt.Println("正在更新制品: ", arti.Name)
		common.UpdateArtifactSum(arti)
	}

	return nil
}