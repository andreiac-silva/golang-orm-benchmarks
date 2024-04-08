package main

import (
	"fmt"
	"testing"

	"github.com/andreiac-silva/golang-orm-benchmarks/benchmark"
	// Auto load .env file.
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := benchmark.NewRawBenchmark()
	r.Init()
	result := testing.Benchmark(r.FindPaginating)
	fmt.Println(result)

	pgx := benchmark.NewPgxBenchmark()
	pgx.Init()
	result3 := testing.Benchmark(pgx.FindPaginating)
	fmt.Println(result3)

	b := benchmark.NewBunBenchmark()
	b.Init()
	result1 := testing.Benchmark(b.FindPaginating)
	fmt.Println(result1)

	g := benchmark.NewGormBenchmark()
	g.Init()
	result2 := testing.Benchmark(g.FindPaginating)
	fmt.Println(result2)
}
