package response

import "time"

type LoginResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"email"`
	CreatedBy   int       `json:"created_by"`
	UpdatedBy   int       `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccessToken string    `json:"access_token"`
}
