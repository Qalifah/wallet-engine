package errors

import "errors"

var (
	ErrInvalidWalletID    = errors.New("invalid wallet id")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidUserID      = errors.New("invalid user id")
	ErrUnExpectedError    = errors.New("n error occurred, please try again")
	ErrWalletNotFound     = errors.New("wallet does not exist")
)
