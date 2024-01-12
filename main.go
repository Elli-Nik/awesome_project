package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func main() {
	urlExample := "postgres://awesome_project:123@localhost:5436/test_db"
	_, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Database is ready")
	}
	FullHouseDb()
}
}
