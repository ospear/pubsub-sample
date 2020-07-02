package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Start")
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %+v\n", err)
		return
	}

	projectID := os.Getenv("GCP_PROJECT")
	topicID := os.Getenv("TOPIC_ID")
	n := 10

	if err := publishThatScales(projectID, topicID, n); err != nil {
		fmt.Printf("Failed: %+v\n", err)
	}
}

func publishThatScales(projectID, topicID string, n int) error {
	fmt.Printf("%s %s %d\n", projectID, topicID, n)
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	var wg sync.WaitGroup
	var totalErrors uint64
	t := client.Topic(topicID)

	for i := 0; i < n; i++ {
		fmt.Printf("loop %d\n", i)
		result := t.Publish(ctx, &pubsub.Message{
			Data: []byte("Message " + strconv.Itoa(i)),
		})

		wg.Add(1)
		go func(i int, res *pubsub.PublishResult) {
			defer wg.Done()
			// The Get method blocks until a server-generated ID or
			// an error is returned for the published message.
			id, err := res.Get(ctx)
			if err != nil {
				// Error handling code can be added here.
				fmt.Printf("Failed to publish: %v\n", err)
				atomic.AddUint64(&totalErrors, 1)
				return
			}
			fmt.Printf("Published message %d; msg ID: %v\n", i, id)
		}(i, result)
	}

	wg.Wait()

	if totalErrors > 0 {
		return fmt.Errorf("%d of %d messages did not publish successfully", totalErrors, n)
	}
	return nil
}
