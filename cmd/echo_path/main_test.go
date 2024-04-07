package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPrintPaths(t *testing.T) {
	testData := []struct {
		name           string
		paths          []string
		expectedOutput string
	}{
		{"Happy path", []string{"/usr/local/bin", "/usr/bin", "/bin"}, "/usr/local/bin\n/usr/bin\n/bin\n"},
		{"Empty path", []string{}, ""},
		{"Path with just one variable", []string{"/usr/bin"}, "/usr/bin\n"},
	}

	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			// Redirect stdout to capture the output
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			printPaths(d.paths)

			// Reset stdout
			w.Close()
			os.Stdout = old

			// Read captured output
			var buf bytes.Buffer
			io.Copy(&buf, r)
			r.Close()

			// Compare output with expected output
			if buf.String() != d.expectedOutput {
				t.Errorf("Expected output: %s, but got: %s", d.expectedOutput, buf.String())
			}
		})
	}
}
