package logger

import (
	"artifactMigrateTools/internal/util"
	"fmt"
	"log"
	"os"
)

const INFO = 1
const ERROR = 2
const DEBUG = 3

type Loggers struct {
	Logs log.Logger
}

type Line struct {
	Level    string // 级别
	Content  string //日志内容
	UnixNano string //日志内容
}

func NewLogger(filePath string) *Loggers {
	// 打开文件
	fileName := fmt.Sprintf(filePath)
	if !util.PathExists(fileName) {
		util.CreatePathDir(fileName)
	}
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	// 通过New方法自定义Logger，New的参数对应的是Logger结构体的output, prefix和flag字段
	logger := log.New(f, "[INFO] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)

	loggerObj := &Loggers{
		Logs: *logger,
	}
	return loggerObj
}

func (l *Loggers) SendLogger(level int, content ...string) {
	log.Println(content)
	log.SetFlags(level)
	l.Logs.Println(content)
}

func (l *Loggers) SendLoggerInfo(content ...string) {
	l.SendLogger(INFO, content...)
}
func (l *Loggers) SendLoggerDebug(content ...string) {
	l.SendLogger(DEBUG, content...)
}

func (l *Loggers) SendLoggerError(content string, err error) {
	l.SendLogger(ERROR, fmt.Sprint(content, err))
}
