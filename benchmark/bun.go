package benchmark

import (
	"context"
	"database/sql"
	"testing"

	"github.com/andreiac-silva/golang-orm-benchmarks/benchmark/utils"
	"github.com/andreiac-silva/golang-orm-benchmarks/model"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type BunBenchmark struct {
	db  *bun.DB
	ctx context.Context
}

func NewBunBenchmark() utils.Benchmark {
	return &BunBenchmark{}
}

func (o *BunBenchmark) Init() error {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(utils.PostgresDSN)))
	o.db = bun.NewDB(sqldb, pgdialect.New())
	sqldb.SetMaxOpenConns(utils.PostgresMaxOpenConn)
	sqldb.SetMaxIdleConns(utils.PostgresMaxIdleConn)
	o.ctx = context.Background()
	return nil
}

func (o *BunBenchmark) Close() error {
	return o.db.Close()
}

func (o *BunBenchmark) Insert(b *testing.B) {
	utils.BeforeBenchmark()
	book := model.NewBook()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		book.ID = 0
		b.StartTimer()

		// Bun insert implementation.
		_, err := o.db.NewInsert().Model(book).Exec(o.ctx)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *BunBenchmark) InsertBulk(b *testing.B) {
	utils.BeforeBenchmark()
	books := model.NewBooks(200)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, book := range books {
			book.ID = 0
		}
		b.StartTimer()

		// Bun bulk insert implementation.
		_, err := o.db.NewInsert().Model(&books).Exec(o.ctx)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *BunBenchmark) Update(b *testing.B) {
	utils.BeforeBenchmark()
	book := model.NewBook()

	_, err := o.db.NewInsert().Model(book).Exec(o.ctx)
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Bun update implementation.
		_, err = o.db.NewUpdate().Model(book).WherePK().Exec(o.ctx)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *BunBenchmark) Delete(b *testing.B) {
	utils.BeforeBenchmark()
	book := model.NewBook()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		book.ID = 0
		_, err := o.db.NewInsert().Model(book).Exec(o.ctx)
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()

		// Bun delete implementation.
		_, err = o.db.NewDelete().Model(book).WherePK().Exec(o.ctx)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *BunBenchmark) FindOne(b *testing.B) {
	utils.BeforeBenchmark()

	n := b.N
	books := model.NewBooks(n)

	_, err := o.db.NewInsert().Model(&books).Exec(o.ctx)
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < n; i++ {
		b.StopTimer()
		bookID := books[i].ID
		book := new(model.Book)
		b.StartTimer()

		// Bun find by id implementation.
		err = o.db.NewSelect().Model(book).Where("id = ?", bookID).Scan(o.ctx)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *BunBenchmark) FindPaginating(b *testing.B) {
	utils.BeforeBenchmark()

	n := b.N
	limit := 10
	books := model.NewBooks(10 * n)
	_, err := o.db.NewInsert().Model(&books).Exec(o.ctx)
	if err != nil {
		b.Error(err)
	}

	var cursor int64
	var saved []*model.Book

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < n; i++ {
		b.StopTimer()
		cursor = model.GetMaxID(saved)
		saved = nil
		b.StartTimer()

		// Bun find with cursor pagination implementation.
		err = o.db.NewSelect().Model(&saved).Where("id > ?", cursor).Limit(limit).Scan(o.ctx)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}
