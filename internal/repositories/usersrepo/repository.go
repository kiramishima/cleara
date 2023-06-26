package usersrepo

import (
	"cleara/internal/core/domain"
	"cleara/internal/repositories"
	"context"
	"database/sql"
	"fmt"
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

// GetProfile Method for retrieve user profile by id
func (repo *UserRepository) GetProfile(id string) (*domain.User, error) {
	if v := strings.TrimSpace(id); v != "" {
		stmt, err := repo.db.PrepareContext(repo.context, "SELECT * FROM customer WHERE id = $1")
		if err != nil {
			return nil, fmt.Errorf("%s: %w", repositories.ErrPrepareStatement, err)
		}
		defer stmt.Close()

		var user = &domain.User{}

		row := stmt.QueryRowContext(repo.context, id)

		err = row.Scan(
			&user.ID,
			&user.Name,
		)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, repositories.ErrUserProfileNotFound
			} else {
				return nil, fmt.Errorf("%s: %w", repositories.ErrScanData, err)
			}
		}
		return user, nil
	}
	return nil, repositories.ErrUserProfileNotFound
}
