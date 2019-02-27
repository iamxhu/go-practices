package logger

import (
	"fmt"
	"log"
	"os"
)

type logger interface {
	Info(...interface{})

	Warn(...interface{})

	Debug(...interface{})

	Error(...interface{})
}

var (
	Logger *MultiLogger
)

// MultiLogger - a group of loggers, which all of logs will be write two log files: data.log and error.log
type MultiLogger struct {
	errorLogger *AsyncLogger
	infoLogger  *AsyncLogger
	warnLogger  *AsyncLogger
	debugLogger *AsyncLogger
}

type AsyncLogger struct {
	*log.Logger
	*BufferedLog
}

func NewAsyncLogger(prefix string, bufferedLog *BufferedLog) (*AsyncLogger, error) {
	asyncLogger := new(AsyncLogger)
	asyncLogger.Logger = log.New(bufferedLog, prefix, log.Ldate|log.Ltime|log.Lshortfile)
	asyncLogger.BufferedLog = bufferedLog

	return asyncLogger, nil
}

func (a *MultiLogger) Info(i ...interface{}) {
	a.infoLogger.Println(i)
}

func (a *MultiLogger) Warn(i ...interface{}) {
	a.warnLogger.Println(i)
}

func (a *MultiLogger) Debug(i ...interface{}) {
	a.debugLogger.Println(i)
}

func (a *MultiLogger) Error(i ...interface{}) {
	a.errorLogger.Println(i)
}

//LoggerInit - Logger Initial, use to initial loggers, you must call this before use logger.
func LoggerInit(logDir string) {
	dataFilename := "data.log"
	if logDir != "" {
		dataFilename = logDir + "/" + dataFilename
	}

	dataLog, e := NewBufferLog(10000, dataFilename)
	if e != nil {
		fmt.Println("init logger error!")
		os.Exit(500)
	}

	errorFn := "error.log"
	if logDir != "" {
		errorFn = logDir + "/" + errorFn
	}

	errorLog, e := NewBufferLog(10000, errorFn)
	if e != nil {
		fmt.Println("init logger error!")
		os.Exit(500)
	}

	debugLogger, err := NewAsyncLogger("DEBUG:", dataLog)
	if err != nil {
		fmt.Println("init logger error!")
		os.Exit(500)
	}

	warnLogger, err := NewAsyncLogger("WARN:", dataLog)
	infoLogger, err := NewAsyncLogger("INFO:", dataLog)

	errorLogger, e := NewAsyncLogger("ERROR:", errorLog)

	Logger = &MultiLogger{
		debugLogger: debugLogger,
		warnLogger:  warnLogger,
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
	}

	go dataLog.writeFile()

	go errorLog.writeFile()
}

func Close() {
	Logger.errorLogger.close()
	Logger.debugLogger.close()
}
