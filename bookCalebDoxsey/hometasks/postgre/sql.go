package main

import (
	"database/sql"
	"fmt"
	"log"
	// Импорт пакета необходим для неявной инициализации драйвера
	_ "github.com/jackc/pgx/stdlib"
)
func main() {
	url := "postgres://postgres:root@localhost:5432/elza"

	// sql.Open открывает соединение, используя драйвер pgx
	db, err := sql.Open("pgx", url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var greeting string
	sqlRequest := "select 'Hello, world!'"

	err = db.QueryRow(sqlRequest).Scan(&greeting)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(greeting)
}
