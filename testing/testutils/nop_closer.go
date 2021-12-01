// nolint: nilerr wrapcheck
package testutils

import (
	"io"
)

type nopCloser struct {
	wr  io.Reader
	err error
}

func (w *nopCloser) Read(p []byte) (int, error) {
	if w.err != nil {
		return 0, w.err
	}

	return w.wr.Read(p)
}
func (w *nopCloser) Close() error {
	return nil
}

// NopCloser transforms io.Reader to io.ReadCloser.
func NopCloser(wr io.Reader) io.ReadCloser {
	return &nopCloser{wr: wr}
}

// ErrorCloser return io.ReadCloser.
// It return passed error when Read method will be called.
func ErrorCloser(err error) io.ReadCloser {
	return &nopCloser{
		err: err,
	}
}
