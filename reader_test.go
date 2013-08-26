package reader

import "testing"

func TestOpenStdin(t *testing.T) {
	r, err := Open("-")
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Close(); err != nil {
		t.Fatal(err)
	}
}
