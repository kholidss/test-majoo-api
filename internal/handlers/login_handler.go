package handlers

import (
	"net/http"

	"test-majoo-api/internal/transport/request"
	"test-majoo-api/internal/transport/response"
	"test-majoo-api/internal/usecases"
	"test-majoo-api/internal/utils"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type LoginHandler struct {
	login usecases.LoginInterface
}

func NewLoginHandler(usecase usecases.LoginInterface) *LoginHandler {
	return &LoginHandler{
		login: usecase,
	}
}

func (h *LoginHandler) UserLogin(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.UserLoginRequest
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, utils.BadRequest(`BAD_REQUEST_PARAM`))
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, utils.InvalidInput(errVal))
	}

	result, token, err := h.login.UserLogin(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.UnprocessableEntity(err.Error()))
	}

	response := &response.LoginResponse{
		ID:          result.Id,
		Name:        result.Name,
		Username:    result.UserName,
		CreatedBy:   result.CreatedBy,
		UpdatedBy:   result.UpdatedBy,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
		AccessToken: token,
	}

	return c.JSON(http.StatusOK, response)
	// return c.JSON(http.StatusOK, role)
}
