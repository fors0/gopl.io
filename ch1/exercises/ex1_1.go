//Prints also the command that invoked it
package main

import (
	"fmt"
	"os"
	"strings"
)

// func main() {
// 	var s, separator string
// 	for i := 0; i < len(os.Args); i++ {
// 		s += separator + os.Args[i]
// 		separator = " "
// 	}
// 	fmt.Println(s)
// }

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
