package ctime

import (
	"os"
	"syscall"
	"time"
)

func timespecToTime(sec, nsec int64) time.Time {
	return time.Unix(sec, nsec)
}

func ctime(fi os.FileInfo) time.Time {
	st := fi.Sys().(*syscall.Stat_t)
	return timespecToTime(st.Ctime, st.CtimeNsec)
}
