package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

type User struct {
    gorm.Model
    Name      string
    Email     string
    BirthDate string
}

func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // Migrate the schema
    db.AutoMigrate(&User{})

    // Create
    db.Create(&User{Name: "John Doe", Email: "john.doe@example.com", BirthDate: "1990-01-01"})

    // Read
    var user User
    db.First(&user, 1) // find user with integer primary key
    log.Println(user)

    // Update
    db.Model(&user).Update("Name", "Jane Doe")

    // Delete
    db.Delete(&user, 1)
}
