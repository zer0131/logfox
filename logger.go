//日志核心类

package logfox

import "time"

type logger struct {
	path          string
	app           string
	backUpDay     int
	splitDuration time.Duration
	iWriter       *writer
	fWriter       *writer
}
