package feedback

import (
	"fmt"
	"io"
	"os"
)

// New returns a new feedback using the given stdout and stderr.
func New(out io.Writer, err io.Writer) Feedback {
	return Feedback{
		out: out,
		err: err,
	}
}

// Default returns the default feedback.
func Default() Feedback {
	return Feedback{
		out: os.Stdin,
		err: os.Stderr,
	}
}

type Feedback struct {
	out io.Writer
	err io.Writer
}

// Println prints the message to stdout.
func (f Feedback) Println(v interface{}) {
	fmt.Fprintln(fb.out, v)
}

// Error prints the error message to stderr.
func (f Feedback) Error(v interface{}) {
	fmt.Fprintln(fb.err, v)
}
