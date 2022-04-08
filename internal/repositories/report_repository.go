package repositories

import (
	"context"

	"test-majoo-api/internal/entities"
	"test-majoo-api/internal/transport/request"

	"github.com/jinzhu/gorm"
)

type ReportRepositoryInterface interface {
	MerchantReport(request request.MerchantReportRequest, ctx context.Context, userId int) (*entities.Transactions, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepositoryInterface {
	return &reportRepository{
		db: db,
	}
}

func (r *reportRepository) MerchantReport(request request.MerchantReportRequest, ctx context.Context, userId int) (*entities.Transactions, error) {
	var transaction entities.Transactions

	limit := request.Limit
	offset := request.Limit * (request.Page - 1)

	result := r.db.Raw("SELECT * FROM Transactions t JOIN Merchants m on m.id = t.merchant_id join Outlets o on o.id = t.outlet_id join Users u on u.id = t.created_by WHERE t.updated_at BETWEEN ? AND ? AND t.merchant_id = ? LIMIT ? OFFSET ?", request.StartDate, request.EndDate, userId, limit, offset).Scan(&transaction)

	if result.Error != nil {
		return nil, result.Error
	}

	return &transaction, nil
}
