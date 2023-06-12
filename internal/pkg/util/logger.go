package util

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var Loggers *logrus.Logger

func init() {
	// 创建并打开日志文件
	file, _ := setOutPutFile()
	// 不为空
	if Loggers != nil {
		// 日志写入文件
		Loggers.Out = file
		return
	}
	// 实例化
	logger := logrus.New()
	// 写入文件
	logger.Out = file
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Loggers = logger
}

func setOutPutFile() (*os.File, error) {
	// 记录当前时间
	now := time.Now()
	// log文件存放路径
	logFilePath := ""
	// 获取当前工作目录
	if workDir, err := os.Getwd(); err == nil {
		logFilePath = workDir + "/logs/"
	}
	// 文件夹是否存在
	_, err := os.Stat(logFilePath)
	// 文件夹不存在
	if os.IsNotExist(err) {
		// 创建文件夹
		if err = os.MkdirAll(logFilePath, os.ModePerm); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	// 日志文件名
	logFileName := now.Format("2006-01-02") + ".log"
	// 文件路径
	filePath := path.Join(logFilePath, logFileName)
	// 创建并写入文件
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	return file, nil
}
