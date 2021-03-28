//Print the index and value of each of its arguments one per line
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args[1:] {
		fmt.Printf("Index: %d, Value: %string\n", index+1, value)
	}
}
