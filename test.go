package main

import (
	"github.com/biezhi/agon/color"
)

func main() {
	color.Print(color.Red, "|| Hello World")
	color.Print(color.Green, "|| Hello World")
	color.Print(color.Yellow, "|| Hello World")
	color.Print(color.Purple, "|| Hello %s", "jack")
}
