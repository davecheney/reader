// Package reader provides io.ReadClosers for anything with a URI.
// The goal of this package is to enable command line applications to
// consume input from a wide variety of sources without having to
// implement the details of the specific input source.
//
// Currently supported schemes are:
//
// 	-	stdin
// 	http://	http 2xx responses
package reader

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gorilla/http"
)

// Open returns an io.ReadCloser representing the contents of the
// source specified by a uri.
func Open(uri string) (io.ReadCloser, error) {
	switch {
	case uri == "-":
		return newStdinReader()
	case strings.HasPrefix(uri, "http://"), strings.HasPrefix(uri, "https://"):
		return newHttpReader(uri)
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

func newHttpReader(uri string) (io.ReadCloser, error) {
	status, _, body, err := http.DefaultClient.Get(uri, nil)
	if err != nil {
		return nil, err
	}
	if !status.IsSuccess() {
		return nil, &http.StatusError{status}
	}
	return body, nil
}
