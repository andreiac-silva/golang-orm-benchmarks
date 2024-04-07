package utils

import (
	"os"
	"strconv"
)

const (
	InsertNumberItems = 2000
	BatchSize         = 10000
	PageSize          = 10
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
