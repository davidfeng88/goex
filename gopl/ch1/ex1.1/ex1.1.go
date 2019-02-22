// Based on echo3
// Prints all of its arguments, including the first one,
// the command itself

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
