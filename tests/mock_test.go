package tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	// Buat mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	// Buat GORM DB dengan mock connection
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("Failed to open GORM DB: %v", err)
	}

	cleanup := func() {
		db.Close()
	}

	return gormDB, mock, cleanup
}
