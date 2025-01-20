// Generate the models and resolvers using GQLGen
//go:generate go run -mod=mod github.com/99designs/gqlgen
/* Define your GraphQL schema in schema.graphql
type User {
    id: ID!
    name: String!
    email: String!
    birthDate: String!
}
*/
package main

import (
    "github.com/99designs/gqlgen/handler"
    "net/http"
    "log"
)

func main() {
    http.Handle("/", handler.Playground("GraphQL playground", "/query"))
    http.Handle("/query", handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{}})))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
