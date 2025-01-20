package main

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Need valid user/password and db name
    db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
}
