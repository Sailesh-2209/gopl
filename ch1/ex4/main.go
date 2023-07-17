package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	type obj struct {
		lineNum int
		file    string
	}
	counts := make(map[string][]obj)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex4: %v\n", err)
			continue
		}
		for i, line := range strings.Split(string(data), "\n") {
			counts[line] = append(counts[line], obj{lineNum: i + 1, file: filename})
		}
	}
	for line, instArr := range counts {
		if len(instArr) > 1 {
			fmt.Printf("* The line %s appears %d times. The files in which it appears are:\n", line, len(instArr))
			for _, inst := range instArr {
				fmt.Printf("\t- %s [%d]\n", inst.file, inst.lineNum)
			}
		}
	}
}
