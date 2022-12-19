package userent

import "time"

type UserProduct struct {
	UUID        string `pg:"type:uuid,pk,default:gen_random_uuid()"`
	ProductUUID string
	Count       uint32
	CreatedAt   time.Time `pg:"default:now()"`
}
