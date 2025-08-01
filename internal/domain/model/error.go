package model

import "errors"

var (
	ErrAPIFetch                = errors.New("API Error")
	ErrGenerationsLimitReached = errors.New("generations monthly limit reached")
	ErrUserNotFound            = errors.New("user not found")
)
