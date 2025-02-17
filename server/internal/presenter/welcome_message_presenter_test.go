package presenter

import (
	"testing"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/require"
)

func TestWelcomeMessagePresenterRender(t *testing.T) {
	setup := func() (
		*WelcomeMessagePresenter,
		*usecase.MockUserRepo,
	) {
		uRepo := new(usecase.MockUserRepo)

		presenter := &WelcomeMessagePresenter{
			UserRepo: uRepo,
		}

		return presenter, uRepo
	}

	t.Run("happy path", func(t *testing.T) {
		presenter, uRepo := setup()

		template := "hello {{.UserDisplayName}} aka {{.UserHandleName}}"
		uRepo.On("GetByID", "user-id").Return(&model.User{
			Username:  "johny",
			FirstName: "John",
			LastName:  "Doe",
		}, nil)

		result, err := presenter.Render(template, "user-id")

		require.Nil(t, err)
		require.Equal(t, "hello John Doe aka @johny", result)
	})

	t.Run("with error response", func(t *testing.T) {
		presenter, uRepo := setup()

		uRepo.On("GetByID", "user-id").Return(nil, &model.AppError{Message: "err"})
		_, err := presenter.Render("any", "user-id")

		require.Equal(t, &model.AppError{Message: "err"}, err)
	})

	t.Run("with empty response", func(t *testing.T) {
		presenter, uRepo := setup()

		uRepo.On("GetByID", "user-id").Return(nil, nil)
		_, err := presenter.Render("any", "user-id")

		require.Equal(t, &model.AppError{Message: "User not found"}, err)
	})
}
