package main

import (
    "log"
    "github.com/IBM/sarama"
)

func main() {
    // Kafka producer
    producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
    if err != nil {
        log.Fatal("Failed to start Sarama producer:", err)
    }
    defer producer.Close()

    msg := &sarama.ProducerMessage{
        Topic: "example_topic",
        Value: sarama.StringEncoder("This is a test message"),
    }
    _, _, err = producer.SendMessage(msg)
    if err != nil {
        log.Fatal("Failed to send message:", err)
    }

    // Kafka consumer
    consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
    if err != nil {
        log.Fatal("Failed to start Sarama consumer:", err)
    }
    defer consumer.Close()

    partitionConsumer, err := consumer.ConsumePartition("example_topic", 0, sarama.OffsetNewest)
    if err != nil {
        log.Fatal("Failed to start partition consumer:", err)
    }
    defer partitionConsumer.Close()

    for message := range partitionConsumer.Messages() {
        log.Printf("Consumed message offset %d: %s", message.Offset, string(message.Value))
    }
}
