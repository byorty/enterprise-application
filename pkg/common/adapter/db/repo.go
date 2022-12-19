package db

import (
	"context"
)

type Model interface {
	GetUUID() string
}

type Repo[T Model] interface {
	DB() DB
	GetByUUID(ctx context.Context, uuid string) (*T, error)
	Save(ctx context.Context, model Model) error
	Delete(ctx context.Context, model Model) error
}

var _ Repo[Model] = (*repo[Model])(nil)

func NewRepo[T Model](
	ctx context.Context,
	db DB,
) Repo[T] {
	return &repo[T]{
		ctx: ctx,
		db:  db,
	}
}

type repo[T Model] struct {
	ctx context.Context
	db  DB
}

func (c repo[T]) DB() DB {
	return c.db
}

func (c repo[T]) GetByUUID(ctx context.Context, uuid string) (*T, error) {
	var model T
	err := c.db.Model(&model).Where("uuid = ?", uuid).Select()
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (c repo[T]) Save(ctx context.Context, model Model) error {
	var err error
	if len(model.GetUUID()) == 0 {
		err = c.db.Insert(model)
	} else {
		_, err = c.db.Model(model).WherePK().Update()
	}

	if err != nil {
		return err
	}

	return nil
}

func (c repo[T]) Delete(ctx context.Context, model Model) error {
	_, err := c.db.Model(model).Where("uuid = ?", model.GetUUID()).Delete()
	if err != nil {
		return err
	}

	return nil
}
