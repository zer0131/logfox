//日志核心类

package logfox

import "time"

type Logger struct {
	path          string
	app           string
	backUpDay     int
	splitDuration time.Duration
	iWriter       *Writer
	fWriter       *Writer
}

func NewLogger(path string, app string, backUpDay int, splitDuration time.Duration, fileSuffixTimeString string) (*Logger, error) {
	return nil, nil
}
