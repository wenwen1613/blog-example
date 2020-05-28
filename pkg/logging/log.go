package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix    = ""
	DefautlCallDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

//Debug 调试日志
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

//Info 记录日志
func Info(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

//Warn 警告日志
func Warn(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

//Error 错误日志
func Error(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

//Fatal 严重错误日志
func Fatal(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefautlCallDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
