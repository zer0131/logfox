//日志输出
//带有f后缀的为可格式化输出

package logfox

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"net/http"
	"time"
)

const (
	DEFAULT_FILEWRITER_MAX_EXPIRE_DAY          = 7                       //日志默认有效天数
	DEFAULT_FILEWRITER_FILE_SUFFIX_TIME_STRING = "2006010215"            //日志文件时间前缀
	DEFAULT_FILEWRITER_MSG_SUFFIX_TIME_STRING  = "06-01-02 15:04:05.999" //日志时间格式
	LogIDKey                                   = "log-id"
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

func GenLogId() string {
	var t int64 = time.Now().UnixNano() / 1000000
	var r int = rand.Intn(10000)
	return fmt.Sprintf("%d%d", t, r)
}

func NewContextWithLogID(ctx context.Context) context.Context {
	md := metadata.New(make(map[string]string))
	md.Set(LogIDKey, GenLogId())
	return metadata.NewIncomingContext(ctx, md)
}

func NewGrpcContextWithLogID(ctx context.Context) context.Context {
	logid := GenLogId()

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		md.Set(LogIDKey, logid)
	}

	return metadata.AppendToOutgoingContext(ctx, LogIDKey, logid)
}

func NewContextWithSpecifyLogID(ctx context.Context, logId string) context.Context {
	md := metadata.New(make(map[string]string))
	md.Set(LogIDKey, logId)
	return metadata.NewIncomingContext(ctx, md)
}

func NewContextWithHttpReq(ctx context.Context, r *http.Request) context.Context {
	logId := r.Header.Get(LogIDKey)
	return NewContextWithSpecifyLogID(ctx, logId)
}

func LogIdFromContext(ctx context.Context) (string, bool) {
	md, _ := metadata.FromIncomingContext(ctx)
	arr := md.Get(LogIDKey)
	if len(arr) == 1 {
		return arr[0], true
	}
	return "", false
}

func DebugWithContext(ctx context.Context, v ...interface{}) {
	logId, ok := LogIdFromContext(ctx)
	if ok {
		Debug(fmt.Sprintf("[%s]", logId) + " " + fmt.Sprint(v...))
	} else {
		Debug(v...)
	}
}

func DebugfWithContext(ctx context.Context, format string, v ...interface{}) {
	logId, ok := LogIdFromContext(ctx)
	if ok {
		Debugf(fmt.Sprintf("[%s] %s", logId, format), v...)
	} else {
		Debugf(format, v...)
	}
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

func InfoWithContext(ctx context.Context, v ...interface{}) {
	logId, ok := LogIdFromContext(ctx)
	if ok {
		Info(fmt.Sprintf("[%s]", logId) + " " + fmt.Sprint(v...))
	} else {
		Info(v...)
	}
}

func InfofWithContext(ctx context.Context, format string, v ...interface{}) {
	logId, ok := LogIdFromContext(ctx)
	if ok {
		Infof(fmt.Sprintf("[%s] %s", logId, format), v...)
	} else {
		Infof(format, v...)
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

func NoticeWithContext(ctx context.Context, v ...interface{}) {
	logId, ok := LogIdFromContext(ctx)
	if ok {
		Notice(fmt.Sprintf("[%s]", logId) + " " + fmt.Sprint(v...))
	} else {
		Notice(v...)
	}
}

func NoticefWithContext(ctx context.Context, format string, v ...interface{}) {
	logId, ok := LogIdFromContext(ctx)
	if ok {
		Noticef(fmt.Sprintf("[%s] %s", logId, format), v...)
	} else {
		Noticef(format, v...)
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

func WarnWithContext(ctx context.Context, v ...interface{}) {
	logId, ok := LogIdFromContext(ctx)
	if ok {
		Warn(fmt.Sprintf("[%s]", logId) + " " + fmt.Sprint(v...))
	} else {
		Warn(v...)
	}
}

func WarnfWithContext(ctx context.Context, format string, v ...interface{}) {
	logId, ok := LogIdFromContext(ctx)
	if ok {
		Warnf(fmt.Sprintf("[%s] %s", logId, format), v...)
	} else {
		Warnf(format, v...)
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

func ErrorWithContext(ctx context.Context, v ...interface{}) {
	logId, ok := LogIdFromContext(ctx)
	if ok {
		Error(fmt.Sprintf("[%s]", logId) + " " + fmt.Sprint(v...))
	} else {
		Error(v...)
	}
}

func ErrorfWithContext(ctx context.Context, format string, v ...interface{}) {
	logId, ok := LogIdFromContext(ctx)
	if ok {
		Errorf(fmt.Sprintf("[%s] %s", logId, format), v...)
	} else {
		Errorf(format, v...)
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

func PanicWithContext(ctx context.Context, v ...interface{}) {
	logId, ok := LogIdFromContext(ctx)
	if ok {
		Panic(fmt.Sprintf("[%s]", logId) + " " + fmt.Sprint(v...))
	} else {
		Panic(v...)
	}
}

func PanicfWithContext(ctx context.Context, format string, v ...interface{}) {
	logId, ok := LogIdFromContext(ctx)
	if ok {
		Panicf(fmt.Sprintf("[%s] %s", logId, format), v...)
	} else {
		Panicf(format, v...)
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
