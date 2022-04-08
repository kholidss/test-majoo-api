package factories

import (
	"test-majoo-api/internal/handlers"
	"test-majoo-api/internal/repositories"
	"test-majoo-api/internal/usecases"

	"github.com/jinzhu/gorm"
)

func CreateReportHandler(db *gorm.DB) *handlers.ReportHandler {
	reportRepo := repositories.NewReportRepository(db)
	reportUsecase := usecases.NewReportusecase(reportRepo)
	return handlers.NewReporthandler(reportUsecase)
}
