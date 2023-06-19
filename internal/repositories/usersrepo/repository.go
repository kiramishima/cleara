package usersrepo

import (
	"cleara/internal/core/domain"
	"context"
	"database/sql"
	"errors"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"strings"
)

type UserRepository struct {
	db      *sqlx.DB
	context context.Context
}

func NewRepository(conn *sqlx.DB, ctx context.Context) *UserRepository {
	return &UserRepository{
		db:      conn,
		context: ctx,
	}
}

func (repo *UserRepository) GetProfile(id string) (domain.User, error) {
	if v := strings.TrimSpace(id); v != "" {
		var user domain.User
		query := "SELECT * FROM users WHERE id = $1"
		err := repo.db.GetContext(repo.context, &user, query, id)
		if err != nil {
			if err != sql.ErrNoRows {
				return domain.User{}, err
			}
		}
		return user, nil
	}
	return domain.User{}, errors.New("user no found")
}
