package testutils

import "io"

type nopCloser struct {
	wr io.Reader
}

func (w *nopCloser) Read(p []byte) (n int, err error) {
	return w.wr.Read(p)
}
func (w *nopCloser) Close() error {
	return nil
}

// NopCloser transforms io.Reader to io.ReadCloser.
func NopCloser(wr io.Reader) io.ReadCloser {
	return &nopCloser{wr: wr}
}
