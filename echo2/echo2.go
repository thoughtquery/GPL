// Echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	s, seperator := "", ""
	for _, argument := range os.Args[1:] {
		s += seperator + argument
		sperator = " "
	}
	fmt.Println(s)
}
