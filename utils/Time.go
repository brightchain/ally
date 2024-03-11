package utils

import (
	"strconv"
	"time"
)

func FormatDate(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	date := tm.Format("2006-01-02 15:04:05")
	return date
}

func FormatDateByString(timestamp string) string {
	t, _ := strconv.ParseInt(timestamp, 10, 64)
	tm := time.Unix(t, 0)
	date := tm.Format("2006-01-02 15:04:05")
	return date
}
