//Echo1 prints its command line arguments with program name.
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var s, sep string
	start := time.Now()
	for i := 0; i < len(os.Args); i++ {
		s += sep + strconv.Itoa(i) + ", " + os.Args[i] + "\n"
	}
	fmt.Print(s)
	fmt.Print(time.Since(start).Seconds(), "\n")
}
