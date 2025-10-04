package tokenprovider

import (
	"crud-go/common"
	"errors"
	"time"
)

type TokenType int

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewCustomErrorResponse(
		errors.New("token not found"),
		"token not found",
		"ERROR_NOT_FOUND")

	ErrEncodingToken = common.NewCustomErrorResponse(
		errors.New("error encoding the token"),
		"error encoding the token",
		"ERROR_ENCODING_TOKEN")

	ErrInvalidToken = common.NewCustomErrorResponse(
		errors.New("invalid token provided"),
		"invalid token provided",
		"ERROR_INVALID_TOKEN",
	)
)

type Token struct {
	Token  string    `json:"Token"`
	Create time.Time `json:"created"`
	Expiry int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	Expiry int
}
