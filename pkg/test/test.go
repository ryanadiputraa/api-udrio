package test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockDB(t *testing.T) (db *gorm.DB, sqlDB *sql.DB, mock sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("fail to create mock db conn: ", err.Error())
	}

	db, err = gorm.Open(postgres.New(
		postgres.Config{
			Conn:       sqlDB,
			DriverName: "postgres",
		},
	), &gorm.Config{})
	if err != nil {
		t.Fatal("fail to open db conn: ", err.Error())
	}
	return
}
