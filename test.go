package main

import (
	"github.com/biezhi/agon/color"
	"github.com/biezhi/agon/log"
	"github.com/biezhi/agon/json"
	"fmt"
	"github.com/biezhi/agon/time"
	"strconv"
)

type Person struct{
	Name string
	Age int
}

func TestColor(){
	color.Println(color.Red, "|| Hello World")
	color.Println(color.Green, "|| Hello World")
	color.Println(color.Yellow, "|| Hello World")
	color.Println(color.Purple, "|| Hello %s", "jack")
}

func TestLog()  {
	log := log.NewLog()
	log.Prefix("[Hello]")
	//log.LogPath("test.log")
	log.Info("Hello Rose")
	log.Debug("Hello %s", "jack")
	log.Warn("Hello %s", "jack")
	log.Trace("Hello %s", "jack")
	log.Error("Hello %s", "jack")
}

func TestJson()  {
	str := "{\"name\":\"jack\", \"age\": 20}"
	json := json.NewJson(str)
	fmt.Println(json.Get("age"))
	fmt.Println(json.Get("name"))
	fmt.Println(json.ToString())

	p := Person{Name:"Rose", Age:20}
	fmt.Println(json.Stringify(p))
}

func TestTime()  {
	log := log.NewLog()
	// 获取当前unix时间戳
	log.Info(strconv.FormatInt(time.UnixTime(), 10))
}

func main() {
	TestColor()
	TestLog()
	//TestJson()
	//TestTime()
}
