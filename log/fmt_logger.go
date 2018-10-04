// (C) Copyright 2018 Hewlett Packard Enterprise Development LP.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

/*
	FMT Logger
	- a very simplistic implementation for logger
	- everything is printed through fmt
	- stores history with sensitive data scrubbing
*/
package log

import (
	"fmt"
	"io"
	"regexp"
)

const redactedText = "<REDACTED>"

type FmtLogger struct {
	debug     bool
	outWriter io.Writer
	errWriter io.Writer
	history   []string
}

var (
	// (?s) enables '.' to match '\n' -- see https://golang.org/pkg/regexp/syntax/
	certRegex = regexp.MustCompile("(?s)-----BEGIN CERTIFICATE-----.*-----END CERTIFICATE-----")
	keyRegex  = regexp.MustCompile("(?s)-----BEGIN RSA PRIVATE KEY-----.*-----END RSA PRIVATE KEY-----")
)

func stripSecrets(original []string) []string {
	stripped := []string{}
	for _, line := range original {
		line = certRegex.ReplaceAllString(line, redactedText)
		line = keyRegex.ReplaceAllString(line, redactedText)
		stripped = append(stripped, line)
	}
	return stripped
}

func (l FmtLogger) Debug(args ...interface{}) {
	fmt.Print(args...)
	l.history = append(l.history, fmt.Sprint(args...))
}

func (l FmtLogger) Debugf(fmtString string, args ...interface{}) {
	fmt.Printf(fmtString, args...)
	l.history = append(l.history, fmt.Sprintf(fmtString, args...))
}

func (l FmtLogger) Error(args ...interface{}) {
	fmt.Print(args...)
	l.history = append(l.history, fmt.Sprint(args...))
}

func (l FmtLogger) Errorf(fmtString string, args ...interface{}) {
	fmt.Printf(fmtString, args...)
	l.history = append(l.history, fmt.Sprintf(fmtString, args...))
}

func (l FmtLogger) Info(args ...interface{}) {
	fmt.Print(args...)
	l.history = append(l.history, fmt.Sprint(args...))
}

func (l FmtLogger) Infof(fmtString string, args ...interface{}) {
	fmt.Printf(fmtString, args...)
	l.history = append(l.history, fmt.Sprintf(fmtString, args...))
}

func (l FmtLogger) Warn(args ...interface{}) {
	fmt.Print(args...)
	l.history = append(l.history, fmt.Sprint(args...))
}

func (l FmtLogger) Warnf(fmtString string, args ...interface{}) {
	fmt.Printf(fmtString, args...)
	l.history = append(l.history, fmt.Sprintf(fmtString, args...))
}

func (l FmtLogger) SetDebug(enable bool) {
	l.debug = enable
}

func (l FmtLogger) SetOutWriter(out io.Writer) {
	l.outWriter = out
}

func (l FmtLogger) SetErrWriter(err io.Writer) {
	l.errWriter = err
}

func (l FmtLogger) History() []string {
	return stripSecrets(l.history)
}
