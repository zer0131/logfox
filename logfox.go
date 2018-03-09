package logfox

import (
	"fmt"
	"time"
)

const (
	DEFAULT_FILEWRITER_MAX_BACKUP_DAY          = 7
	DEFAULT_FILEWRITER_SPLIT_DURATION          = time.Hour
	DEFAULT_FILEWRITER_FILE_SUFFIX_TIME_STRING = "2006010215"
	DEFAULT_FILEWRITER_MSG_SUFFIX_TIME_STRING  = "06-01-02 15:04:05.999"
)

var loggerObj *Logger = nil
var logLevel string

func Init(path string, app string, level string) error {
	if loggerObj != nil {
		return nil
	}
	logLevel = level
	var err error
	if loggerObj, err = NewLogger(path, app, DEFAULT_FILEWRITER_MAX_BACKUP_DAY,
		DEFAULT_FILEWRITER_SPLIT_DURATION, DEFAULT_FILEWRITER_FILE_SUFFIX_TIME_STRING); err != nil {
		return err
	}
	return nil
}

func Debug(v ...interface{}) {
	if DebugLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(fmt.Sprint(v...), DebugLevel)
	}
}

func Info(v ...interface{}) {
	if InfoLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(fmt.Sprint(v...), InfoLevel)
	}
}

func Notice(v ...interface{}) {
	if NoticeLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(fmt.Sprint(v...), NoticeLevel)
	}
}

func Warn(v ...interface{}) {
	if WarnLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.OutputF(fmt.Sprint(v...), WarnLevel)
	}
}

func Error(v ...interface{}) {
	if ErrorLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.OutputF(fmt.Sprint(v...), ErrorLevel)
	}
}

func Panic(v ...interface{}) {
	if PanicLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.OutputF(fmt.Sprint(v...), PanicLevel)
		panic(fmt.Sprint(v...))
	}
}

func Close() {
	loggerObj.Close()
}
