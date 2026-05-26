package tracerr

import (
	"sync"
)

// DefaultLinesAfter is number of source lines after traced line to display.
var DefaultLinesAfter = 2

// DefaultLinesBefore is number of source lines before traced line to display.
var DefaultLinesBefore = 3

var cache = map[string][]string{}

var mutex sync.RWMutex

// Print prints error message with stack trace.
func Print(err error) { _ = "STUB: not implemented"; return }

// PrintSource prints error message with stack trace and source fragments.
//
// By default, 6 lines of source code will be printed,
// see DefaultLinesAfter and DefaultLinesBefore.
//
// Pass a single number to specify a total number of source lines.
//
// Pass two numbers to specify exactly how many lines should be shown
// before and after traced line.
func PrintSource(err error, nums ...int) { _ = "STUB: not implemented"; return }

// PrintSourceColor prints error message with stack trace and source fragments,
// which are in color.
// Output rules are the same as in PrintSource.
func PrintSourceColor(err error, nums ...int) { _ = "STUB: not implemented"; return }

// Sprint returns error output by the same rules as Print.
func Sprint(err error) string { _ = "STUB: not implemented"; return "" }

// SprintSource returns error output by the same rules as PrintSource.
func SprintSource(err error, nums ...int) string { _ = "STUB: not implemented"; return "" }

// SprintSourceColor returns error output by the same rules as PrintSourceColor.
func SprintSourceColor(err error, nums ...int) string { _ = "STUB: not implemented"; return "" }

func calcRows(nums []int) (before, after int, withSource bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

// Extra line goes to "before" rather than "after".

func readLines(path string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func sourceRows(rows []string, frame Frame, before, after int, colorized bool) []string {
	_ = "STUB: not implemented"
	return nil
}

func sprint(err error, nums []int, colorized bool) string { _ = "STUB: not implemented"; return "" }
