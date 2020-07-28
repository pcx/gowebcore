package errors

import (
	"fmt"
	"runtime"

	"github.com/pcx/gowebcore/logs"
)

const maxStacktraceDepth = 25

type ErrorCode int

type Error struct {
	error
	code       ErrorCode
	frames     []runtime.Frame
	reportable bool
	statusCode int // Use this for HTTP Status codes for failures caused by this error
}

func getFrames() []runtime.Frame {
	frames := make([]runtime.Frame, 0)

	// Ask runtime.Callers for up to 10 pcs
	pc := make([]uintptr, maxStacktraceDepth)
	n := runtime.Callers(3, pc)
	if n == 0 {
		// No pcs available. Stop now.
		// This can happen if the first argument to runtime.Callers is large.
		return nil
	}

	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	runtimeFrames := runtime.CallersFrames(pc)
	for {
		if runtimeFrames == nil {
			break
		}

		runtimeFrame, more := runtimeFrames.Next()
		frames = append(frames, runtimeFrame)
		if !more {
			break
		}
	}
	return frames
}

func New(error error) Error {
	if error == nil {
		panic("Cannot create Error with error = nil")
	}

	return Error{
		error:      error,
		frames:     getFrames(),
		reportable: true,
		statusCode: 500,
	}
}

func Format(message string, a ...interface{}) Error {
	return Error{
		error:      fmt.Errorf(message, a...),
		frames:     getFrames(),
		reportable: true,
		statusCode: 500,
	}
}

func build(err error, code ErrorCode, reportable bool, statusCode int) Error {
	return Error{
		error:      err,
		code:       code,
		frames:     getFrames(),
		reportable: reportable,
		statusCode: statusCode,
	}
}

func Nil() Error {
	return Error{error: nil}
}

func (e Error) Stack() []runtime.Frame {
	return e.frames
}

func (e Error) Present() bool {
	return e.error != nil
}

func (e Error) Is(code ErrorCode) bool {
	return e.Present() && e.code == code
}

func (e Error) IsRecordNotFound() bool {
	return e.Present() && e.code == RecordNotFoundCode
}

func (e Error) Code() ErrorCode {
	return e.code
}

func (e Error) Error() string {
	if !e.Present() {
		panic("Error.error should not be nil")
	}
	return e.error.Error()
}

func (e Error) Report() {
	if e.IsReportable() {
		// Use this to report the error to your error management service
		// Use error.frames for getting stacktrace details
		// rollbar.Error(e)
	}

	logger := logs.ErrLogger()
	logger.Printf("ERROR: %s", e.Error())

	// Loop to get frames.
	for _, frame := range e.frames {
		logger.Printf("%s:%d (%s)", frame.File, frame.Line, frame.Func.Name())
	}
}

func (e Error) IsReportable() bool {
	return e.reportable
}

// ErrorClass implements newrelic.ErrorClasser
func (e Error) ErrorClass() string {
	return fmt.Sprintf("%d", e.code)
}

func (e Error) StatusCode() int {
	return e.statusCode
}
