package usecases

import (
	"context"

	"test-majoo-api/internal/entities"
	"test-majoo-api/internal/repositories"
	"test-majoo-api/internal/transport/request"
)

type ReportInterface interface {
	MerchantReport(request request.MerchantReportRequest, ctx context.Context, userId int) (*entities.Transactions, error)
}

type reportUsecase struct {
	ReportRepository repositories.ReportRepositoryInterface
}

func NewReportusecase(reportRepository repositories.ReportRepositoryInterface) ReportInterface {
	return &reportUsecase{
		ReportRepository: reportRepository,
	}
}

func (u reportUsecase) MerchantReport(request request.MerchantReportRequest, ctx context.Context, userId int) (*entities.Transactions, error) {
	result, _ := u.ReportRepository.MerchantReport(request, ctx, userId)

	return result, nil
}
