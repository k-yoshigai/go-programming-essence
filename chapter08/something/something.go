package something

import (
	"fmt"
)

func MakeSomething1(n int) []string {
	var r []string
	for i := 0; i < n; i++ {
		r = append(r, fmt.Sprintf("%05d something", i))
	}
	return r
}

func MakeSomething2(n int) []string {
	r := make([]string, n)

	for i := 0; i < n; i++ {
		r[i] = fmt.Sprintf("%05d something", i)
	}
	return r
}
