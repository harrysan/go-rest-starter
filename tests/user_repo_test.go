package tests

import (
	"fmt"
	"testing"
	"time"

	"finance-tracker/internal/mods/auth/dal"
	"finance-tracker/internal/mods/auth/schema"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite" // Import driver modernc
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, mock, cleanup := SetupMockDB(t)
	defer cleanup()
	repo := dal.NewUserDal(db)

	t.Run("success create user", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "users"`).
			WithArgs(
				"johndoe",
				"John Doe",
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				"johndoe@example.com",
				sqlmock.AnyArg(),
				"activated",
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		user := &schema.User{
			Username:  "johndoe",
			Name:      "John Doe",
			Password:  "secret",
			Phone:     "62889900",
			Email:     "johndoe@example.com",
			Remark:    "test user john",
			Status:    "activated",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := repo.Create(user)
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			fmt.Println("Mock expectations error:", err)
		}
	})
}
