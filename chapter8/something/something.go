package something

import (
	"fmt"
)

func MakeSomething(n int) []string {
	var r []string
	for i := 0; i < n; i++ {
		r = append(r, fmt.Sprintf("%05d something", i))
	}
	return r
}
