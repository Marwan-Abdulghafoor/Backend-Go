package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"murwan.net/fiephrs-backend/utils"
)

func StartKafka() <-chan utils.ProfileInfo {
	// Define a channel to receive messages
	messageChan := make(chan utils.ProfileInfo)

	// Start a goroutine to read messages and send them to the channel
	go func() {
		conf := kafka.ReaderConfig{
			Brokers:  []string{"localhost:9092"},
			Topic:    "FIEPHRS_TOPIC",
			GroupID:  "g1",
			MaxBytes: 10,
		}

		reader := kafka.NewReader(conf)

		defer close(messageChan) // Close the channel when the goroutine finishes

		for {
			message, err := reader.ReadMessage(context.Background())
			if err != nil {
				fmt.Println("An error occurred:", err)
				continue
			}

			var info utils.ProfileInfo
			err = json.Unmarshal(message.Value, &info)
			if err != nil {
				fmt.Println("Error unmarshaling message:", err)
				continue
			}

			// Send the profile info to the channel
			messageChan <- info
		}
	}()

	return messageChan
}
