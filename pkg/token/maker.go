package token

import (
	"time"

	"github.com/google/uuid"
)

type Maker interface {
	///Create a new token with username, roleid, duration
	CreateToken(tokenId uuid.UUID, user_id int64, username string, duration time.Duration) (string, error)

	//Is token valid?
	VerifyToken(token string) (*Payload, error)
}
