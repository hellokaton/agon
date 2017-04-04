# agon

`agon` 这个名字起源于希腊神话中的竞赛之神，胜利女神。这个项目存储了Golang开发常用的类库。

## 安装

```bash
go get github.com/biezhi/agon
```

## 使用

### 输出彩色文本

```go
color.Println(color.Red, "|| Hello World")
color.Println(color.Green, "|| Hello World")
color.Println(color.Yellow, "|| Hello World")
color.Println(color.Purple, "|| Hello %s", "jack")
```

### 日志

```go
//log.ConfigLog("test.log")
log.Info("Hello Rose")
log.Debug("Hello %s", "jack")
log.Warn("Hello %s", "jack")
log.Trace("Hello %s", "jack")
log.Error("Hello %s", "jack")
```

### JSON解析


#### 加载一个配置文件

```go
configMap := LoadJson("config.json")
fmt.Println(configMap["api_key"])
```

### 字符串转换为JSON对象

```go
str := "{\"name\":\"jack\", \"age\": 20}"
json := json.NewJson(str)
fmt.Println(json.Get("age"))
fmt.Println(json.Get("name"))
fmt.Println(json.ToString())
```

### 类型转换为JSON

```go
p := Person{Name:"Rose", Age:20}
fmt.Println(json.Stringify(p))
```
