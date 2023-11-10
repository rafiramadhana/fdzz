package usecase

import (
	"context"
	"fdzz/infrastructure"
	"fdzz/model"
)

type BookStore struct {
	db infrastructure.DB
}

func NewBookStore(db infrastructure.DB) *BookStore {
	return &BookStore{
		db: db,
	}
}

func (bs *BookStore) Create(ctx context.Context, book model.Book) (model.Book, error) {
	panic("implement me")
}

func (bs *BookStore) Read(ctx context.Context) ([]model.Book, error) {
	return bs.db.Read(ctx)
}

func (bs *BookStore) Update(ctx context.Context, id string, book model.Book) (model.Book, error) {
	panic("implement me")
}

func (bs *BookStore) Delete(ctx context.Context, id string) (model.Book, error) {
	panic("implement me")
}
