package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type BufferedLog struct {
	logChannel chan []byte
	logFile    *os.File
	lastOffset int64
}

func (b *BufferedLog) Write(p []byte) (n int, err error) {
	b.logChannel <- p

	return len(p), nil
}

func NewBufferLog(length int, filename string) (*BufferedLog, error) {
	bufferedLog := new(BufferedLog)
	bufferedLog.logChannel = make(chan []byte, length)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open error log logFile:", err)
		return nil, err
	}

	info, err := file.Stat()
	bufferedLog.lastOffset = info.Size()
	bufferedLog.logFile = file
	//bufferedLog.mu = sync.Mutex{}

	return bufferedLog, nil
}

func (b *BufferedLog) writeFile() {
	for {
		logLine, ok := <-b.logChannel
		if !ok {
			fmt.Println("Work is done")
			return
		}

		n, err := b.logFile.WriteAt(logLine, b.lastOffset)
		if err != nil {
			fmt.Printf("Write logLine file error:%v\n", err)
		}

		b.lastOffset = b.lastOffset + int64(n)

	}
}

func (b *BufferedLog) close() {
	close(b.logChannel)

	time.Sleep(time.Duration(100) * time.Millisecond)

	e := b.logFile.Close()
	if e != nil {
		fmt.Printf("Close file error:%v", e)
	}
}
