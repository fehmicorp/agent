package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Logger struct {
	mu      sync.Mutex
	file    *os.File
	path    string
	maxSize int64
	console bool
}

var Default *Logger

func Init(path, filename string) error {

	if filename == "" {
		filename = "agent.log"
	}

	if path == "" {
		exe, _ := os.Executable()
		path = filepath.Dir(exe)
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	full := filepath.Join(path, filename)

	file, err := os.OpenFile(full,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644)

	if err != nil {
		return err
	}

	Default = &Logger{
		file:    file,
		path:    full,
		maxSize: 10 * 1024 * 1024,
		console: true,
	}

	Info("===================================")
	Info("Fehmi Logger Initialized")
	Info("===================================")

	return nil
}

func (l *Logger) write(level, msg string) {

	l.mu.Lock()
	defer l.mu.Unlock()

	l.rotate()

	line := fmt.Sprintf(
		"%s [%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		level,
		msg,
	)

	if l.console {
		fmt.Print(line)
	}

	if l.file != nil {
		l.file.WriteString(line)
	}
}

func Close() {

	if Default == nil {
		return
	}

	Default.file.Close()
}
