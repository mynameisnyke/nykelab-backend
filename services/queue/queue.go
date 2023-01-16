package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

type QueueService struct {
	client    *pubsub.Client
	queueName string
}

type QueueServiceConfig struct {
	Project   string
	QueueName string
}

type StorageObjectData struct {
	Bucket  string `json:"bucket"`
	Name    string `json:"name"`
	Type    string `json:"contentType"`
	MediaID string `json:"metadata.media_id"`
}

type Queue struct {
	Status    string `json:"status"`
	InputUri  string `json:"input_uri"`
	OutputUri string `json:"output_uri"`
	MediaID   string `json:"media_id"`
	Type      string `json:"type"`
}

func NewQueueService(c *QueueServiceConfig) (*QueueService, error) {
	client, err := pubsub.NewClient(context.Background(), c.Project)
	if err != nil {
		log.Fatal(err)
	}
	return &QueueService{client: client, queueName: c.QueueName}, nil
}

func (ns *QueueService) Create(queue *Queue) (*string, error) {

	t := ns.client.Topic(ns.queueName)

	msgData, err := json.Marshal(queue)
	result := t.Publish(context.Background(), &pubsub.Message{
		Data: msgData,
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(context.Background())
	if err != nil {
		return nil, fmt.Errorf("pubsub: result.Get: %v", err)
	}

	return &id, nil
}
