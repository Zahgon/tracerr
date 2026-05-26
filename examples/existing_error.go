package main

import (
	"github.com/ztrue/tracerr"
)

func main() {
	if err := read(); err != nil {
		tracerr.PrintSourceColor(err)
	}
}

func read() error { _ = "STUB: not implemented"; return nil }

func readNonExistent() error { _ = "STUB: not implemented"; return nil }

// Add stack trace to existing error, no matter if it's nil.
