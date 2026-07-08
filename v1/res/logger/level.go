package logger

import "fmt"

func log(level string, args ...any) {

	if Default == nil {
		return
	}

	Default.write(level, fmt.Sprint(args...))
}

func Debug(args ...any) {
	log("DEBUG", args...)
}

func Info(args ...any) {
	log("INFO", args...)
}

func Warn(args ...any) {
	log("WARN", args...)
}

func Error(args ...any) {
	log("ERROR", args...)
}

func Fatal(args ...any) {
	log("FATAL", args...)
}

func Security(args ...any) {
	log("SECURITY", args...)
}

func Inventory(args ...any) {
	log("INVENTORY", args...)
}

func Sync(args ...any) {
	log("SYNC", args...)
}

func Remote(args ...any) {
	log("REMOTE", args...)
}

func Update(args ...any) {
	log("UPDATE", args...)
}

func Network(args ...any) {
	log("NETWORK", args...)
}

func Database(args ...any) {
	log("DATABASE", args...)
}

func Worker(args ...any) {
	log("WORKER", args...)
}
