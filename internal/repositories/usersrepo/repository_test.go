package usersrepo

import (
	"cleara/internal/core/domain"
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGetProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	ctx := context.Background()

	repo := NewRepository(db, ctx)

	user := &domain.User{
		ID:   1,
		Name: "Uta",
	}

	t.Run("OK", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(user.ID, user.Name)

		mock.ExpectPrepare("SELECT * FROM customer WHERE id = ").
			ExpectQuery().
			WithArgs(user.ID).
			WillReturnRows(rows)

		id := strconv.Itoa(int(user.ID))
		userProfile, err := repo.GetProfile(id)
		assert.NoError(t, err)
		assert.Equal(t, user, userProfile)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Query Failed", func(t *testing.T) {
		mock.ExpectPrepare("SELECT * FROM customer WHERE id = ").
			ExpectQuery().
			WithArgs(user.ID).
			WillReturnError(sql.ErrConnDone)

		id := strconv.Itoa(int(user.ID))
		userProfile, err := repo.GetProfile(id)
		assert.Error(t, err)
		assert.Empty(t, userProfile)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Prepare Failed", func(t *testing.T) {
		mock.ExpectPrepare("SELECT \\* FROM customer WHERE id = ").
			WillReturnError(sql.ErrConnDone)

		id := strconv.Itoa(int(user.ID))
		userProfile, err := repo.GetProfile(id)
		assert.Error(t, err)
		assert.Empty(t, userProfile)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Not Found", func(t *testing.T) {
		mock.ExpectPrepare("SELECT \\* FROM customer WHERE id = ").
			ExpectQuery().
			WithArgs(user.ID).
			WillReturnError(sql.ErrNoRows)

		id := strconv.Itoa(int(user.ID))
		userProfile, err := repo.GetProfile(id)
		assert.Error(t, err)
		assert.Empty(t, userProfile)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
