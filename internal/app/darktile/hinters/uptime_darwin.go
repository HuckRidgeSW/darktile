package hinters

import (
	"time"

	"golang.org/x/sys/unix"
)

func getUptime() int64 {
	tv, err := unix.SysctlTimeval("kern.boottime")
	if err != nil {
		panic(err)
	}
	sec, nsec := tv.Unix()
	return int64(time.Now().Sub(time.Unix(sec, nsec)) / time.Second)
}
