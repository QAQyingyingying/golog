package main

import "fmt"

const (
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textPurple
	textCyan
	textWhite
)

func Black(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textBlack, str)
}

func red(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textRed, str)
}
func yellow(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textYellow, str)
}
func green(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textGreen, str)
}
func Cyan(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textCyan, str)
}
func Blue(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textBlue, str)
}
func Purple(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textPurple, str)
}
func white(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textWhite, str)
}

func textColor(color int, str string) string {
	return fmt.Sprintf("\x1b[;%dm%s\x1b[0m", color, str)
}
