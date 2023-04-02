package something_test

import (
	"testing"

	"github.com/k-yoshigai/go-programming-essence/chapter8/something"
)

func BenchmarkMakeSomething1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = something.MakeSomething1(1000)
	}
}

// Benchmark Result(BenchmarkMakeSomething1)
// ❯ go test -benchmem -bench MakeSomething1
// goos: darwin
// goarch: arm64
// pkg: github.com/k-yoshigai/go-programming-essence/chapter8/something
// BenchmarkMakeSomething1-8          14095             81212 ns/op           72390 B/op       1756 allocs/op
// PASS
// ok      github.com/k-yoshigai/go-programming-essence/chapter8/something 2.364s

func BenchmarkMakeSomething2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = something.MakeSomething2(1000)
	}
}

// Benchmark Result(BenchmarkMakeSomething1)
// ❯ go test -benchmem -bench MakeSomething2
// goos: darwin
// goarch: arm64
// pkg: github.com/k-yoshigai/go-programming-essence/chapter8/something
// BenchmarkMakeSomething2-8          15102             78295 ns/op           38347 B/op       1745 allocs/op
// PASS
// ok      github.com/k-yoshigai/go-programming-essence/chapter8/something 1.897s
