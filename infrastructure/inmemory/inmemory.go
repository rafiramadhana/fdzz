package inmemory

import (
	"context"
	"fdzz/model"
)

var defaultBook = model.Book{
	ID:   1,
	Name: "Filosofi Teras",
}

type DB struct {
	data []model.Book
}

func NewDB() *DB {
	return &DB{
		data: []model.Book{defaultBook},
	}
}

func (db *DB) Create(ctx context.Context, book model.Book) (model.Book, error) {
	panic("implement me")
}

func (db *DB) Read(ctx context.Context) ([]model.Book, error) {
	return db.data, nil
}

func (db *DB) Update(ctx context.Context, id string, book model.Book) (model.Book, error) {
	panic("implement me")
}

func (db *DB) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
