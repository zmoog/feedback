package feedback_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zmoog/feedback"
)

func TestPrintln(t *testing.T) {
	t.Run("Print message to stdout", func(t *testing.T) {
		stdout := bytes.Buffer{}
		stderr := bytes.Buffer{}
		fb := feedback.New(&stdout, &stderr)
		feedback.SetDefault(fb)

		feedback.Println("welcome to this f* world!")

		assert.Equal(t, "welcome to this f* world!\n", stdout.String())
		assert.Equal(t, "", stderr.String())
	})
}
