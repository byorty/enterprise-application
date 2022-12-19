package suite_helper

import (
	"github.com/pkg/errors"
)

var (
	ErrService   = errors.New("api.service_error")
	ErrService2  = errors.New("api.service_error_2")
	ErrService3  = errors.New("api.service_error_3")
	ErrService4  = errors.New("api.service_error_4")
	ErrService5  = errors.New("api.service_error_5")
	ErrRepo      = errors.New("api.repo_error")
	ErrRepoBegin = errors.New("api.repo_begin_error")
)
