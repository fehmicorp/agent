package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type AgentLogger struct {
	file *os.File
}

var Logger *AgentLogger

func InitLogger(customPath string, fileName string) {
	var dir string
	execPath, err := os.Executable()
	if err != nil {
		dir = "."
	} else {
		dir = filepath.Dir(execPath)
	}
	if fileName == "" {
		fileName = "logger.txt"
	}
	var logPath string
	if customPath == "" {
		logPath = filepath.Join(dir, fileName)
	} else {
		logPath = filepath.Join(customPath, fileName)
	}
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to initialize file logger at %s: %v\n", logPath, err)
		return
	}
	Logger = &AgentLogger{file: file}
	Logger.Log("INFO", "=============================================")
	Logger.Log("INFO", "Fehmi Endpoint Agent Logger Initialized System")
	Logger.Log("INFO", "=============================================")
}

func (l *AgentLogger) Log(category string, message string) {
	if l == nil || l.file == nil {
		fmt.Printf("[%s] %s\n", category, message)
		return
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logLine := fmt.Sprintf("%s [%s] %s\n", timestamp, category, message)
	fmt.Print(logLine)
	l.file.WriteString(logLine)
}

func (l *AgentLogger) Close() {
	if l != nil && l.file != nil {
		l.file.Close()
	}
}
