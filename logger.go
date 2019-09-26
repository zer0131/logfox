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
	path      string
	app       string
	expireDay int
	iWriter   *Writer
	fWriter   *Writer
}

func NewLogger(path string, app string, expireDay int, fileSuffixTimeString string) (*Logger, error) {
	if err := os.MkdirAll(path, os.FileMode(0755)); err != nil {
		return nil, err
	}
	//正常输出
	iWriter, iErr := NewWriter(path, fmt.Sprintf("%s.log", app), fileSuffixTimeString, expireDay)
	if iErr != nil {
		return nil, iErr
	}
	//错误输出
	fWriter, wErr := NewWriter(path, fmt.Sprintf("%s.log.wf", app), fileSuffixTimeString, expireDay)
	if wErr != nil {
		return nil, wErr
	}
	logger := &Logger{
		path:      path,
		app:       app,
		expireDay: expireDay,
		iWriter:   iWriter,
		fWriter:   fWriter,
	}
	return logger, nil
}

//正常输出
func (lo *Logger) Output(msg string, level Level) {
	lo.iWriter.write(lo.msgInput(msg, level))
}

//错误输出
func (lo *Logger) OutputWf(msg string, level Level) {
	lo.fWriter.write(lo.msgInput(msg, level))
}

func (lo *Logger) msgInput(msg string, level Level) string {
	timeNow := time.Now().Format(DEFAULT_FILEWRITER_MSG_SUFFIX_TIME_STRING)
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s %s %s:%d: %s\n", level.String(), timeNow, filepath.Base(file), line, msg)
}

func (lo *Logger) Close() {
	lo.iWriter.Close()
	lo.fWriter.Close()
}
