package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	s1, sep := "", " "
	for _, arg := range os.Args[1:] {
		s1 += arg + sep
	}
	fmt.Println(s1)
}

func echo2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {

	start := time.Now()
	echo1()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Time taken by echo approach 1: %d\n", elapsed)

	start = time.Now()
	echo2()
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("Time taken by echo approach 2: %d\n", elapsed)
}
