package productsrv

import "github.com/pkg/errors"

var (
	ErrProductNotFound = errors.New("Товар не найден")
)
