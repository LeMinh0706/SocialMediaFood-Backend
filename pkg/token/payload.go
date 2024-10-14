package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var TokenOutDate = errors.New("Token is out of date")

type Payload struct {
	Id       uuid.UUID `json:"id"`
	UserId   int64     `json:"user_id"`
	RoleID   int32     `json:"role_id"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

func NewPayload(user_id int64, role_id int32, username string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		Id:       tokenId,
		UserId:   user_id,
		RoleID:   role_id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
	return payload, nil
}
