package main

import (
	"fmt"

	"github.com/ztrue/tracerr"
)

func main() {
	if err := read(); err != nil {
		// Dump raw stack trace.
		frames := tracerr.StackTrace(err)
		fmt.Printf("%#v\n", frames)
	}
}

func read() error { _ = "STUB: not implemented"; return nil }

func readNonExistent() error { _ = "STUB: not implemented"; return nil }
