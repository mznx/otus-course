package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID         uint64 `db:"id"`
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	BirthDate  string `db:"birthdate"`
	City       string `db:"city"`
}

func getDbConn() *sqlx.DB {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASS"), os.Getenv("PG_DB"))

	db, err := sqlx.Connect("postgres", connString)

	if err != nil {
		panic(err)
	}

	return db
}

func processSearchMethod(db *sqlx.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		firstName := r.URL.Query().Get("first_name")
		secondName := r.URL.Query().Get("second_name")

		var res []User

		err := db.Select(&res, "SELECT * FROM users WHERE first_name LIKE $1 AND second_name LIKE $2", "%"+firstName+"%", "%"+secondName+"%")

		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("OK"))
		}

	}
}

func main() {
	db := getDbConn()

	router := chi.NewRouter()

	router.Get("/user/search", processSearchMethod(db))

	http.ListenAndServe(":3000", router)
}
