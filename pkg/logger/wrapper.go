package logger

import (
	"errors"
	"io"

	"github.com/sirupsen/logrus"
)

// logrusWrapper is a local wrapper around logrus
type logrusWrapper struct {
	*logrus.Entry
}

func (w logrusWrapper) Writer() *io.PipeWriter {
	return w.Entry.Writer()
}

// Log is a wrapper around logrus.Log
func (w logrusWrapper) Log(level Level, args ...interface{}) {
	w.Entry.Log(logrus.Level(level), args...)
}

// Logf is a wrapper around logrus.Logf
func (w logrusWrapper) Logf(level Level, format string, args ...interface{}) {
	w.Entry.Logf(logrus.Level(level), format, args...)
}

// Logln is a wrapper around logrus.Logln
func (w logrusWrapper) Logln(level Level, args ...interface{}) {
	w.Entry.Logln(logrus.Level(level), args...)
}

// WithFields is a wrapper around logrus.WithFields
func (w logrusWrapper) WithField(key string, value interface{}) Logger {
	return logrusWrapper{w.Entry.WithField(key, value)}
}
func (w logrusWrapper) WithFields(fields Fields) Logger {
	logrusFields := logrus.Fields(fields)
	return logrusWrapper{w.Entry.WithFields(logrusFields)}
}

// WithError appends to logger 'error' key
func (w logrusWrapper) WithError(err error) Logger {
	return w.WithFields(Fields{ErrorField: err})
}

// WrapLogrusEntry wrap logrus logger to our interface
func WrapLogrusEntry(l *logrus.Entry) (Logger, error) {
	if l != nil {
		return &logrusWrapper{l}, nil
	}
	return nil, errors.New("can't wrap nil")

}

// WrapLogrusLogger wrap logrus entry to our interface
func WrapLogrusLogger(l *logrus.Logger) (Logger, error) {
	if l != nil {
		return &logrusWrapper{l.WithField("wrapped", "wrapped")}, nil
	}
	return nil, errors.New("can't wrap nil")

}
