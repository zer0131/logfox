//日志核心类

package logfox

import (
	"time"
	"os"
	"fmt"
	"runtime"
	"path/filepath"
)

type Logger struct {
	path          string
	app           string
	backUpDay     int
	splitDuration time.Duration
	iWriter       *Writer
	fWriter       *Writer
}

func NewLogger(path string, app string, backUpDay int, splitDuration time.Duration, fileSuffixTimeString string) (*Logger, error) {
	if err := os.MkdirAll(path, os.FileMode(0755)); err != nil {
		return nil, err
	}
	iWriter, erri := NewWriter(path, fmt.Sprintf("%s.log", app), fileSuffixTimeString, backUpDay, splitDuration)
	if erri != nil {
		return nil, erri
	}
	fWriter, errw := NewWriter(path, fmt.Sprintf("%s.log.wf", app), fileSuffixTimeString, backUpDay, splitDuration)
	if errw != nil {
		return nil, errw
	}
	logger := &Logger{
		path:          path,
		app:           app,
		backUpDay:     backUpDay,
		splitDuration: splitDuration,
		iWriter:       iWriter,
		fWriter:       fWriter,
	}
	return logger, nil
}

//正常输出
func (this *Logger) Output(msg string, level Level) {
	timeNow := time.Now().Format(DEFAULT_FILEWRITER_MSG_SUFFIX_TIME_STRING)
	_, file, line, _ := runtime.Caller(2)//获取文件和行号
	msgInput := fmt.Sprintf("%s %s %s:%d: %s\n", level.String(), timeNow, filepath.Base(file), line, msg)
	this.iWriter.write(msgInput)
}

//错误输出
func (this *Logger) OutputF(msg string, level Level) {
	timeNow := time.Now().Format(DEFAULT_FILEWRITER_MSG_SUFFIX_TIME_STRING)
	_, file, line, _ := runtime.Caller(2)
	msgInput := fmt.Sprintf("%s %s %s:%d: %s\n", level.String(), timeNow, filepath.Base(file), line, msg)
	this.fWriter.write(msgInput)
}

func (this *Logger) Close() {
	this.iWriter.Close()
	this.fWriter.Close()
}
