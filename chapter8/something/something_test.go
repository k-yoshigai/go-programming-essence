package something_test

import (
	"testing"

	"github.com/k-yoshigai/go-programming-essence/chapter8/something"
)

func BenchmarkMakeSomething(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = something.MakeSomething(1000)
	}
}
