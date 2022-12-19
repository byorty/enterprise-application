package userrp

import (
	"github.com/byorty/enterprise-application/pkg/common/adapter/db"
	userent "github.com/byorty/enterprise-application/pkg/user/domain/entity"
)

type UserProductRepo interface {
	db.Repo[userent.UserProduct]
}
