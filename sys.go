// +build linux

package h

import (
	"io"
	"os"
	"syscall"
)

func LiftRLimits() (rLimit syscall.Rlimit, err error) {
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	rLimit.Cur = rLimit.Max
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &syscall.Rlimit{Cur: 1048576, Max: rLimit.Max})
		if err != nil {
			return
		}
	}
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	return
}


type Flock struct {
	*os.File
}

func NewFlock(path string) (*Flock, error) {
	f, err := os.OpenFile(path, syscall.O_CREAT|syscall.O_RDWR|syscall.O_CLOEXEC, 0666)
	if err != nil {
		return nil, err
	}

	err = syscall.FcntlFlock(
		f.Fd(),
		syscall.F_SETLK,
		&syscall.Flock_t{
			Type:   syscall.F_WRLCK,
			Whence: io.SeekStart,
			Start:  0,
			Len:    0,
		})
	if err != nil {
		return nil, err
	}

	return &Flock{f}, nil
}

func (f *Flock) UnLock() error {
	err := syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	if err != nil {
		return err
	}
	return f.Close()
}
