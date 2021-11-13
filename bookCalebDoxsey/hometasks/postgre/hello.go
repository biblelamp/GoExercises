package main

// Install all required packages:
// go get github.com/jackc/pgx
// go get github.com/jackc/pgconn
// go get github.com/jackc/pgtype

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"os"
)

func main() {
	ctx := context.Background()
	dbUrl := "postgres://postgres:root@localhost:5432/hello"

	// получение соединения
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	var result string
	sqlRequest := "select 'Hello, world!'"

	// отправка запроса и получение результата
	err = conn.QueryRow(ctx, sqlRequest).Scan(&result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
