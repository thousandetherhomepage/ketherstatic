package log

import (
	l "log"

	raven "github.com/getsentry/raven-go"
)

func Errorf(format string, network string, v ...interface{}) {
	// v is other arguments to format string, last one is always the error
	err := v[len(v)-1].(error)
	raven.CaptureError(err, map[string]string{"network": network})
	l.Printf(format, v...)
}

func Printf(format string, v ...interface{}) {
	l.Printf(format, v...)
}
