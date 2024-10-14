package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
}

// CreateToken implements Maker.
func (j JWTMaker) CreateToken(user_id int64, role_id int32, username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(user_id, role_id, username, duration)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(j.secretKey))
}

// VerifyToken implements Maker.
func (j JWTMaker) VerifyToken(token string) (*Payload, error) {
	claims := &Payload{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpected
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, ErrOutDate
		case errors.Is(err, jwt.ErrSignatureInvalid):
			return nil, ErrInvalid
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, ErrUnSupported
		default:
			return nil, err
		}
	}
	return claims, nil
}

func NewJWTMaker(secretKey string) (Maker, error) {
	return JWTMaker{secretKey}, nil
}
