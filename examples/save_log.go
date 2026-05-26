package main

import (
	"os"

	"github.com/ztrue/tracerr"
)

func main() {
	if err := read(); err != nil {
		// Save output to variable.
		text := tracerr.SprintSource(err)
		os.WriteFile("/tmp/tracerr.log", []byte(text), 0644)
	}
}

func read() error { _ = "STUB: not implemented"; return nil }

func readNonExistent() error { _ = "STUB: not implemented"; return nil }
