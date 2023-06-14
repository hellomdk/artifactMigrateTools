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

	// 设置日志级别
	//logger.SetOutput(os.Stdout) // 输出到控制台
	//logger.SetOutput(os.Stderr) // 输出到标准错误流

	loggerObj := &Loggers{
		Logs: *logger,
	}
	return loggerObj
}

func (l *Loggers) SendLoggerInfo(content ...string) {
	l.Logs.SetPrefix("[INFO]")
	l.Logs.Println(content)
	log.SetPrefix("[INFO]")
	log.Println(content)
}
func (l *Loggers) SendLoggerDebug(content ...string) {
	l.Logs.SetPrefix("[DEBUG]")
	l.Logs.Println(content)
}

func (l *Loggers) SendLoggerError(content string, err error) {
	l.Logs.SetPrefix("[ERROR]")
	l.Logs.Println(fmt.Sprint(content, err))
	log.SetPrefix("[ERROR]")
	log.Println(fmt.Sprint(content, err))
}
