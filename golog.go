package golog

import (
	"log"
	"runtime"
)

// Info 普通信息
func Info(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix("INFO  " + "[" + funcName + "] ")
	log.Print(white(msg...))
}

// Debug 调试信息
func Debug(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(Green("DEBUG " + "[" + funcName + "] "))
	log.Print(Green(msg...))
}

// Warn 警告信息
func Warn(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(Yellow("WARN  " + "[" + funcName + "] "))
	log.Print(Yellow(msg...))
}

// Error 错误信息
func Error(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(Red("ERROR " + "[" + funcName + "] "))
	log.Print(Red(msg...))
}

// ErrorAndExit 错误信息并退出程序
func ErrorAndExit(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(Purple("EXIT  " + "[" + funcName + "] "))
	log.Fatal(Purple(msg...))
}
