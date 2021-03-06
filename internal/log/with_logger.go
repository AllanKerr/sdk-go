// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package log

import (
	"go.temporal.io/sdk/log"
)

// With returns Logger instance that prepend every log entry with keyvals. If logger implments WithLogger it is used, otherwise every log call will be intercepted.
func With(logger log.Logger, keyvals ...interface{}) log.Logger {
	if wl, ok := logger.(log.WithLogger); ok {
		return wl.With(keyvals...)
	}

	return newWithLogger(logger, keyvals...)
}

type withLogger struct {
	logger  log.Logger
	keyvals []interface{}
}

func newWithLogger(logger log.Logger, keyvals ...interface{}) *withLogger {
	return &withLogger{logger: logger, keyvals: keyvals}
}

func (l *withLogger) prependKeyvals(keyvals []interface{}) []interface{} {
	return append(l.keyvals, keyvals...)
}

// Debug writes message to the log.
func (l *withLogger) Debug(msg string, keyvals ...interface{}) {
	l.logger.Debug(msg, l.prependKeyvals(keyvals)...)
}

// Info writes message to the log.
func (l *withLogger) Info(msg string, keyvals ...interface{}) {
	l.logger.Info(msg, l.prependKeyvals(keyvals)...)
}

// Warn writes message to the log.
func (l *withLogger) Warn(msg string, keyvals ...interface{}) {
	l.logger.Warn(msg, l.prependKeyvals(keyvals)...)
}

// Error writes message to the log.
func (l *withLogger) Error(msg string, keyvals ...interface{}) {
	l.logger.Error(msg, l.prependKeyvals(keyvals)...)
}
