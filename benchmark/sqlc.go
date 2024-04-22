package benchmark

import (
	"context"
	"testing"

	"github.com/andreiac-silva/golang-orm-benchmarks/benchmark/sqlc/repository"
	"github.com/andreiac-silva/golang-orm-benchmarks/benchmark/utils"
	"github.com/andreiac-silva/golang-orm-benchmarks/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type SqlcBenchmark struct {
	repository *repository.Queries
	db         *pgx.Conn
	ctx        context.Context
}

func NewSqlcBenchmark() Benchmark {
	return &SqlcBenchmark{ctx: context.Background()}
}

func (s *SqlcBenchmark) Init() error {
	conn, err := pgx.Connect(context.Background(), utils.PostgresDSN)
	if err != nil {
		return err
	}
	s.db = conn
	s.repository = repository.New(conn)
	return nil
}

func (s *SqlcBenchmark) Close() error {
	return s.db.Close(s.ctx)
}

func (s *SqlcBenchmark) Insert(b *testing.B) {
	book := model.NewBook()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		book.ID = 0
		b.StartTimer()

		err := s.repository.Create(s.ctx, repository.CreateParams{
			Isbn:         book.ISBN,
			Title:        book.Title,
			Author:       book.Author,
			Genre:        book.Genre,
			Quantity:     int32(book.Quantity),
			PublicizedAt: pgtype.Timestamp{Time: book.PublicizedAt, Valid: true},
		})

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (s *SqlcBenchmark) InsertBulk(b *testing.B) {
	// Not supported, but there is an alternative. If the pgx is used, it is possible to use the CopyFrom function.
	b.SkipNow()
}

func (s *SqlcBenchmark) Update(b *testing.B) {
	book := model.NewBook()

	id, err := s.repository.CreateReturningID(s.ctx, repository.CreateReturningIDParams{
		Isbn:         book.ISBN,
		Title:        book.Title,
		Author:       book.Author,
		Genre:        book.Genre,
		Quantity:     int32(book.Quantity),
		PublicizedAt: pgtype.Timestamp{Time: book.PublicizedAt, Valid: true},
	})
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err = s.repository.Update(s.ctx, repository.UpdateParams{
			ID:           id,
			Isbn:         book.ISBN,
			Title:        book.Title,
			Author:       book.Author,
			Genre:        book.Genre,
			Quantity:     int32(book.Quantity),
			PublicizedAt: pgtype.Timestamp{Time: book.PublicizedAt, Valid: true},
		})

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (s *SqlcBenchmark) Delete(b *testing.B) {
	n := b.N
	book := model.NewBook()
	bookIDs := make([]int32, n)
	for i := 0; i < n; i++ {
		id, err := s.repository.CreateReturningID(s.ctx, repository.CreateReturningIDParams{
			Isbn:         book.ISBN,
			Title:        book.Title,
			Author:       book.Author,
			Genre:        book.Genre,
			Quantity:     int32(book.Quantity),
			PublicizedAt: pgtype.Timestamp{Time: book.PublicizedAt, Valid: true},
		})
		if err != nil {
			b.Error(err)
		}
		bookIDs[i] = id
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := s.repository.Delete(s.ctx, bookIDs[i])

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (s *SqlcBenchmark) FindByID(b *testing.B) {
	n := b.N
	book := model.NewBook()
	savedIDs := make([]int32, n)
	for i := 0; i < n; i++ {
		id, err := s.repository.CreateReturningID(s.ctx, repository.CreateReturningIDParams{
			Isbn:         book.ISBN,
			Title:        book.Title,
			Author:       book.Author,
			Genre:        book.Genre,
			Quantity:     int32(book.Quantity),
			PublicizedAt: pgtype.Timestamp{Time: book.PublicizedAt, Valid: true},
		})
		if err != nil {
			b.Error(err)
		}
		savedIDs[i] = id
	}

	b.ReportAllocs()
	b.ResetTimer()

	var bookID int32
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bookID = savedIDs[i]
		b.StartTimer()

		_, err := s.repository.Get(s.ctx, bookID)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (s *SqlcBenchmark) FindPaginating(b *testing.B) {
	n := b.N
	book := model.NewBook()
	bookIDs := make([]int32, n)
	for i := 0; i < n; i++ {
		id, err := s.repository.CreateReturningID(s.ctx, repository.CreateReturningIDParams{
			Isbn:         book.ISBN,
			Title:        book.Title,
			Author:       book.Author,
			Genre:        book.Genre,
			Quantity:     int32(book.Quantity),
			PublicizedAt: pgtype.Timestamp{Time: book.PublicizedAt, Valid: true},
		})
		if err != nil {
			b.Error(err)
		}
		bookIDs[i] = id
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := s.repository.ListPaginating(s.ctx, repository.ListPaginatingParams{
			ID:    int32(i),
			Limit: utils.PageSize,
		})

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}
