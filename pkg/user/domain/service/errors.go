package usersrv

import "github.com/pkg/errors"

var (
	ErrUserNotFound        = errors.New("Пользователь не найден")
	ErrUserAlreadyExists   = errors.New("Пользователь уже существует")
	ErrUserProductNotFound = errors.New("Товары пользователя не найдены")
)
