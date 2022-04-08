package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req UserLoginRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Username, validation.Required, validation.Length(2, 20)),
		validation.Field(&req.Password, validation.Required, validation.Length(2, 100)),
	)
}
