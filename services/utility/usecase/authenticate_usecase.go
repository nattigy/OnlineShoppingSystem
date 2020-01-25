package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/utility"
)

type AuthenticateUsecase struct {
	authRepo utility.AuthenticationRepository
}

func NewAuthenticateUsecase(authRepo utility.AuthenticationRepository) *AuthenticateUsecase {
	return &AuthenticateUsecase{authRepo: authRepo}
}

func (a *AuthenticateUsecase) Authenticate(u models.User) (bool, models.User, error) {
	auth, user, err := a.authRepo.Authenticate(u)
	return auth, user, err
}
