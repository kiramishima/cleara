package user_service

import (
	"cleara/internal/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProfile(t *testing.T) {
	repo := mock.NewMockUserRepository()
	uc := New(repo)

	t.Run("OK", func(t *testing.T) {
		b, err := uc.GetProfile("1")
		assert.NoError(t, err)
		assert.Equal(t, "Nami", b.Name)
	})
	t.Run("Not Found", func(t *testing.T) {
		_, err := uc.GetProfile("0")
		assert.Error(t, err)
	})
}
