package handlers

import (
	"net/http"
	"strconv"

	"test-majoo-api/internal/transport/request"
	"test-majoo-api/internal/usecases"
	"test-majoo-api/internal/utils/token"

	"github.com/labstack/echo/v4"
)

type ReportHandler struct {
	report usecases.ReportInterface
}

func NewReporthandler(usecase usecases.ReportInterface) *ReportHandler {
	return &ReportHandler{
		report: usecase,
	}
}

func (h *ReportHandler) MerchantReport(c echo.Context) error {
	ctx := c.Request().Context()

	admin := c.Get("admin").(*token.Payload)

	pageQParam := c.QueryParam("page")
	limitQParam := c.QueryParam("limit")
	startDateQParam := c.QueryParam("start_date")
	endDateQParam := c.QueryParam("end_date")

	userPage, _ := strconv.Atoi(pageQParam)
	userLimit, _ := strconv.Atoi(limitQParam)

	reqParam := request.MerchantReportRequest{
		Page:      userPage,
		Limit:     userLimit,
		StartDate: startDateQParam,
		EndDate:   endDateQParam,
	}

	result, _ := h.report.MerchantReport(reqParam, ctx, admin.Id)

	return c.JSON(http.StatusOK, result)
}
