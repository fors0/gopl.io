//Benchmark echo implementations
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo2() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3() string {
	return strings.Join(os.Args[1:], " ")
}

func main() {
	start := time.Now()
	echo1()
	fmt.Printf("Time elapsed for echo1: %2f\n", time.Since(start).Seconds())
	start = time.Now()
	echo2()
	fmt.Printf("Time elapsed for echo2: %2f\n", time.Since(start).Seconds())
	start = time.Now()
	echo3()
	fmt.Printf("Time elapsed for echo3: %2f\n", time.Since(start).Seconds())
}
