package main

import (
	"fmt"
	"testing"

	"github.com/andreiac-silva/golang-orm-benchmarks/benchmark"

	// Automatic load environment variables from .env.
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	gorm := benchmark.NewGormBenchmark()
	gorm.Init()

	result := testing.Benchmark(gorm.FindPaginating)

	fmt.Println(result)
}
