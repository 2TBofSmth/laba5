package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/stellafff25/Lab5/db/sqlc"
	"github.com/stellafff25/Lab5/internal/server"
)

func main() {
	fmt.Println("Inside main is starting on port 3000...")
	connPool, err := pgxpool.New(context.Background(), "postgresql://postgres:postgres@localhost:5432/api_db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer connPool.Close()

	store := db.NewStore(connPool)

	server := server.NewServer(store)
	fmt.Println("Server is starting on port 3000...")
	server.Run(":3000")
	fmt.Println("Server stopped running")
}
