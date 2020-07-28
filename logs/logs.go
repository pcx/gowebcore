package logs

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

// Create separate loggers for writing to stdout and stderr
var outLogger, errLogger *log.Logger
var trim bool

func ErrLogger() *log.Logger {
	if errLogger == nil {
		panic("Must call Init() before using ErrLogger()")
	}

	return errLogger
}

func InitLoggers(logPrefix string, _trim bool) {
	trim = _trim
	if trim {
		outLogger = log.New(os.Stdout, "", 0)
		errLogger = log.New(os.Stderr, "", 0)
	} else {
		outLogger = log.New(os.Stdout, fmt.Sprintf("%s ", logPrefix),
			log.Lmicroseconds|log.LUTC)
		errLogger = log.New(os.Stderr, logPrefix,
			log.Lmicroseconds|log.LUTC)
	}
}

func decorateLog(format string, level string) string {
	if trim {
		format = fmt.Sprintf("%s %s\n", level, format)
	} else {
		_, file, line, _ := runtime.Caller(2)
		format = fmt.Sprintf("%s:%d %s %s\n", file, line, level, format)
	}
	return format
}

func Debugf(format string, a ...interface{}) {
	outLogger.Printf(decorateLog(format, "DEBUG"), a...)
}

func Infof(format string, a ...interface{}) {
	outLogger.Printf(decorateLog(format, "INFO"), a...)
}

func Warnf(format string, a ...interface{}) {
	outLogger.Printf(decorateLog(format, "WARN"), a...)
}

func Printf(format string, a ...interface{}) {
	outLogger.Printf(format, a...)
}

func Panicf(format string, a ...interface{}) {
	panic(fmt.Sprintf(decorateLog(format, "PANIC"), a...))
}

func Debug(format string) {
	outLogger.Println(decorateLog(format, "DEBUG"))
}

func Info(format string) {
	outLogger.Println(decorateLog(format, "INFO"))
}

func Warn(format string) {
	outLogger.Println(decorateLog(format, "WARN"))
}

func Print(format string) {
	outLogger.Println(format)
}

func Panic(format string) {
	panic(decorateLog(format, "PANIC"))
}
