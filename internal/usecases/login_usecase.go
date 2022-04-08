package usecases

import (
	"context"
	"os"
	"time"

	"test-majoo-api/internal/entities"
	"test-majoo-api/internal/repositories"
	"test-majoo-api/internal/transport/request"
	"test-majoo-api/internal/utils"

	"github.com/dgrijalva/jwt-go"
)

type LoginInterface interface {
	UserLogin(ctx context.Context, params *request.UserLoginRequest) (*entities.User, string, error)
}

type loginUsecase struct {
	LoginRepository repositories.LoginRepositoryInterface
}

func NewLoginUsecase(loginRepository repositories.LoginRepositoryInterface) LoginInterface {
	return &loginUsecase{
		LoginRepository: loginRepository,
	}
}

func (u loginUsecase) UserLogin(ctx context.Context, params *request.UserLoginRequest) (*entities.User, string, error) {
	result, _ := u.LoginRepository.UserLogin(ctx, params)
	if result == nil {
		return nil, "", utils.NotFound("WRONG_USERNAME_OR_PASSWORD")
	}

	if result.Password != params.Password {
		return nil, "", utils.NotFound("WRONG_USERNAME_OR_PASSWORD")
	}

	var mySigningKey = []byte(os.Getenv("JWT_SECRET"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = result.Id
	claims["user_name"] = result.UserName
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return nil, "", utils.InternalServerError("COULD_NOT_LOGIN")
	}

	return result, tokenString, nil
}
