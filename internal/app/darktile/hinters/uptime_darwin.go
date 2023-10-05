package hinters

import "time"

// Supposed to return uptime in seconds.  Instead, just to make it build,
// return current time in seconds.
func getUptime() int64 {
	return time.Now().Unix()
}
