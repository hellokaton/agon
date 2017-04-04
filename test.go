package main

import (
	"github.com/biezhi/agon/color"
)

func main() {
	color.Println(color.Red, "|| Hello World")
	color.Println(color.Green, "|| Hello World")
	color.Println(color.Yellow, "|| Hello World")
	color.Print(color.Purple, "|| Hello %s", "jack")
}
