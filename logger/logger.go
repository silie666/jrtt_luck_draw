package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"jrtt/config"
	"jrtt/model"
	"os"
	"time"
)

func LoggerToFile(errs string) {
	config := config.GetJrttUrl()
	fileName := config["LOGS_PATH"].(string)
	src,err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err",err)
	}

	logger := logrus.New()
	logger.Out = src

	logger.SetLevel(logrus.DebugLevel)

	logWriter,err := rotatelogs.New(
		// 分割后的文件名称
		fileName + ".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})
	logger.AddHook(lfHook)
	startTime := time.Now()
	logger.WithFields(logrus.Fields{
		"start_time"  : startTime,
		"err"      : errs,
	}).Error()
}

func LogToMysql(errs string,exit bool) {
	var log model.Log
	log.Add(errs)
	fmt.Println(errs)
	if exit {
		os.Exit(0)
	}
}

