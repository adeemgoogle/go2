package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func init() {
	// Инициализация подключения к базе данных
	var err error
	Db, err = sqlx.Connect("postgres", "user=postgres password=123 dbname=postgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}

}
