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

func (this *Writer) Close() {
	//等chan里的写完
	this.waitQueue.Wait()
	close(this.msgs)
}

func (this *Writer) newFile() (*os.File, error) {
	filePath := filepath.Join(this.path, this.fileName)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		file.Close()
		return nil, err
	}
	return file, nil
}

//将当前log文件按时间存起来
func (this *Writer) dumpFile() error {
	this.file.Close()

	//减一小时
	now := time.Now()
	h, _ := time.ParseDuration("-1h")
	preHour := now.Add(h)

	timeSuffix := preHour.Format(this.fileSuffixTimeString)
	preFilePath := filepath.Join(this.path, this.fileName+"."+timeSuffix)
	curFilePath := filepath.Join(this.path, this.fileName)
	if err := os.Rename(curFilePath, preFilePath); err != nil {
		return err
	}

	newFile, err := this.newFile()
	if err != nil {
		return err
	}
	this.file = newFile
	return nil
}

func (this *Writer) write(msg string) {
	this.waitQueue.Add(1)
	this.msgs <- msg
}

func (this *Writer) rotate() {
	tDuration := this.getTDuration()

	//定时器
	tick := time.NewTimer(tDuration)

	for {
		select {
		case msg, ok := <-this.msgs:
			if ok || msg != "" {
				this.waitQueue.Done()
				if err := this.fileAppend(msg); err != nil {
					panic(err)
				}
			} else {
				//chan已关闭并且里面消息已消费完
				this.file.Close()
				return
			}

		case <-tick.C:
			//日志切分
			err := this.dumpFile()
			if err != nil {
				panic(err)
			}

			//重新计时
			tDuration = this.getTDuration()
			tick.Reset(tDuration)

			//处理过期数据
			go this.clean()
		}
	}
}

func (this *Writer) fileAppend(msg string) error {
	_, err := this.file.WriteString(msg)
	return err
}

func (this *Writer) clean() {
	filepath.Walk(this.path, func(filePath string, info os.FileInfo, err error) error {
		if filePath == "." || filePath == ".." || filePath == this.path {
			return nil
		}
		idx := strings.LastIndex(filePath, ".")
		rs := []byte(filePath)
		suffix := rs[idx+1: len(rs)]
		fileSuffix := string(suffix)
		match := strings.LastIndex(filePath, this.fileName+"."+fileSuffix)
		if match < 0 {
			return nil
		}
		now := time.Now()
		day, _ := time.ParseDuration(fmt.Sprintf("-%dh", 24*this.expireDay))
		timeLast := now.Add(day).Format(this.fileSuffixTimeString)
		if timeLast > fileSuffix {
			os.Remove(filePath)
		}
		return nil
	})
}

func (this *Writer) getTDuration() time.Duration {
	//按照小时为单位四舍五入为整点
	roundHour := time.Now().Round(time.Hour)
	if roundHour.Before(time.Now()) {
		roundHour = roundHour.Add(time.Hour)
	}
	return roundHour.Sub(time.Now())
}
