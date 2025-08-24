package storage

import (
	"dialog-service/infrastructure/config"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func Connect(config *config.Config) *Storage {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB.PG_HOST, config.DB.PG_PORT, config.DB.PG_USER, config.DB.PG_PASS, config.DB.PG_DBNAME)

	log.Println("connect to database ...")

	db, err := sqlx.Connect("postgres", connString)

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
