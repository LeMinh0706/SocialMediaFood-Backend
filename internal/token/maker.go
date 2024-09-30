package token

import (
	"time"
)

type Maker interface {
	///Create a new token with username, roleid, duration
	CreateToken(username string, duration time.Duration) (string, error)

	//Is token valid?
	VerifyToken(token string) (*Payload, error)
}
