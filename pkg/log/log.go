package log

import (
	"io"
	"log"
)

type Logger struct {
	Error *log.Logger
}

func NewLogger(out io.Writer) *Logger {
	return &Logger{
		Error: log.New(out, "[ERROR] ", log.Ldate|log.Ltime|log.Llongfile),
	}
}
