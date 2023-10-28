package golog

import (
	"log"
	"runtime"
)

func Info(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix("INFO  " + "[" + funcName + "] ")
	log.Print(white(msg...))
}
func Debug(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(Green("DEBUG " + "[" + funcName + "] "))
	log.Print(Green(msg...))
}
func Warn(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(Yellow("WARN  " + "[" + funcName + "] "))
	log.Print(Yellow(msg...))
}
func Error(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(Red("ERROR " + "[" + funcName + "] "))
	log.Print(Red(msg...))
}
func ErrorAndExit(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(Purple("EXIT  " + "[" + funcName + "] "))
	log.Fatal(Purple(msg...))
}
