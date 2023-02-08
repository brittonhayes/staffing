package logging

import "github.com/ThreeDotsLabs/watermill"

// Logger is a wrapper around watermill.LoggerAdapter that provides a
// default implementation of the LoggerAdapter interface.
type Logger struct {
	watermill.LoggerAdapter
}

// New returns a new Logger instance.
func New(debug bool, trace bool, domain string) *Logger {
	return &Logger{
		LoggerAdapter: watermill.NewStdLogger(debug, trace).With(watermill.LogFields{
			"domain": domain,
		}),
	}
}
