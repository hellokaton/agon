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
	_TimePattern string
	_LogPath     string
	_LogLevel    Level
	_LogColor    map[Level]int
	_TextColor   int
	_Prefix      string
}

func (log *Logger) Prefix(Prefix string) *Logger {
	log._Prefix = Prefix
	return log
}

func (log *Logger) LogPath(LogPath string) *Logger {
	log._LogPath = LogPath
	return log
}

func (log *Logger) TimePattern(TimePattern string) *Logger {
	log._TimePattern = TimePattern
	return log
}

func (log *Logger) TextColor(TextColor int) *Logger {
	log._TextColor = TextColor
	return log
}

func (log *Logger) LogColor(LogColor map[Level]int) *Logger {
	log._LogColor = LogColor
	return log
}

// log level
const (
	LEVEL_DEBUG = iota
	LEVEL_TRACE
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
)

// config logger
func NewLog() *Logger {
	log := &Logger{}
	log._TimePattern = "01/02 15:04:05.000"
	log._LogLevel = LEVEL_DEBUG
	return log
}

// print info log
func (log *Logger) Info(format string, v ...interface{}) {
	if log._LogLevel <= LEVEL_INFO {
		log.Log(os.Stdout, LEVEL_INFO, fmt.Sprintf(format, v...))
	}
}

// print trace log
func (log *Logger) Trace(format string, v ...interface{}) {
	if log._LogLevel <= LEVEL_TRACE {
		log.Log(os.Stdout, LEVEL_TRACE, fmt.Sprintf(format, v...))
	}
}

// print debug log
func (log *Logger) Debug(format string, v ...interface{}) {
	if log._LogLevel <= LEVEL_DEBUG {
		log.Log(os.Stdout, LEVEL_DEBUG, fmt.Sprintf(format, v...))
	}
}

// print warning log
func (log *Logger) Warn(format string, v ...interface{}) {
	if log._LogLevel <= LEVEL_WARN {
		log.Log(os.Stdout, LEVEL_WARN, fmt.Sprintf(format, v...))
	}
}

// print error log
func (log *Logger) Error(format string, v ...interface{}) {
	if log._LogLevel <= LEVEL_ERROR {
		log.Log(os.Stdout, LEVEL_ERROR, fmt.Sprintf(format, v...))
	}
}

// get log level text
func (log *Logger) LevelText(level Level) string {
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
func (log *Logger) Color(str string, level Level) string {
	if c, ok := log._LogColor[level]; ok {
		return color.Color(c, str)
	}
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
func (log *Logger) TimeLog() string {
	return time.Now().Format(log._TimePattern)
}

// print log
func (log *Logger) Log(out io.Writer, level Level, txt string) {
	prefix := fmt.Sprintf("%s[%s] %s\t=> ", log._Prefix, log.TimeLog(), log.LevelText(level))
	out.Write([]byte(log.Color(prefix, level)))
	out.Write([]byte(txt + "\n"))
	log.store(prefix + txt)
}

// store log
func (log *Logger) store(txt string) {
	if !strings.EqualFold("", log._LogPath) {
		os.OpenFile(log._LogPath, os.O_CREATE, 0711)
		fout, err := os.OpenFile(log._LogPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println(color.Color(color.Red, log._LogPath+err.Error()))
			return
		}
		defer fout.Close()
		if log._TextColor > 0 {
			fout.WriteString(color.Color(log._TextColor, txt+"\n"))
		} else {
			fout.WriteString(txt + "\n")
		}
	}
}
