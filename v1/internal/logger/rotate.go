package logger

import (
	"fmt"
	"os"
	"time"
)

func (l *Logger) rotate() {

	if l.file == nil {
		return
	}

	info, err := l.file.Stat()

	if err != nil {
		return
	}

	if info.Size() < l.maxSize {
		return
	}

	l.file.Close()

	newName := fmt.Sprintf(
		"%s.%s",
		l.path,
		time.Now().Format("20060102150405"),
	)

	os.Rename(l.path, newName)

	file, err := os.OpenFile(
		l.path,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)

	if err != nil {
		return
	}

	l.file = file
}
