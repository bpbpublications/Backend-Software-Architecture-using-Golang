// Generate the models using the Ent code generator
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema

package main

import (
    "context"
    "log"
    "entgo.io/ent"
)

func main() {
    client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }
    defer client.Close()
    ctx := context.Background()

    // Run the auto migration tool.
    if err := client.Schema.Create(ctx); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

    // Create a new user.
    user, err := client.User.
        Create().
        SetName("John Doe").
        SetEmail("john.doe@example.com").
        SetBirthDate("1990-01-01").
        Save(ctx)
    if err != nil {
        log.Fatalf("failed creating user: %v", err)
    }
    log.Println("user was created:", user)
}
