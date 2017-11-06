//Echo1 prints its command line arguments with program name.
package benchmarkStringRead

import (
	"fmt"
	"os"
	"testing"
)

func TestbenchmarkStringRead(*testing.B) {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
