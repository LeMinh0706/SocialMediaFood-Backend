package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var TokenOutDate = errors.New("Token is out of date")

type Payload struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	RoleID    int32     `json:"role_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, role_id int32, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		Id:        tokenId,
		Username:  username,
		RoleID:    role_id,
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
