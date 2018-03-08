//write类

package logfox

import (
	"os"
	"sync"
)

type Writer struct {
	path                 string
	fileName             string
	fileSuffixTimeString string
	backUpDay            int
	file                 *os.File
	msgs                 chan string
	waitQueue            *sync.WaitGroup
}
