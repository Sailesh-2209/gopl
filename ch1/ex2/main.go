package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("argument %2d: %s\n", i+1, arg)
	}
}
