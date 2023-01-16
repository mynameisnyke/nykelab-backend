package publish

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/mynameisnyke/nykelab-backend/services/media"
	"github.com/mynameisnyke/nykelab-backend/services/queue"
	"github.com/mynameisnyke/nykelab-backend/services/transcode"
)

var (
	collection        string
	project           string
	videoOutputBucket string
	parentPath        string
)

func init() {
	collection = os.Getenv("COLLECTION")
	project = os.Getenv("PROJECT")
	parentPath = os.Getenv("PARENT_PATH")
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
	Data string `json:"data"`
}

// publishStorageEvent consumes a CloudEvent message and logs details about the changed object.
func createTranscode(ctx context.Context, e event.Event) error {

	log.Printf("Event ID: %s", e.ID())
	log.Printf("Event Type: %s", e.Type())

	var msg MessagePublishedData
	if err := e.DataAs(&msg); err != nil {
		return fmt.Errorf("event.DataAs: %v", err)
	}

	svc, err := transcode.NewTranscodeService(&transcode.TranscodeServiceConfig{
		ParentPath: parentPath,
	})

	if err != nil {
		return err
	}

	var q queue.Queue
	data, err := base64.StdEncoding.DecodeString(msg.Message.Data)
	err = json.Unmarshal(data, &q)

	// Create the transcode job
	job, err := svc.CreateJob(&transcode.CreateJobInput{
		InputUri:  q.InputUri,
		OutputUri: q.OutputUri,
		Preset:    "h264",
	})

	if err != nil {
		return err
	}

	ms, err := media.NewMediaService(&media.MediaServiceConfig{
		Project:    project,
		Collection: collection})

	var updates []firestore.Update
	updates = append(updates, firestore.Update{Path: "transcode_id", Value: job.Name})
	ms.Update(q.MediaID, &updates)
	return nil
}
