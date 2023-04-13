package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/utils"
)

type loginService struct {
	UsersRepo     repos.UsersRepo
	UsersService  UsersService
	TokensService TokensService
	SigningMethod jwt.SigningMethod
	HashCost      int
}

func NewLoginService(
	usersRepo repos.UsersRepo,
	usersService UsersService,
	tokensService TokensService,
	signingMethod jwt.SigningMethod,
	hashCost int,
) *loginService {
	return &loginService{
		UsersRepo:     usersRepo,
		UsersService:  usersService,
		TokensService: tokensService,
		SigningMethod: signingMethod,
		HashCost:      hashCost,
	}
}

func (ss *loginService) Login(username string, password string) (*models.LoginResponse, bool, error) {
	var item *models.LoginResponse

	user, err := ss.UsersService.GetOneByUsername(username)
	if err != nil {
		return item, false, err
	}

	isMatch := utils.CheckPasswordHash(password, user.Password)
	if !isMatch {
		return item, false, nil
	}

	token := jwt.NewWithClaims(ss.SigningMethod, models.CustomJWTClaims{
		CustomClaims: models.CustomClaims{
			UserID: user.ID,
		},
	})

	signedToken, err := ss.TokensService.SignTokenWithCustomClaims(token)
	if err != nil {
		return item, false, err
	}

	item = &models.LoginResponse{
		UserGetAllResponse: models.UserGetAllResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			LastLogin: user.LastLogin,
			IsAdmin:   user.IsAdmin,
			UserType:  user.UserType,
			UpdatedAt: user.UpdatedAt,
			CreatedAt: user.CreatedAt,
		},

		Token: signedToken,
	}

	return item, true, nil
}
