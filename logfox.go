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
	return nil
}

func Debug(v ...interface{}) {
	fmt.Println(v...)
}
