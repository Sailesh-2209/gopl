// source code for exercise 1.7
// this file implements a version of fetch that writes the response of request directly
// to STDOUT by using the io.Copy method, without requiring to read the entire response
// into a memory buffer

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex7: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex7: %v\n", err)
		}
	}
}
