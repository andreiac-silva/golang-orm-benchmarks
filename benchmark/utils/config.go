package utils

import (
	"os"
)

const (
	BulkInsertNumber = 2000
	BatchSize        = 10000
	PageSize         = 10
)

var PostgresDSN string

func init() {
	PostgresDSN = os.Getenv("POSTGRES_DSN")
	if PostgresDSN == "" {
		panic("POSTGRES_DSN is required")
	}
}
