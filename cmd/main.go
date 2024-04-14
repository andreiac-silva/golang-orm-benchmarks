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
	result := testing.Benchmark(r.FindByID)
	fmt.Println(fmt.Sprintf("Raw: %v", result))

	pgx := benchmark.NewPgxBenchmark()
	pgx.Init()
	result1 := testing.Benchmark(pgx.FindByID)
	fmt.Println(fmt.Sprintf("Pgx: %v", result1))

	b := benchmark.NewBunBenchmark()
	b.Init()
	result2 := testing.Benchmark(b.FindPaginating)
	fmt.Println(fmt.Sprintf("Bun: %v", result2))

	g := benchmark.NewGormBenchmark()
	g.Init()
	result3 := testing.Benchmark(g.FindPaginating)
	fmt.Println(fmt.Sprintf("Gorm: %v", result3))

	e := benchmark.NewEntBenchmark()
	e.Init()
	result4 := testing.Benchmark(e.FindPaginating)
	fmt.Println(fmt.Sprintf("Ent: %v", result4))

	s := benchmark.NewSqlcBenchmark()
	s.Init()
	result5 := testing.Benchmark(s.FindPaginating)
	fmt.Println(fmt.Sprintf("Sqlc: %v", result5))
}
