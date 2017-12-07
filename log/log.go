package log

import (
	l "log"

	raven "github.com/getsentry/raven-go"
)

func Fatal(message string) {
	raven.CaptureMessageAndWait(message, nil)
	l.Fatal(message)
}

func Fatalf(format string, err error) {
	raven.CaptureErrorAndWait(err, nil)
	l.Fatal(format, err)
}

func Errorf(format string, network string, v ...interface{}) {
	// v is other arguments to format string, last one is always the error
	err := v[len(v)-1].(error)
	raven.CaptureError(err, map[string]string{"network": network})
	l.Printf(format, v...)
}

func Printf(format string, v ...interface{}) {
	l.Printf(format, v...)
}
