package publish

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	services "github.com/mynameisnyke/nykelab-backend/services/media"
	"github.com/mynameisnyke/nykelab-backend/services/queue"
	"github.com/mynameisnyke/nykelab-backend/services/transcode"
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
	functions.CloudEvent("CreateTranscode", createTranscode)
}

// MessagePublishedData contains the full Pub/Sub message
// See the documentation for more details:
// https://cloud.google.com/eventarc/docs/cloudevents#pubsub
type MessagePublishedData struct {
	Message PubSubMessage
}

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data queue.Queue `json:"data"`
}

// publishStorageEvent consumes a CloudEvent message and logs details about the changed object.
func createTranscode(ctx context.Context, e event.Event) error {

	log.Printf("Event ID: %s", e.ID())
	log.Printf("Event Type: %s", e.Type())

	var msg MessagePublishedData
	if err := e.DataAs(&msg); err != nil {
		return fmt.Errorf("event.DataAs: %v", err)
	}

	svc, err := transcode.NewTranscodeService()

	if err != nil {
		return err
	}

	// Create the transcode job
	job, err := svc.CreateJob(&transcode.CreateJobInput{
		InputUri:  msg.Message.Data.InputUri,
		OutputUri: msg.Message.Data.OutputUri,
		Preset:    "h264",
	})

	if err != nil {
		return err
	}
	services.Media.

	// Update the DB
	return nil
}
