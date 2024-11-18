package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	Id       uuid.UUID `json:"id"`
	UserId   int64     `json:"user_id"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

func NewPayload(tokenId uuid.UUID, user_id int64, username string, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		Id:       tokenId,
		UserId:   user_id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt.Time) {
		return ErrOutDate
	}
	return nil
}
