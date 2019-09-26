//write类

package logfox

import (
	"os"
	"sync"
	"time"
	"path/filepath"
	"strings"
	"fmt"
)

type Writer struct {
	path                 string
	fileName             string
	fileSuffixTimeString string
	expireDay            int
	file                 *os.File
	msgs                 chan string
	waitQueue            *sync.WaitGroup
}

func NewWriter(path string, fileName string, fileSuffixTimeString string, expireDay int) (*Writer, error) {
	writer := &Writer{
		path:                 path,
		fileName:             fileName,
		fileSuffixTimeString: fileSuffixTimeString,
		expireDay:            expireDay,
		msgs:                 make(chan string, 10000),
		waitQueue:            new(sync.WaitGroup),
	}

	file, err := writer.newFile()
	if err != nil {
		return nil, err
	}

	writer.file = file

	go writer.rotate()

	return writer, nil
}

func (w *Writer) Close() {
	//等chan里的写完
	w.waitQueue.Wait()
	close(w.msgs)
}

func (w *Writer) newFile() (*os.File, error) {
	filePath := filepath.Join(w.path, w.fileName)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		_ = file.Close()
		return nil, err
	}
	return file, nil
}

//将当前log文件按时间存起来
func (w *Writer) dumpFile() error {
	_ = w.file.Close()

	//减一小时
	now := time.Now()
	h, _ := time.ParseDuration("-1h")
	preHour := now.Add(h)

	timeSuffix := preHour.Format(w.fileSuffixTimeString)
	preFilePath := filepath.Join(w.path, w.fileName+"."+timeSuffix)
	curFilePath := filepath.Join(w.path, w.fileName)
	if err := os.Rename(curFilePath, preFilePath); err != nil {
		return err
	}

	newFile, err := w.newFile()
	if err != nil {
		return err
	}
	w.file = newFile
	return nil
}

func (w *Writer) write(msg string) {
	w.waitQueue.Add(1)
	w.msgs <- msg
}

func (w *Writer) rotate() {
	tDuration := w.getTDuration()

	//定时器
	tick := time.NewTimer(tDuration)

	for {
		select {
		case msg, ok := <-w.msgs:
			if ok || msg != "" {
				w.waitQueue.Done()
				if err := w.fileAppend(msg); err != nil {
					panic(err)
				}
			} else {
				//chan已关闭并且里面消息已消费完
				_ = w.file.Close()
				return
			}

		case <-tick.C:
			//日志切分
			err := w.dumpFile()
			if err != nil {
				panic(err)
			}

			//重新计时
			tDuration = w.getTDuration()
			tick.Reset(tDuration)

			//处理过期数据
			go w.clean()
		}
	}
}

func (w *Writer) fileAppend(msg string) error {
	_, err := w.file.WriteString(msg)
	return err
}

func (w *Writer) clean() {
	_ = filepath.Walk(w.path, func(filePath string, info os.FileInfo, err error) error {
		if filePath == "." || filePath == ".." || filePath == w.path {
			return nil
		}
		idx := strings.LastIndex(filePath, ".")
		rs := []byte(filePath)
		suffix := rs[idx+1:len(rs)]
		fileSuffix := string(suffix)
		match := strings.LastIndex(filePath, w.fileName+"."+fileSuffix)
		if match < 0 {
			return nil
		}
		now := time.Now()
		day, _ := time.ParseDuration(fmt.Sprintf("-%dh", 24*w.expireDay))
		timeLast := now.Add(day).Format(w.fileSuffixTimeString)
		if timeLast > fileSuffix {
			_ = os.Remove(filePath)
		}
		return nil
	})
}

func (w *Writer) getTDuration() time.Duration {
	//按照小时为单位四舍五入为整点
	roundHour := time.Now().Round(time.Hour)
	if roundHour.Before(time.Now()) {
		roundHour = roundHour.Add(time.Hour)
	}
	return roundHour.Sub(time.Now())
}
