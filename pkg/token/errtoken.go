package token

import "errors"

var ErrUnexpected = errors.New("unexpected signing method")
var ErrUnSupported = errors.New("unsupported type")
var ErrInvalid = errors.New("invalid token")
var ErrParse = errors.New("could not parse token")
var ErrOutDate = errors.New("token out of date")
var ErrLenKey = errors.New("secret key must be at least 32")
