package usersrepo

import (
	"cleara/internal/core/domain"
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"strings"
)

type UserRepository struct {
	db      *sql.DB
	context context.Context
}

func NewRepository(conn *sql.DB, ctx context.Context) *UserRepository {
	return &UserRepository{
		db:      conn,
		context: ctx,
	}
}

func (repo *UserRepository) GetProfile(id string) (domain.User, error) {
	if v := strings.TrimSpace(id); v != "" {
		var user domain.User
		query := "SELECT * FROM customer WHERE id = $1"
		err := repo.db.QueryRowContext(repo.context, query, id).Scan(
			&user.ID,
			&user.Name,
		)
		if err != nil {
			if err != sql.ErrNoRows {
				return domain.User{}, err
			}
		}
		return user, nil
	}
	return domain.User{}, errors.New("user no found")
}
