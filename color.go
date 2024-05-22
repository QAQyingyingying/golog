package golog

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

// Black 黑色信息
func Black(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textBlack, str)
}

// Red 红色信息
func Red(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textRed, str)
}

// Yellow 黄色信息
func Yellow(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textYellow, str)
}

// Green 绿色信息
func Green(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textGreen, str)
}

// Cyan 青色信息
func Cyan(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textCyan, str)
}

// Blue 蓝色信息
func Blue(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textBlue, str)
}

// Purple 紫色信息
func Purple(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textPurple, str)
}

// White 白色信息
func White(msg ...any) string {
	str := fmt.Sprint(msg...)
	return textColor(textWhite, str)
}

func textColor(color int, str string) string {
	return fmt.Sprintf("\x1b[;%dm%s\x1b[0m", color, str)
}
