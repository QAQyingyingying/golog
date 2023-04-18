package golog

import (
	"log"
	"runtime"
)

func Info(msg string) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(White("INFO  " + "[" + funcName + "] "))
	log.Print(White(msg))
}
func DEBUG(msg string) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(Green("DEBUG " + "[" + funcName + "] "))
	log.Print(Green(msg))
}
func WARN(msg string) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(Yellow("WARN  " + "[" + funcName + "] "))
	log.Print(Yellow(msg))
}
func ERROR(msg string) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(Red("ERROR " + "[" + funcName + "] "))
	log.Print(Red(msg))
}
func ERRORANDSTOP(msg string) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(Red("ERROR " + "[" + funcName + "] "))
	log.Fatal(Red(msg))
}
