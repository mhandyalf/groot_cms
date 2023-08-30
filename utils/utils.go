package utils

import (
	"fmt"
	"time"
)

func LogRequest(method, path string, start time.Time) {
	fmt.Printf("[%s] - HTTP request sent to %s %s\n", start.Format("2006/01/02 15:04:05"), method, path)
}
