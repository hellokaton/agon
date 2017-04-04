package color

import (
	"fmt"
	"runtime"
	"strconv"
)

const (
	// common
	Reset = "\033[0m"
	Normal = 0
	Bold = 1
	Dim = 2
	Underline = 4
	Blink = 5
	Reverse = 7
	Hidden = 8

	// color
	Black = 30
	Red = 31
	Green = 32
	Yellow = 33
	Blue = 34
	Purple = 35
	Cyan = 36
	LightGray = 37
	DarkGray = 90
	LightRed = 91
	LightGreen = 92
	LightYellow = 93
	LightBlue = 94
	LightPurple = 95
	LightCyan = 96
	White = 97
	Success = Green
	Error = Red
	Warning = Yellow
)

// 打印颜色到控制台
func Print(c int, format string, v... interface{}) () {
	fmt.Print(Color(c, fmt.Sprintf(format, v...)))
}

func Println(c int, format string, v... interface{}) () {
	fmt.Println(Color(c, fmt.Sprintf(format, v...)))
}

// 返回带颜色的字符串
// c为颜色代码，f为字体格式
func Color(c int, str string) string {
	return Render(c, str)
}

// 判断是否是windows系统
func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}

// render string to output
func Render(color int, txt string) string {
	if IsWindows() {
		return txt
	}
	return "\x1b\033[" + strconv.Itoa(color) + "m" + txt + Reset
}