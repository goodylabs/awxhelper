package prompter

import (
	"io"
)

type noBellWriter struct {
	w io.Writer
}

func (n noBellWriter) Write(p []byte) (int, error) {
	filtered := make([]byte, 0, len(p))
	for _, b := range p {
		if b != 0x07 {
			filtered = append(filtered, b)
		}
	}
	return n.w.Write(filtered)
}

func (n noBellWriter) Close() error {
	return nil
}
