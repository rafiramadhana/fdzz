package sqlite

import (
	"context"
	"fdzz/model"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sqlx.DB
}

func NewDB() *DB {
	// TODO: Refactor to use a .db file
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Fail to open sqlite3: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Fail to ping: ", err)
	}

	// TODO: Add more proper migration
	err = migrate(db)
	if err != nil {
		log.Fatal("Fail to migrate: ", err)
	}

	return &DB{db}
}

func (db *DB) Create(ctx context.Context, book model.Book) (model.Book, error) {
	panic("implement me")
}

func (db *DB) Read(ctx context.Context) ([]model.Book, error) {
	b := []model.Book{}

	err := db.Select(&b, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (db *DB) Update(ctx context.Context, id string, book model.Book) (model.Book, error) {
	panic("implement me")
}

func (db *DB) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func migrate(db *sqlx.DB) error {
	q := `
CREATE TABLE books (
	id INTEGER PRIMARY KEY,
	name TEXT
);`

	_, err := db.Exec(q)
	if err != nil {
		return err
	}

	q2 := `
INSERT INTO books (name)
VALUES
	('Tiwas tak gondeli tenanan'),
	('Sayangku wes ora kurang kurang');`

	_, err = db.Exec(q2)
	if err != nil {
		return err
	}

	return nil
}
