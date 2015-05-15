// Package ctime provides a platform-independent way to get ctimes for files.
package ctime

import (
	"os"
	"time"
)

// Get returns the Last Access Time for the given FileInfo
func Get(fi os.FileInfo) time.Time {
	return ctime(fi)
}

// Stat returns the Last Access Time for the given filename
func Stat(name string) (time.Time, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return time.Time{}, err
	}
	return ctime(fi), nil
}
