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

func NewBunBenchmark() Benchmark {
	return &BunBenchmark{ctx: context.Background()}
}

func (o *BunBenchmark) Init() error {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(utils.PostgresDSN)))
	o.db = bun.NewDB(sqldb, pgdialect.New())
	return nil
}

func (o *BunBenchmark) Close() error {
	return o.db.Close()
}

func (o *BunBenchmark) Insert(b *testing.B) {
	book := model.NewBook()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		book.ID = 0
		b.StartTimer()

		_, err := o.db.NewInsert().Model(book).Exec(o.ctx)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *BunBenchmark) InsertBulk(b *testing.B) {
	books := model.NewBooks(utils.BulkInsertNumber)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for _, book := range books {
			book.ID = 0
		}
		b.StartTimer()

		_, err := o.db.NewInsert().Model(&books).Exec(o.ctx)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *BunBenchmark) Update(b *testing.B) {
	book := model.NewBook()

	_, err := o.db.NewInsert().Model(book).Exec(o.ctx)
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = o.db.NewUpdate().Model(book).WherePK().Exec(o.ctx)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *BunBenchmark) Delete(b *testing.B) {
	n := b.N
	books := model.NewBooks(n)

	_, err := o.db.NewInsert().Model(&books).Exec(o.ctx)
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	var book *model.Book
	for i := 0; i < n; i++ {
		b.StopTimer()
		book = new(model.Book)
		book.ID = books[i].ID
		b.StartTimer()

		_, err = o.db.NewDelete().Model(book).WherePK().Exec(o.ctx)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *BunBenchmark) FindByID(b *testing.B) {
	n := b.N
	books := model.NewBooks(n)
	_, err := o.db.NewInsert().Model(&books).Exec(o.ctx)
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	var book *model.Book
	var bookID int64
	for i := 0; i < n; i++ {
		b.StopTimer()
		book = new(model.Book)
		bookID = books[i].ID
		b.StartTimer()

		err = o.db.NewSelect().Model(book).Where("id = ?", bookID).Scan(o.ctx)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *BunBenchmark) FindPage(b *testing.B) {
	n := b.N
	books := model.NewBooks(n)
	_, err := o.db.NewInsert().Model(&books).Exec(o.ctx)
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	booksPage := make([]model.Book, utils.PageSize)
	for i := 0; i < n; i++ {
		b.StopTimer()
		booksPage = make([]model.Book, utils.PageSize)
		b.StartTimer()

		err = o.db.NewSelect().Model(&booksPage).Where("id > ?", i).Limit(utils.PageSize).Scan(o.ctx)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}
