// Run with `go run graphql.go` and test
// test with `curl -XPOST -d '{"query": "{ hello }"}' localhost:8080/graphql` in another terminal
package main

import (
    "log"
    "net/http"

    "github.com/graph-gophers/graphql-go"
    "github.com/graph-gophers/graphql-go/relay"
)

// Schema definition
const schemaString = `
    type Query {
        hello: String!
    }
`

// Resolver
type resolver struct{}

func (r *resolver) Hello() string {
    return "Hello, world!"
}

func main() {
    // Parse schema
    schema := graphql.MustParseSchema(schemaString, &resolver{})

    // Set up HTTP handler
    http.Handle("/graphql", &relay.Handler{Schema: schema})

    // Start HTTP server
    log.Println("Now server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}