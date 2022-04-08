package factories

import (
	"test-majoo-api/internal/handlers"
	"test-majoo-api/internal/repositories"
	"test-majoo-api/internal/usecases"

	"github.com/jinzhu/gorm"
)

func CreateLoginHandler(db *gorm.DB) *handlers.LoginHandler {
	loginRepo := repositories.NewLoginRepository(db)
	loginUsecase := usecases.NewLoginUsecase(loginRepo)
	return handlers.NewLoginHandler(loginUsecase)
}
