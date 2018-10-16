// directly derived from the Go io package io.TeeReader
//
// the missing io.TeeReadCloser

package ioaux

import "io"

// TeeReadCloser is a ReadCloser variant of io.TeeReader
//
// see io.TeeReader
func TeeReadCloser(r io.ReadCloser, w io.Writer) io.ReadCloser {
	return &teeReaderCloser{r, w}
}

type teeReaderCloser struct {
	r io.ReadCloser
	w io.Writer
}

func (t *teeReaderCloser) Read(p []byte) (n int, err error) {
	n, err = t.r.Read(p)
	if n > 0 {
		if n, err := t.w.Write(p[:n]); err != nil {
			return n, err
		}
	}
	return
}

func (t *teeReaderCloser) Close() error {
	return t.r.Close()
}
