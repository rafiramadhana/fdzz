package usecase

import (
	"context"
	"fdzz/infrastructure"
	"fdzz/model"
	"log"
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
	b, err := bs.db.Read(ctx)
	if err != nil {
		log.Println("Fail to read: ", err)
	}

	return b, err
}

func (bs *BookStore) Update(ctx context.Context, id string, book model.Book) (model.Book, error) {
	panic("implement me")
}

func (bs *BookStore) Delete(ctx context.Context, id string) (model.Book, error) {
	panic("implement me")
}
