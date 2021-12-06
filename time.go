package gocom

import (
	"time"
)

func NowUTC() time.Time {
	return time.Now().UTC()
}

func Now() time.Time {
	return time.Now()
}

func TimeString(t time.Time) string {
	return t.Format(time.RFC3339Nano)
}
