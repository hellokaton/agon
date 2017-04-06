package time

import (
	"time"
	"strings"
)

// return unix time
func UnixTime() int64 {
	return time.Now().Unix()
}

// format date
func Format(_time int64, pattern string) string {
	if strings.EqualFold("", pattern){
		pattern = "2006-01-02 15:04:05"
	}
	return time.Unix(_time, 0).Format(pattern)
}