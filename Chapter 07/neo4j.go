package main

import (
    "github.com/neo4j/neo4j-go-driver/v4/neo4j"
    "log"
)

func main() {
    driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("username", "password", ""))
    if err != nil {
        log.Fatal(err)
    }
    defer driver.Close()
}
