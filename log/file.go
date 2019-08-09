/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by my-gin.
 * User: 605724193@qq.com
 * Date: 2019/08/07
 * Time: 11:18
 ********************************************************************************************
 */

package log

import (
	"time"
	"path"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"github.com/lestrrat/go-file-rotatelogs"
)

var (
	Logger *logrus.Logger
)

func init() {
	Logger = logrus.New()
}

func ConfigLogger(logger *logrus.Logger, logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)

	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d",
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		logger.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})

	logger.AddHook(lfHook)
}
