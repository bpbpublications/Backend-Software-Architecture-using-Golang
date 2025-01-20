package main

import (
    "context"
    "github.com/jackc/pgx/v4"
    "log"
)

func main() {
    // Need valid user/password and db name
    conn, err := pgx.Connect(context.Background(), "postgresql://username:password@localhost:5432/database_name")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close(context.Background())
}
