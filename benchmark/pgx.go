package benchmark

import (
	"context"
	"testing"

	"github.com/andreiac-silva/golang-orm-benchmarks/benchmark/utils"
	"github.com/andreiac-silva/golang-orm-benchmarks/model"

	"github.com/jackc/pgx/v5"
)

var columns = []string{"isbn", "title", "author", "genre", "quantity", "publicized_at"}

type PgxBenchmark struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewPgxBenchmark() Benchmark {
	return &PgxBenchmark{
		ctx: context.Background(),
	}
}

func (p *PgxBenchmark) Init() error {
	var err error
	p.db, err = pgx.Connect(p.ctx, utils.PostgresDSN)
	return err
}

func (p *PgxBenchmark) Close() error {
	return p.db.Close(p.ctx)
}

func (p *PgxBenchmark) Insert(b *testing.B) {
	book := model.NewBook()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := p.db.Exec(p.ctx, utils.InsertQuery,
			book.ISBN, book.Title, book.Author, book.Genre, book.Quantity, book.PublicizedAt)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (p *PgxBenchmark) InsertBulk(b *testing.B) {
	var rows = make([][]interface{}, 0)
	for _, book := range model.NewBooks(utils.BulkInsertNumber) {
		rows = append(rows, []interface{}{book.ISBN, book.Title, book.Author, book.Genre, book.Quantity, book.PublicizedAt})
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := p.db.CopyFrom(p.ctx, pgx.Identifier{"books"}, columns, pgx.CopyFromRows(rows))

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (p *PgxBenchmark) Update(b *testing.B) {
	book := model.NewBook()
	var id int64
	err := p.db.QueryRow(p.ctx, utils.InsertReturningIDQuery,
		book.ISBN, book.Title, book.Author, book.Genre, book.Quantity, book.PublicizedAt).Scan(&id)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = p.db.Exec(p.ctx, utils.UpdateQuery,
			book.ISBN, book.Title, book.Author, book.Genre, book.Quantity, book.PublicizedAt, id)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (p *PgxBenchmark) Delete(b *testing.B) {
	book := model.NewBook()
	savedIDs := make([]int64, b.N)
	for i := 0; i < b.N; i++ {
		var id int64
		err := p.db.QueryRow(p.ctx, utils.InsertReturningIDQuery,
			book.ISBN, book.Title, book.Author, book.Genre, book.Quantity, book.PublicizedAt).Scan(&id)
		if err != nil {
			b.Error(err)
		}
		savedIDs[i] = id
	}

	b.ReportAllocs()
	b.ResetTimer()

	var bookID int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bookID = savedIDs[i]
		b.StartTimer()

		_, err := p.db.Exec(p.ctx, utils.DeleteQuery, bookID)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (p *PgxBenchmark) FindByID(b *testing.B) {
	book := model.NewBook()
	savedIDs := make([]int64, b.N)
	for i := 0; i < b.N; i++ {
		var id int64
		err := p.db.QueryRow(p.ctx, utils.InsertReturningIDQuery,
			book.ISBN, book.Title, book.Author, book.Genre, book.Quantity, book.PublicizedAt).Scan(&id)
		if err != nil {
			b.Error(err)
		}
		savedIDs[i] = id
	}

	b.ReportAllocs()
	b.ResetTimer()

	var bookID int64
	var foundBook model.Book
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bookID = savedIDs[i]
		foundBook = model.Book{}
		b.StartTimer()

		err := p.db.QueryRow(p.ctx, utils.SelectByIDQuery, bookID).Scan(
			&foundBook.ID,
			&foundBook.ISBN,
			&foundBook.Title,
			&foundBook.Author,
			&foundBook.Genre,
			&foundBook.Quantity,
			&foundBook.PublicizedAt,
		)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (p *PgxBenchmark) FindPage(b *testing.B) {
	n := b.N
	var rows = make([][]interface{}, 0)
	for _, book := range model.NewBooks(n) {
		rows = append(rows, []interface{}{book.ISBN, book.Title, book.Author, book.Genre, book.Quantity, book.PublicizedAt})
	}

	_, err := p.db.CopyFrom(p.ctx, pgx.Identifier{"books"}, columns, pgx.CopyFromRows(rows))
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < n; i++ {
		b.StopTimer()
		booksPage := make([]model.Book, utils.PageSize)
		b.StartTimer()

		result, err := p.db.Query(p.ctx, utils.SelectPaginatingQuery, i, utils.PageSize)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()

		for j := 0; result.Next() && j < utils.PageSize; j++ {
			err = result.Scan(
				&booksPage[j].ID,
				&booksPage[j].ISBN,
				&booksPage[j].Title,
				&booksPage[j].Author,
				&booksPage[j].Genre,
				&booksPage[j].Quantity,
				&booksPage[j].PublicizedAt,
			)
			b.StopTimer()
			if err != nil {
				b.Error(err)
			}
			b.StartTimer()
		}
	}
}
