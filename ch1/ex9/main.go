package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex9: %v\n", err)
			os.Exit(1)
		}
		data, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex9: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Response status code: %d\n", res.StatusCode)
		fmt.Printf("Response data: %s\n", data)
	}
}
