package benchmark

import (
	"testing"

	"github.com/andreiac-silva/golang-orm-benchmarks/benchmark/utils"
	"github.com/andreiac-silva/golang-orm-benchmarks/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormBenchmark struct {
	db *gorm.DB
}

func NewGormBenchmark() utils.Benchmark {
	return &GormBenchmark{}
}

func (o *GormBenchmark) Init() error {
	// The config follows the performance section of the GORM documentation: https://gorm.io/docs/performance.html.
	pgConfig := postgres.New(postgres.Config{
		DSN:                  utils.PostgresDSN,
		PreferSimpleProtocol: true,
	})
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	}
	db, err := gorm.Open(pgConfig, gormConfig)
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	o.db = db
	sqlDB.SetMaxIdleConns(utils.PostgresMaxIdleConn)
	sqlDB.SetMaxOpenConns(utils.PostgresMaxOpenConn)
	return nil
}

func (o *GormBenchmark) Close() error {
	sqlDB, err := o.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (o *GormBenchmark) Insert(b *testing.B) {
	utils.BeforeBenchmark()
	book := model.NewBook()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		book.ID = 0
		b.StartTimer()

		// Gorm insert implementation.
		err := o.db.Create(book).Error

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *GormBenchmark) InsertBulk(b *testing.B) {
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

		// Gorm insert bulk implementation.
		err := o.db.Create(&books).Error

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *GormBenchmark) Update(b *testing.B) {
	utils.BeforeBenchmark()
	book := model.NewBook()

	err := o.db.Create(book).Error
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		book.Title = "Updated title"
		b.StartTimer()

		// Gorm update implementation.
		err = o.db.Save(book).Error

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *GormBenchmark) Delete(b *testing.B) {
	utils.BeforeBenchmark()

	n := b.N
	books := model.NewBooks(n)

	err := o.db.Create(books).Error
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bookID := books[i].ID
		b.StartTimer()

		// Gorm delete implementation.
		err = o.db.Delete(&model.Book{}, bookID).Error

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *GormBenchmark) FindByID(b *testing.B) {
	utils.BeforeBenchmark()

	n := b.N
	books := model.NewBooks(n)

	err := o.db.Create(books).Error
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		book := &model.Book{}
		bookID := books[i].ID
		b.StartTimer()

		// Gorm get by id implementation.
		err = o.db.First(book, bookID).Error

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}

func (o *GormBenchmark) FindPaginating(b *testing.B) {
	utils.BeforeBenchmark()

	n := b.N
	limit := 10
	books := model.NewBooks(limit * n)

	// Persist it in batches > https://github.com/bitmagnet-io/bitmagnet/issues/126.
	batches := model.Chunk(books, 10000)
	for _, chunk := range batches {
		err := o.db.Create(chunk).Error
		if err != nil {
			b.Error(err)
		}

	}

	b.ReportAllocs()
	b.ResetTimer()

	var cursor int64
	var saved []*model.Book

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		cursor = model.GetMaxID(saved)
		saved = nil
		b.StartTimer()

		// Gorm find paginating implementation.
		err := o.db.Limit(limit).Where("id > ?", cursor).Find(&saved).Error

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}
