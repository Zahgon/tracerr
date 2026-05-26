// Package tracerr makes error output more informative.
// It adds stack trace to error and can display error with source fragments.
//
// Check example of output here https://github.com/ztrue/tracerr
package tracerr

// DefaultCap is a default cap for frames array.
// It can be changed to number of expected frames
// for purpose of performance optimisation.
var DefaultCap = 20

// Error is an error with stack trace.
type Error interface {
	Callers() []uintptr
	Error() string
	StackTrace() []Frame
	Unwrap() error
}

type errorData struct {
	// err contains original error.
	err error
	// pcs contains raw program counters, resolved lazily to frames.
	pcs []uintptr
	// frames contains pre-resolved stack trace.
	frames []Frame
}

// CustomError creates an error with provided frames.
func CustomError(err error, frames []Frame) Error { _ = "STUB: not implemented"; return *new(Error) }

// CustomErrorFromCallers creates an error with provided program counters.
func CustomErrorFromCallers(err error, pcs []uintptr) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

// Errorf creates new error with stacktrace and formatted message.
// Formatting works the same way as in fmt.Errorf.
func Errorf(message string, args ...interface{}) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

// New creates new error with stacktrace.
func New(message string) Error { _ = "STUB: not implemented"; return *new(Error) }

// Wrap adds stacktrace to existing error.
func Wrap(err error) Error { _ = "STUB: not implemented"; return *new(Error) }

// Unwrap returns the original error.
func Unwrap(err error) error { _ = "STUB: not implemented"; return nil }

// Callers returns raw program counters of the stack trace.
func (e *errorData) Callers() []uintptr {
	_ = "STUB: not implemented"

	// Error returns error message.
	return nil
}

func (e *errorData) Error() string { _ = "STUB: not implemented"; return "" }

// StackTrace resolves and returns the stack trace, caching the result.
func (e *errorData) StackTrace() []Frame { _ = "STUB: not implemented"; return nil }

// Unwrap returns the original error.
func (e *errorData) Unwrap() error {
	_ = "STUB: not implemented"

	// Frame is a single step in stack trace.
	return nil
}

type Frame struct {
	// Func contains a function name.
	Func string
	// Line contains a line number.
	Line int
	// Path contains a file path.
	Path string
}

// StackTrace returns stack trace of an error.
// It will be empty if err is not of type Error.
func StackTrace(err error) []Frame { _ = "STUB: not implemented"; return nil }

// String formats Frame to string.
func (f Frame) String() string { _ = "STUB: not implemented"; return "" }

func trace(err error, skip int) Error { _ = "STUB: not implemented"; return *new(Error) }
