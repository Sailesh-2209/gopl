package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, arg := range os.Args[1:] {
		data, err := os.ReadFile(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, cnt := range counts {
		if cnt > 1 {
			fmt.Printf("%d\t%s\n", cnt, line)
		}
	}
}
