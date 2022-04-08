package token

import (
	"errors"
	"time"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("TOKEN_IS_INVALID")
	ErrExpiredToken = errors.New("TOKEN_HAS_EXPIRED")
)

// Payload contains the payload data of the token
type Payload struct {
	Exp  int    `json:"exp"`
	Iss  string `json:"iss"`
	Name string `json:"name"`
	Nbf  int    `json:"nbf"`
	Sub  int    `json:"sub"`
	Id   int    `json:"id"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(email string, iss string, name string, nbf int, role string, sub int, duration time.Duration) (*Payload, error) {

	payload := &Payload{
		Iss: iss,
		Nbf: nbf,
		Sub: sub,
		Exp: int(time.Now().Add(duration).Unix()),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	tm := time.Unix(int64(payload.Exp), 0)
	if time.Now().After(tm) {
		return ErrExpiredToken
	}
	return nil
}
