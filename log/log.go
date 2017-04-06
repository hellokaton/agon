package log

import (
	"io"
	"fmt"
	"os"
	"github.com/biezhi/agon/color"
	"time"
	"strings"
)

type Level int

type Logger struct {
	TimePattern string
	LogPath     string
	LogLevel    Level
	LogColor    int
}

// log level
const (
	LEVEL_DEBUG = iota
	LEVEL_TRACE
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
)

var _logger Logger = Logger{TimePattern: "01/02 15:04:05.000", LogLevel: LEVEL_DEBUG}

// config logger
func ConfigLog(log Logger) {
	_logger = log
}

// print info log
func Info(format string, v ...interface{}) {
	if _logger.LogLevel <= LEVEL_INFO {
		Log(os.Stdout, LEVEL_INFO, fmt.Sprintf(format, v...))
	}
}

// print trace log
func Trace(format string, v ...interface{}) {
	if _logger.LogLevel <= LEVEL_TRACE {
		Log(os.Stdout, LEVEL_TRACE, fmt.Sprintf(format, v...))
	}
}

// print debug log
func Debug(format string, v ...interface{}) {
	if _logger.LogLevel <= LEVEL_DEBUG {
		Log(os.Stdout, LEVEL_DEBUG, fmt.Sprintf(format, v...))
	}
}

// print warning log
func Warn(format string, v ...interface{}) {
	if _logger.LogLevel <= LEVEL_WARN {
		Log(os.Stdout, LEVEL_WARN, fmt.Sprintf(format, v...))
	}
}

// print error log
func Error(format string, v ...interface{}) {
	if _logger.LogLevel <= LEVEL_ERROR {
		Log(os.Stdout, LEVEL_ERROR, fmt.Sprintf(format, v...))
	}
}

// get log level text
func LevelText(level Level) string {
	switch {
	case level == LEVEL_INFO:
		return "INFO"
	case level == LEVEL_DEBUG:
		return "DEBUG"
	case level == LEVEL_TRACE:
		return "TRACE"
	case level == LEVEL_WARN:
		return "WARN"
	case level == LEVEL_ERROR:
		return "ERROR"
	default:
		return ""
	}
}

// get log level color
func Color(str string, level Level) string {
	switch {
	case level == LEVEL_INFO:
		return color.Color(color.Green, str)
	case level == LEVEL_DEBUG:
		return color.Color(color.LightGray, str)
	case level == LEVEL_TRACE:
		return color.Color(color.LightPurple, str)
	case level == LEVEL_WARN:
		return color.Color(color.Yellow, str)
	case level == LEVEL_ERROR:
		return color.Color(color.Red, str)
	default:
		return color.Color(color.Cyan, str)
	}
}

// return time log
func TimeLog() string {
	return time.Now().Format(_logger.TimePattern)
}

// print log
func Log(out io.Writer, level Level, txt string) {
	prefix := fmt.Sprintf("[%s] %s\t=> ", TimeLog(), LevelText(level))
	out.Write([]byte(Color(prefix, level)))
	out.Write([]byte(txt + "\n"))
	store(prefix + txt)
}

// store log
func store(txt string) {
	if !strings.EqualFold("", _logger.LogPath) {
		os.OpenFile(_logger.LogPath, os.O_CREATE, 0711)
		fout, err := os.OpenFile(_logger.LogPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println(color.Color(color.Red, _logger.LogPath+err.Error()))
			return
		}
		defer fout.Close()
		if _logger.LogColor > 0 {
			fout.WriteString(color.Color(_logger.LogColor, txt+"\n"))
		} else {
			fout.WriteString(txt + "\n")
		}
	}
}
