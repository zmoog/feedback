package feedback

var (
	fb = Default()
)

// SetDefault sets the default Feedback for the whole feedback package.
func SetDefault(f Feedback) {
	fb = f
}

// Println prints the message to stdout.
func Println(v interface{}) {
	fb.Println(v)
}

// Error prints the error message to stderr.
func Error(v interface{}) {
	fb.Error(v)
}
