package storage

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func Connect() *Storage {
	db, err := sqlx.Connect("postgres", "host=127.0.0.1 user=user password=pass dbname=test sslmode=disable")

	log.Println("connect ...")

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("database connected")

	return &Storage{db: db}
}

func (s *Storage) Disconnect() {
	s.db.Close()
}

func (s *Storage) GetDB() *sqlx.DB {
	return s.db
}
