// Package reader provides io.ReadClosers for anything with a URI
package reader

import (
	"fmt"
	"io"
	"os"
)

func Open(uri string) (io.ReadCloser, error) {
	switch {
	case uri == "-":
		return newStdinReader()
	}
	return nil, fmt.Errorf("no handler registered for %q", uri)
}

type readCloser struct {
	io.Reader
}

func (r *readCloser) Close() error { return nil }

func newStdinReader() (io.ReadCloser, error) {
	return &readCloser{os.Stdin}, nil
}
