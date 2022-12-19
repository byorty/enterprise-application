package db

import (
	"github.com/pkg/errors"
)

var (
	ErrRowNotFound  = errors.New("api.row_not_found_error")
	ErrRowsNotFound = errors.New("api.rows_not_found_error")
	ErrCantSave     = errors.New("api.saving_error")
	ErrCantDelete   = errors.New("api.deleting_error")
	ErrInvalidKind  = errors.New("api.no_db_by_kind_error")
)
