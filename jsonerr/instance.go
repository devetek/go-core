package jsonerr

import (
	"fmt"
	"runtime"
	"time"

	"go.uber.org/zap"
)

/**
 * instance is a struct to store error data
 */
type instance struct {
	alwaysPrint bool // indicate if error should always print to stderr
	caller      bool
	message     string
	err         error
	zap         *zap.Logger
}

// create new instance of JSONError
func New(options ...Options) Error {
	logger, _ := zap.NewProduction(
		zap.WithCaller(false),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	defer logger.Sync()

	var e = &instance{
		alwaysPrint: false,
		caller:      false,
		message:     "",
		err:         nil,
		zap:         logger,
	}

	// set options
	for _, opt := range options {
		opt(e)
	}

	return e
}

// Create log error data to be processed
func (e *instance) Create(error NewError) {
	e.err = error.Error
	e.message = error.Message
}

// Error returns the error message
func (e *instance) Error() string {
	if e.caller {
		_, filename, line, ok := runtime.Caller(1)
		if ok {
			if e.message != "" {
				return fmt.Sprintf("[%s] %s: %s:%d", e.datetime(), e.message, filename, line)
			}

			if e.err != nil {
				return fmt.Sprintf("[%s] %s: %s:%d", e.err.Error(), e.message, filename, line)
			}
		}
	}

	if e.message != "" {
		return fmt.Sprintf("[%s] %s", e.datetime(), e.message)
	}

	if e.err != nil {
		return fmt.Sprintf("[%s] %s", e.datetime(), e.err.Error())
	}

	return ""
}

// Public interface to call print error to stderr
func (e *instance) Print() {
	e.print()
}

// Print error to stdout with zap logger
func (e *instance) print() {
	if e.caller {
		_, filename, line, ok := runtime.Caller(1)
		if ok {
			if e.message != "" {
				e.zap.Error(e.message,
					zap.String("datetime", e.datetime()),
					zap.String("caller", fmt.Sprintf("%s:%d", filename, line)),
				)

				return
			}

			if e.err != nil {
				e.zap.Error(e.err.Error(),
					zap.String("datetime", e.datetime()),
					zap.String("caller", fmt.Sprintf("%s:%d", filename, line)),
				)

				return
			}
		}

	}

	if e.message != "" {
		e.zap.Error(e.message,
			zap.String("datetime", e.datetime()),
		)
		return
	}

	if e.err != nil {
		e.zap.Error(e.err.Error(),
			zap.String("datetime", e.datetime()),
		)
		return
	}
}

// Unwrap returns the wrapped error
func (e *instance) Unwrap() error {
	// always print error to stderr if condition is met
	if e.alwaysPrint {
		e.print()
	}

	return e.err
}

// Get current datetime with format YYYY-MM-DD HH:MM:SS
func (e *instance) datetime() string {
	now := time.Now() // get this early.
	year, month, day := now.Date()
	hour, min, sec := now.Clock()

	// TODO: use concatenation for better performance
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, min, sec)
}
