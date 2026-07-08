package logger

import (
	"io"
)

type Writer struct{}

func (Writer) Write(p []byte) (int, error) {

	Info(string(p))

	return len(p), nil
}

func StdWriter() io.Writer {
	return Writer{}
}
