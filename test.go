package main

import (
	"github.com/biezhi/agon/color"
	"github.com/biezhi/agon/log"
)

func TestColor(){
	color.Println(color.Red, "|| Hello World")
	color.Println(color.Green, "|| Hello World")
	color.Println(color.Yellow, "|| Hello World")
	color.Println(color.Purple, "|| Hello %s", "jack")
}

func TestLog()  {
	log.Info("Hello Rose")
	log.Debug("Hello %s", "jack")
	log.Warn("Hello %s", "jack")
	log.Trace("Hello %s", "jack")
	log.Error("Hello %s", "jack")
}

func main() {
	TestColor()
	TestLog()
	//fmt.Println("\033[32;1m我被变成了蓝色，\033[0m我是原来的颜色")
}
