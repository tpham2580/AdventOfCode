package logger

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(message string) {
	l.infoLogger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	l.infoLogger.Output(2, message)
}

func (l *Logger) InfoInts(nums []int) {
	message := fmt.Sprintf("%v", nums)
	l.infoLogger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	l.infoLogger.Output(2, message)
}

func (l *Logger) InfoBool(b bool) {
	message := fmt.Sprintf("%v", b)
	l.infoLogger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	l.infoLogger.Output(2, message)
}

func (l *Logger) InfoStrings(strings []string) {
	message := fmt.Sprintf("%v", strings)
	l.infoLogger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	l.infoLogger.Output(2, message)
}

func (l *Logger) Error(message string) {
	l.errorLogger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	l.errorLogger.Output(2, message)
	l.errorLogger.Println("Stack trace: \n" + string(debug.Stack()))
}
