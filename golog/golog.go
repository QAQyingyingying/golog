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
	log.SetPrefix(white("INFO  " + "[" + funcName + "] "))
	log.Print(white(msg...))
}
func DEBUG(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(green("DEBUG " + "[" + funcName + "] "))
	log.Print(green(msg...))
}
func WARN(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(yellow("WARN  " + "[" + funcName + "] "))
	log.Print(yellow(msg...))
}
func ERROR(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(red("ERROR " + "[" + funcName + "] "))
	log.Print(red(msg...))
}
func ERRORANDSTOP(msg ...any) {
	pc, _, _, ok := runtime.Caller(1)
	funcName := "unKnow"
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.SetPrefix(red("ERROR " + "[" + funcName + "] "))
	log.Fatal(red(msg...))
}
