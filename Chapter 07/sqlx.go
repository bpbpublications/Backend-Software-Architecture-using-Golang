package main

import (
    "database/sql"
    "log"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq" // Postgres driver
)

type User struct {
    ID        int    `db:"id"`
    Name      string `db:"name"`
    Email     string `db:"email"`
    BirthDate string `db:"birth_date"`
}

func main() {
    db, err := sqlx.Connect("postgres", "user=username dbname=mydb sslmode=disable")
    if err != nil {
        log.Fatalln(err)
    }

    user := User{}
    err = db.Get(&user, "SELECT * FROM users WHERE id=$1", 1)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println(user)
}
