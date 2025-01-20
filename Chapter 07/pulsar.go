package main

import (
    "context"
    "log"
    "github.com/apache/pulsar-client-go/pulsar"
)

func main() {
    client, err := pulsar.NewClient(pulsar.ClientOptions{
        URL: "pulsar://localhost:6650",
    })
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // Producer
    producer, err := client.CreateProducer(pulsar.ProducerOptions{
        Topic: "my-topic",
    })
    if err != nil {
        log.Fatal(err)
    }
    defer producer.Close()

    _, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
        Payload: []byte("hello"),
    })
    if err != nil {
        log.Fatal(err)
    }

    // Consumer
    consumer, err := client.Subscribe(pulsar.ConsumerOptions{
        Topic:            "my-topic",
        SubscriptionName: "my-subscription",
        Type:             pulsar.Shared,
    })
    if err != nil {
        log.Fatal(err)
    }
    defer consumer.Close()

    msg, err := consumer.Receive(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Received message msgId: %v -- content: '%s'", msg.ID(), string(msg.Payload()))
    consumer.Ack(msg)
}
