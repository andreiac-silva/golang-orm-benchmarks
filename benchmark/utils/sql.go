package utils

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	queries "github.com/andreiac-silva/golang-orm-benchmarks/sql"

	// Postgres driver.
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	PostgresDSN         string
	PostgresMaxOpenConn int
	PostgresMaxIdleConn int
)

func init() {
	var err error

	PostgresDSN = os.Getenv("POSTGRES_DSN")
	if PostgresDSN == "" {
		panic("POSTGRES_DSN is required")
	}

	maxOpenConnStr := os.Getenv("POSTGRES_MAX_OPEN_CONN")
	PostgresMaxOpenConn, err = strconv.Atoi(maxOpenConnStr)
	if err != nil {
		panic("POSTGRES_MAX_OPEN_CONN is required")
	}

	maxIdleConnStr := os.Getenv("POSTGRES_MAX_IDLE_CONN")
	PostgresMaxIdleConn, err = strconv.Atoi(maxIdleConnStr)
	if err != nil {
		panic("POSTGRES_MAX_IDLE_CONN is required")
	}
}

func recreateDatabase() {
	db, err := sql.Open("pgx", PostgresDSN)
	if err != nil {
		log.Fatal("the benchmark execution was aborted", err)
	}

	defer func() {
		_ = db.Close()
	}()

	_, err = db.Exec(queries.RecreateDatabaseSQL)
	if err != nil {
		log.Fatal("the benchmark execution was aborted", err)
	}
}
