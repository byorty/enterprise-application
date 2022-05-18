package ordersrv

import "github.com/pkg/errors"

var (
	ErrOrderNotFound       = errors.New("Заказ не найден")
	ErrOrderAmountNotEqual = errors.New("Неверная сумма заказа")
)
