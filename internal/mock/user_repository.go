package mock

import (
	"cleara/internal/core/domain"
	err "cleara/internal/repositories"
	"strconv"
)

type mockUserRepository struct {
	users []*domain.User
}

func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		users: []*domain.User{
			{
				ID:   1,
				Name: "Nami",
			},
			{
				ID:   2,
				Name: "Robin",
			},
		},
	}
}

func (r *mockUserRepository) GetProfile(id string) (*domain.User, error) {
	for _, b := range r.users {
		ID, _ := strconv.Atoi(id)
		if b.ID == int16(ID) {
			return b, nil
		}
	}

	return nil, err.ErrUserProfileNotFound
}
