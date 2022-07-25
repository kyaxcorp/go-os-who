package who

import (
	"runtime"
	"time"
)

type User struct {
	Name    string
	Time    time.Time
	PID     int64
	Comment string
	Idle    string // todo:
	Line    string
}

func List() ([]User, error) {
	switch runtime.GOOS {
	case "linux":
		return linux()
	case "windows":
	}

	// empty...
	return []User{}, nil
}
