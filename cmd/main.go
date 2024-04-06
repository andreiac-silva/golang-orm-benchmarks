package main

import (
	"fmt"
	"testing"

	"github.com/andreiac-silva/golang-orm-benchmarks/benchmark"

	// Automatic load environment variables from .env.
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	bun := benchmark.NewBunBenchmark()
	bun.Init()

	result := testing.Benchmark(bun.FindPaginating)

	fmt.Println(result)
}
