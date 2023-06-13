package config

import (
	"artifactMigrateTools/internal/logger"
	"fmt"
	"os"
	"path"
	"time"
)

const (
	Workspace        string = "SYSTEM_WORKSPACE"      // 工作空间
	LoggerPath       string = "SYSTEM_LOGGER_PATHs"   // 日志路径
	ConfigFileName   string = "MIRRATE_CONFIG_NAME"   // 迁移配置文件名称
	ProjectFileName  string = "MIRRATE_PROJECT_NAME"  // 迁移空间文件名称
	RepoFileName     string = "MIRRATE_REPO_NAME"     // 迁移仓库文件名称
	ArtifactFileName string = "MIRRATE_ARTIFACT_NAME" // 迁移制品文件名称
)

type Context struct {
	Pool    map[string]string
	Loggers logger.Loggers
}

func NewContextNoArgs() *Context {
	return &Context{}
}

func NewContext(pool map[string]string) *Context {
	return &Context{
		Pool: pool,
	}
}

func (c Context) InitLogger(context *Context) {
	loggers := logger.NewLogger(c.GetLoggerPath())
	context.Loggers = *loggers
}

func (c Context) GetWorkspace() string {
	return c.GetEvn(Workspace, "/home/migrate")
}

func (c Context) GetLoggerPath() string {
	return path.Join(c.GetWorkspace(), "log", c.GetEvn(LoggerPath, fmt.Sprintf("app_%s.log", time.Now().Format("20060102"))))
}

func (c Context) GetConfigPath() string {
	return path.Join(c.GetWorkspace(), "yaml", c.GetEvn(ConfigFileName, "config.yaml"))
}
func (c Context) GetProjectPath() string {
	return path.Join(c.GetWorkspace(), "yaml", c.GetEvn(ProjectFileName, "project.yaml"))
}
func (c Context) GetRepoPath() string {
	return path.Join(c.GetWorkspace(), "yaml", c.GetEvn(RepoFileName, "repo.yaml"))
}
func (c Context) GetArtifactPath() string {
	return path.Join(c.GetWorkspace(), "yaml", c.GetEvn(ArtifactFileName, "artifact.yaml"))
}

// 获取系统变量
func (c Context) GetEvn(key string, defaultValue string) (val string) {
	result, ok := c.Pool[key]
	if ok && result != "" {
		return result
	} else {
		env := os.Getenv(key)
		if len(env) == 0 {
			return defaultValue
		}
		return env
	}
}
