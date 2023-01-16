package transcode

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"

// 	"cloud.google.com/go/pubsub"
// 	"github.com/mynameisnyke/nykelab-backend/internal/upload"
// )

// type TranscodeService struct {
// 	client    *pubsub.Client
// 	queueName string
// }

// type TranscodeServiceConfig struct {
// 	Project   string
// 	QueueName string
// }

// type Transcode struct {
// 	Status  string `json:"status"`
// 	MediaID string `json:"media_id"`
// 	Type    string `json:"type"`
// 	Bucket  string `json:"bucket"`
// 	Name    string `json:"name"`
// }

// func NewTranscodeService(c *TranscodeServiceConfig) (*TranscodeService, error) {
// 	client, err := pubsub.NewClient(context.Background(), c.Project)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return &TranscodeService{client: client, queueName: c.QueueName}, nil
// }

// func (ns *TranscodeService) Create(upload *upload.Upload) (*string, error) {

// 	t := ns.client.Topic(ns.queueName)

// 	msgData, err := json.Marshal(transcode)
// 	result := t.Publish(context.Background(), &pubsub.Message{
// 		Data: msgData,
// 	})
// 	// Block until the result is returned and a server-generated
// 	// ID is returned for the published message.
// 	id, err := result.Get(context.Background())
// 	if err != nil {
// 		return nil, fmt.Errorf("pubsub: result.Get: %v", err)
// 	}

// 	return &id, nil
// }
