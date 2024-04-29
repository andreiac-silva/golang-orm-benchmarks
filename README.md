## Golang ORM Benchmarks

#### This repository contains benchmarks for the following projects:
- [GORM](https://gorm.io/)
- [Ent](https://entgo.io/)
- [Bun](https://bun.uptrace.dev/)
- [Sqlc](https://sqlc.dev/)

#### And also, pure SQL benchmarks using:
- [pgx](https://github.com/jackc/pgx)
- [database/sql](https://pkg.go.dev/database/sql)

<p>To execute all benchmarks operations, run the following command:

```bash
$ make benchmark-all
```

<p>Keep it in mind that running all the benchmarks at once can take some time to complete. 
<p>If you want to run a specific benchmark, you can use the following commands:

```bash
$ make benchmark-insert
$ make benchmark-insert-bulk
$ make benchmark-update
$ make benchmark-delete
$ make benchmark-select-one
$ make benchmark-select-page
```

Modeling credits: [go-orm-benchmarks](https://github.com/efectn/go-orm-benchmarks).