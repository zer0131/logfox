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

//根据level值获取字符串
func (level Level) StringLevel() string {
	if str, ok := levelMapper[level]; ok {
		return str
	}
	panic(errors.New("Wrong Level Number"))
}

//校验level是否有效
func (level Level) IsValidLevel() bool {
	for _, val := range allLevels {
		if val == level {
			return true
		}
	}
	return false
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

//声明级别数组
var allLevels = [...]Level{
	DebugLevel,
	InfoLevel,
	NoticeLevel,
	WarnLevel,
	ErrorLevel,
	PanicLevel,
}
