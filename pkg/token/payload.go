package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var TokenOutDate = errors.New("Token is out of date")

type Payload struct {
	Id        uuid.UUID `json:"id"`
	UserId    int64     `json:"user_id"`
	RoleID    int32     `json:"role_id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(user_id int64, role_id int32, username string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		Id:        tokenId,
		UserId:    user_id,
		RoleID:    role_id,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return TokenOutDate
	}
	return nil
}
