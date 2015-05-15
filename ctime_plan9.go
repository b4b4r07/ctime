package ctime

import (
	"os"
	"time"
)

func ctime(fi os.FileInfo) time.Time {
	return time.Unix(int64(fi.Sys().(*syscall.Dir).Ctime), 0)
}
