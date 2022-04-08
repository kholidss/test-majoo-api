package request

type MerchantReportRequest struct {
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
