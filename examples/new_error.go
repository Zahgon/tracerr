package main

import (
	"github.com/ztrue/tracerr"
)

func main() {
	if err := foo(); err != nil {
		tracerr.PrintSourceColor(err)
	}
}

func foo() error { _ = "STUB: not implemented"; return nil }

func bar(i int) error {
	_ = "STUB: not implemented"

	// Create new error with stack trace.
	return nil
}
