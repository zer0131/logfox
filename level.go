//日志级别定义

package logfox

import "errors"

//别名
type Level uint8

//级别常量
const (
	DebugLevel  Level = iota
	InfoLevel
	NoticeLevel
	WarnLevel
	ErrorLevel
	PanicLevel
)

//实现String方法，输出时默认转换为字符串
func (this Level) String() string {
	if str, ok := levelMapper[this]; ok {
		return str
	}
	panic(errors.New("Wrong Level Number"))
}

var levelMapper = map[Level]string{
	DebugLevel:  "DEBUG",
	InfoLevel:   "INFO",
	NoticeLevel: "NOTICE",
	WarnLevel:   "WARN",
	ErrorLevel:  "ERROR",
	PanicLevel:  "PANIC",
}

var levelMapperRev = map[string]Level{
	"DEBUG":  DebugLevel,
	"INFO":   InfoLevel,
	"NOTICE": NoticeLevel,
	"WARN":   WarnLevel,
	"ERROR":  ErrorLevel,
	"PANIC":  PanicLevel,
}

