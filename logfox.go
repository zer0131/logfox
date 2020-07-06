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
	if DebugLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(parseStr(ctx, v...), DebugLevel)
	}
}

func DebugfWithContext(ctx context.Context, format string, v ...interface{}) {
	if DebugLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(parseStrf(ctx, format, v...), DebugLevel)
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
	if InfoLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(parseStr(ctx, v...), InfoLevel)
	}
}

func InfofWithContext(ctx context.Context, format string, v ...interface{}) {
	if InfoLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(parseStrf(ctx, format, v...), InfoLevel)
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
	if NoticeLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(parseStr(ctx, v...), NoticeLevel)
	}
}

func NoticefWithContext(ctx context.Context, format string, v ...interface{}) {
	if NoticeLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.Output(parseStrf(ctx, format, v...), NoticeLevel)
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
	if WarnLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.OutputWf(parseStr(ctx, v...), WarnLevel)
	}
}

func WarnfWithContext(ctx context.Context, format string, v ...interface{}) {
	if WarnLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.OutputWf(parseStrf(ctx, format, v...), WarnLevel)
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
	if ErrorLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.OutputWf(parseStr(ctx, v...), ErrorLevel)
	}
}

func ErrorfWithContext(ctx context.Context, format string, v ...interface{}) {
	if ErrorLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		loggerObj.OutputWf(parseStrf(ctx, format, v...), ErrorLevel)
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
	if PanicLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		f := parseStr(ctx, v...)
		loggerObj.OutputWf(f, PanicLevel)
		panic(f)
	}
}

func PanicfWithContext(ctx context.Context, format string, v ...interface{}) {
	if PanicLevel >= levelMapperRev[logLevel] && loggerObj != nil {
		f := parseStrf(ctx, format, v...)
		loggerObj.OutputWf(f, PanicLevel)
		panic(f)
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

func parseStr(ctx context.Context, v ...interface{}) string {
	logId, ok := LogIdFromContext(ctx)
	var format string
	if ok {
		format = fmt.Sprintf("[%s]", logId) + " " + fmt.Sprint(v...)
	} else {
		format = fmt.Sprint(v...)
	}
	return format
}

func parseStrf(ctx context.Context, format string, v ...interface{}) string {
	logId, ok := LogIdFromContext(ctx)
	if ok {
		format = fmt.Sprintf("[%s] %s", logId, format)
	}
	return fmt.Sprintf(format, v...)
}
