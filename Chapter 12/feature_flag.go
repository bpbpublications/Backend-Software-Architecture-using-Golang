package main

import (
    "fmt"
    ffclient "github.com/thomaspoignant/go-feature-flag"
    "github.com/thomaspoignant/go-feature-flag/ffcontext"
    "github.com/thomaspoignant/go-feature-flag/retriever"
    "github.com/thomaspoignant/go-feature-flag/retriever/fileretriever"
    "time"
)

func main() {
    client, err := ffclient.New(ffclient.Config{
        PollingInterval: 10 * time.Second,
        Retrievers: []retriever.Retriever{
           &fileretriever.Retriever{Path: "./flags.yaml"},
        },
    })
    if err != nil {
        panic(err)
    }
    defer client.Close()
    user := ffcontext.NewEvaluationContext("new-feature-key")
    enabled, err := client.BoolVariationDetails("new-feature", user, false)
    if enabled.Value {
        fmt.Println("The new feature is enabled!")
    } else {
        fmt.Println("The new feature is disabled!")
    }
}
