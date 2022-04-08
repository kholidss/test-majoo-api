package repositories

import (
	"context"

	"test-majoo-api/internal/entities"
	"test-majoo-api/internal/transport/request"

	"github.com/jinzhu/gorm"
)

type LoginRepositoryInterface interface {
	UserLogin(ctx context.Context, params *request.UserLoginRequest) (*entities.User, error)
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) LoginRepositoryInterface {
	return &loginRepository{
		db: db,
	}
}

func (r *loginRepository) UserLogin(ctx context.Context, params *request.UserLoginRequest) (*entities.User, error) {

	user := &entities.User{
		UserName: params.Username,
	}

	result := r.db.Table("Users").Where("user_name = ?", params.Username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
