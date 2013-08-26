// cat consumes the input from a reader specified on the command line
// and outputs it to stdout.
//
// Examples:
//
// Read from standard in:
// 	$ echo "foo" | ./cat -
// 	foo
//
// Cat the contents of a remote url:
//
// 	$ ./cat http://dave.cheney.net/
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/davecheney/reader"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %v $URI\n", os.Args[0])
		os.Exit(2)
	}
	r, err := reader.Open(os.Args[1])
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer r.Close()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
