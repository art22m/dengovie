package usecase

import "github.com/pkg/errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrChatAlreadyExists = errors.New("chat already exists")

	ErrUserNotFound = errors.New("user not found")
	ErrChatNotFound = errors.New("chat not found")
	ErrDebtNotFound = errors.New("debt not found")
)
