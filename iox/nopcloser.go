// Package iox
// Date: 2023/4/10 15:41
// Author: Amu
// Description:
package iox

import "io"

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error {
	return nil
}

func NopCloser(w io.Writer) io.WriteCloser {
	return nopCloser{w}
}
