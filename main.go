package main

import (
	"golog/golog"
	"time"
)

type user struct {
	name     string
	age      int
	birthday string
}

func main() {
	wanglei := user{
		name:     "wanglei",
		age:      0,
		birthday: time.Now().Format("2006-01-2 15:4"),
	}
	golog.Info(wanglei)
	golog.DEBUG(wanglei)
}
