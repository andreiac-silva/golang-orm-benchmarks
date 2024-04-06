package utils

import "testing"

// Benchmark interface was inspired by https://github.com/efectn/go-orm-benchmarks/blob/master/helper/suite.go.
type Benchmark interface {
	Init() error
	Close() error
	Insert(b *testing.B)
	InsertBulk(b *testing.B)
	Update(b *testing.B)
	Delete(b *testing.B)
	FindOne(b *testing.B)
	FindPaginating(b *testing.B)
}

func BeforeBenchmark() {
	recreateDatabase()
}
