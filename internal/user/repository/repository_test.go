package repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ryanadiputraa/api-udrio/internal/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestSaveOrUpdate(t *testing.T) {
	db, sqlDB, mock := test.NewMockDB(t)
	defer sqlDB.Close()
	r := NewUserRepository(db)

	cases := []struct {
		description       string
		mockRepoBehaviour func(mock sqlmock.Sqlmock)
	}{
		{
			description: "should insert new user data",
			mockRepoBehaviour: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("^INSERT INTO \"users\" *").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				currTime := time.Now()
				user := domain.User{
					ID:        "1",
					FirstName: "John",
					LastName:  "Doe",
					Email:     "john@mail.com",
					Picture:   "https://domain.com/image.png",
					Locale:    "id",
					CreatedAt: currTime,
					UpdatedAt: currTime,
				}

				err := r.SaveOrUpdate(context.TODO(), user)
				assert.Nil(t, err)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			c.mockRepoBehaviour(mock)
		})
	}
}

func TestFindByID(t *testing.T) {
	db, sqlDB, mock := test.NewMockDB(t)
	defer sqlDB.Close()

	cases := []struct {
		description       string
		mockRepoBehaviour func(mock sqlmock.Sqlmock, r domain.UserRepository)
	}{
		{
			description: "should return user data with id = '1'",
			mockRepoBehaviour: func(mock sqlmock.Sqlmock, r domain.UserRepository) {
				currTime := time.Now()
				rows := sqlmock.NewRows([]string{
					"id",
					"first_name",
					"last_name",
					"email",
					"picture",
					"locale",
					"created_at",
					"updated_at",
				}).AddRow("1", "John", "Doe", "john@mail.com", "https://domain.com/image.png", "id", currTime, currTime)
				mock.ExpectQuery("^SELECT (.+) FROM \"users\" *").WillReturnRows(rows)

				user, err := r.FindByID(context.TODO(), "1")
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, domain.User{
					ID:        "1",
					FirstName: "John",
					LastName:  "Doe",
					Email:     "john@mail.com",
					Picture:   "https://domain.com/image.png",
					Locale:    "id",
					CreatedAt: currTime,
					UpdatedAt: currTime,
				}, user)
			},
		},
		{
			description: "should return error for no record user data",
			mockRepoBehaviour: func(mock sqlmock.Sqlmock, r domain.UserRepository) {
				rows := sqlmock.NewRows([]string{
					"id",
					"first_name",
					"last_name",
					"email",
					"picture",
					"locale",
					"created_at",
					"updated_at",
				})
				mock.ExpectQuery("^SELECT (.+) FROM \"users\" *").WillReturnRows(rows)

				user, err := r.FindByID(context.TODO(), "2")
				assert.Error(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, "", user.ID)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			r := NewUserRepository(db)
			c.mockRepoBehaviour(mock, r)
		})
	}
}
