//日志输出
//带有f后缀的为可格式化输出

package logfox

import (
	"fmt"
)

const (
	DEFAULT_FILEWRITER_MAX_EXPIRE_DAY          = 7                       //日志默认有效天数
	DEFAULT_FILEWRITER_FILE_SUFFIX_TIME_STRING = "2006010215"            //日志文件时间前缀
	DEFAULT_FILEWRITER_MSG_SUFFIX_TIME_STRING  = "06-01-02 15:04:05.999" //日志时间格式
)

var loggerObj *Logger = nil
var logLevel string

//日志初始化
func Init(path string, app string, level string, expireDay int) error {
	if loggerObj != nil {
		return nil
	}
	logLevel = level
	if expireDay <= 0 {
		expireDay = DEFAULT_FILEWRITER_MAX_EXPIRE_DAY
	}
	var err error
	if loggerObj, err = NewLogger(path, app, expireDay, DEFAULT_FILEWRITER_FILE_SUFFIX_TIME_STRING); err != nil {
		return err
	}
	return nil
}

func Debug(v ...interface{}) {
	if DebugLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(fmt.Sprint(v...), DebugLevel)
	}
}

func Debugf(format string, v ...interface{}) {
	if DebugLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(fmt.Sprintf(format, v...), DebugLevel)
	}
}

func Info(v ...interface{}) {
	if InfoLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(fmt.Sprint(v...), InfoLevel)
	}
}

func Infof(format string, v ...interface{}) {
	if InfoLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(fmt.Sprintf(format, v...), InfoLevel)
	}
}

func Notice(v ...interface{}) {
	if NoticeLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(fmt.Sprint(v...), NoticeLevel)
	}
}

func Noticef(format string, v ...interface{}) {
	if NoticeLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(fmt.Sprintf(format, v...), NoticeLevel)
	}
}

func Warn(v ...interface{}) {
	if WarnLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.OutputWf(fmt.Sprint(v...), WarnLevel)
	}
}

func Warnf(format string, v ...interface{}) {
	if WarnLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.OutputWf(fmt.Sprintf(format, v...), WarnLevel)
	}
}

func Error(v ...interface{}) {
	if ErrorLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.OutputWf(fmt.Sprint(v...), ErrorLevel)
	}
}

func Errorf(format string, v ...interface{}) {
	if ErrorLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.OutputWf(fmt.Sprintf(format, v...), ErrorLevel)
	}
}

func Panic(v ...interface{}) {
	if PanicLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.OutputWf(fmt.Sprint(v...), PanicLevel)
		panic(fmt.Sprint(v...))
	}
}

func Panicf(format string, v ...interface{}) {
	if PanicLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.OutputWf(fmt.Sprintf(format, v...), PanicLevel)
		panic(fmt.Sprintf(format, v...))
	}
}

func Close() {
	loggerObj.Close()
}
