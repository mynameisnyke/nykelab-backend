package publish

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/mynameisnyke/nykelab-backend/services/queue"
)

var (
	queueName         string
	project           string
	videoOutputBucket string
)

func init() {
	queueName = os.Getenv("QUEUE_NAME")
	project = os.Getenv("PROJECT")
	videoOutputBucket = os.Getenv("VIDEO_OUTPUT_BUCKET")
	functions.CloudEvent("PublishtStorageEvent", publishStorageEvent)
}

// StorageObjectData contains metadata of the Cloud Storage object.
type StorageObjectData struct {
	Bucket         string    `json:"bucket,omitempty"`
	Name           string    `json:"name,omitempty"`
	Metageneration int64     `json:"metageneration,string,omitempty"`
	TimeCreated    time.Time `json:"timeCreated,omitempty"`
	ContentType    string    `json:"contentType,omitempty"`
	MediaID        string    `json:"metadata.media_id"`
}

// publishStorageEvent consumes a CloudEvent message and logs details about the changed object.
func publishStorageEvent(ctx context.Context, e event.Event) error {

	log.Printf("Event ID: %s", e.ID())
	log.Printf("Event Type: %s", e.Type())

	var data StorageObjectData
	if err := e.DataAs(&data); err != nil {
		return fmt.Errorf("event.DataAs: %v", err)
	}

	switch strings.Split(data.ContentType, "/")[0] {
	case "video":
		svc, err := queue.NewQueueService(&queue.QueueServiceConfig{
			Project:   project,
			QueueName: queueName,
		})

		if err != nil {
			log.Panic(err)
		}

		_, err = svc.Create(&queue.Queue{
			Status:    "initialized",
			InputUri:  fmt.Sprintf("gs://%s/%s", data.Bucket, data.Name),
			OutputUri: fmt.Sprintf("gs://%s/%s", videoOutputBucket, data.MediaID),
			MediaID:   data.MediaID,
			Type:      "transcode",
		})

		if err != nil {
			log.Panic(err)
		}

		return nil
	}

	fmt.Printf("Type was empty")

	return nil
}
