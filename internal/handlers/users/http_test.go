package users

import (
	"cleara/internal/core/domain"
	"cleara/internal/mock"
	repoErr "cleara/internal/repositories"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProfile(t *testing.T) {
	testCases := map[string]struct {
		ID            any
		buildStubs    func(uc *mock.MockUserService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		"OK": {
			ID: "1",
			buildStubs: func(uc *mock.MockUserService) {
				uc.EXPECT().
					GetProfile(gomock.Eq(1)).
					Times(1).
					Return(&domain.User{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		"Invalid URL Param": {
			ID: "ID",
			buildStubs: func(uc *mock.MockUserService) {
				uc.EXPECT().
					GetProfile(gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		"Not Found": {
			ID: 1,
			buildStubs: func(uc *mock.MockUserService) {
				uc.EXPECT().
					GetProfile(gomock.Any()).
					Times(1).
					Return(nil, repoErr.ErrUserProfileNotFound)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		"Unexpected Error": {
			ID: 1,
			buildStubs: func(uc *mock.MockUserService) {
				uc.EXPECT().
					GetProfile(gomock.Any()).
					Times(1).
					Return(nil, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := mock.NewMockUserService(ctrl)
			tc.buildStubs(uc)

			recorder := httptest.NewRecorder()

			url := fmt.Sprint("/v1/users/", tc.ID)
			println("url", url)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			assert.NoError(t, err)

			app := gin.New()
			v1 := app.Group("/v1")
			userRoutes := v1.Group("/users")

			var userHandlers = NewUserHandlers(uc)
			userRoutes.GET("/:id", userHandlers.GetProfile)
			app.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
