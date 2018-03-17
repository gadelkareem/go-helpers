package helpers

import (
	"math/rand"
	"time"
	"crypto/md5"
	"encoding/hex"
	"syscall"
	"unicode/utf8"
	"strings"
	"html"
	"github.com/astaxie/beego/logs"
)

func RandomNumber(min int, max int) int {
	if min == max {
		return min
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func Md5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func LiftRLimits() (rLimit syscall.Rlimit) {
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		logs.Error("Error Getting Rlimit %v", err)
	}
	rLimit.Cur = rLimit.Max
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		logs.Error("Error Setting Rlimit %v", err)
	}
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		logs.Error("Error Getting Rlimit %v", err)
	}

	return
}

type NoRetryError struct {
	err error
}

func NewNoRetryError(err error) NoRetryError {
	return NoRetryError{err}
}
func (e NoRetryError) Error() string { return e.err.Error() }

func Retry(retryFunc func() error, maxRetries int) (err error) {
	retry := 0
	for retry < maxRetries {
		err = retryFunc()
		if err == nil {
			return
		}
		if nre, isNoRetry := err.(NoRetryError); isNoRetry {
			return nre.err
		}
		retry++
	}
	return
}

func Utf8Encode(s string) string {
	if utf8.ValidString(s) {
		return s
	}
	v := make([]rune, 0, len(s))
	for i, r := range s {
		if r == utf8.RuneError {
			_, size := utf8.DecodeRuneInString(s[i:])
			if size == 1 {
				continue
			}
		}
		v = append(v, r)
	}
	return string(v)
}

func CleanString(s string) string {
	s = Utf8Encode(strings.TrimSpace(html.UnescapeString(s)))
	s = strings.Replace(s, "\\", "", -1)
	return s
}

var slugReplacer = strings.NewReplacer(" ", "-", "\t", "-", "/", "-", "\\", "-")

func Slug(s string) string {
	return slugReplacer.Replace(s)
}

func SubString(s string, l int) string {
	if len(s) <= l {
		return s
	}
	return s[0:l]
}
