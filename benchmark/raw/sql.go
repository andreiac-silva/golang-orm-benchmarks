package benchmark

import _ "embed"

var (
	//go:embed sql/insert.sql
	insertQuery string
	//go:embed sql/insert_returning_id.sql
	insertReturningIDQuery string
	//go:embed sql/insert_bulk.sql
	insertBulkQuery string
	//go:embed sql/update.sql
	updateQuery string
	//go:embed sql/delete.sql
	deleteQuery string
	//go:embed sql/select_by_id.sql
	selectByIDQuery string
	//go:embed sql/select_paginating.sql
	selectPaginatingQuery string
)
