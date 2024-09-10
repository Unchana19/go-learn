package exceptions

import "errors"

var (
	ErrDuplicatedEmail = errors.New("duplicated email")
	ErrLoginFailed = errors.New("login failed")
)
