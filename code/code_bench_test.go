package code

import (
	"testing"
)

// Run with:
// $ go test -bench=BenchmarkV1 --benchtime=30s
func BenchmarkV1(b *testing.B) {
	dropAndCreateTables()

	for i := 0; i < b.N; i++ {
		v1()
	}
}

// Run with:
// $ go test -bench=BenchmarkV2 --benchtime=30s
func BenchmarkV2(b *testing.B) {
	dropAndCreateTables()

	for i := 0; i < b.N; i++ {
		v2()
	}
}

// Run with:
// $ go test -bench=BenchmarkV3 --benchtime=30s
func BenchmarkV3(b *testing.B) {
	dropAndCreateTables()

	for i := 0; i < b.N; i++ {
		v3()
	}
}
