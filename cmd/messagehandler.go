package cmd

import (
	"fmt"
	"strings"
	"time"
)

func isPrintable(str string) bool {
	for _, r := range str {
		if r < 32 || r == 127 {
			return false // Unprintable ASCII characters
		}
		if r >= 0x80 {
			return false // ANSI escape codes
		}
	}
	return true
}

func GetTimeFromMsg(msg string) time.Time {
	parts := strings.SplitN(msg, "][", 2)
	if len(parts) < 2 {
		return time.Time{}
	}
	t, err := time.Parse("2006-01-02 15:04:05", strings.Trim(parts[0], "["))
	if err != nil {
		return time.Time{}
	}
	return t
}

func GetFormatTime() string {
	now := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

func CheckMsg(str string) bool {
	return len(strings.TrimSpace(str)) == 0 || str == "" || !isPrintable(str)
}
