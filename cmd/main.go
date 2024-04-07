package main

import (
	"fmt"
	"testing"

	"github.com/andreiac-silva/golang-orm-benchmarks/benchmark"
	raw "github.com/andreiac-silva/golang-orm-benchmarks/benchmark/raw"

	// Auto load .env file.
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := raw.NewRawBenchmark()
	r.Init()
	result := testing.Benchmark(r.FindPaginating)
	fmt.Println(result)

	b := benchmark.NewBunBenchmark()
	b.Init()
	result1 := testing.Benchmark(b.FindPaginating)
	fmt.Println(result1)

	g := benchmark.NewGormBenchmark()
	g.Init()
	result2 := testing.Benchmark(g.FindPaginating)
	fmt.Println(result2)
}
