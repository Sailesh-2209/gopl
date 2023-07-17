package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for txt, cnt := range counts {
		if cnt > 1 {
			fmt.Printf("%d\t%s\n", cnt, txt)
		}
	}
}
