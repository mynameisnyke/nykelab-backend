package probe_function

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/functions/metadata"
	"github.com/mynameisnyke/nykelab-backend/pkg/firebase"
	"github.com/mynameisnyke/nykelab-backend/pkg/label"
	"github.com/mynameisnyke/nykelab-backend/pkg/probe"
)

// FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	CreateTime time.Time `json:"createTime"`
	// Fields is the data for this value. The type depends on the format of your
	// database. Log an interface{} value and inspect the result to see a JSON
	// representation of your database fields.
	Fields     map[string]map[string]string `json:"fields"`
	Name       string                       `json:"name"`
	UpdateTime time.Time                    `json:"updateTime"`
}

// MyData represents a value from Firestore. The type definition depends on the
// format of your database.

// GOOGLE_CLOUD_PROJECT is automatically set by the Cloud Functions runtime.
var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

// MakeUpperCase is triggered by a change to a Firestore document. It updates
// the `original` value of the document to upper case.
func ProbeFile(ctx context.Context, e FirestoreEvent) error {

	meta, err := metadata.FromContext(ctx)
	if err != nil {
		// Assume an error on the function invoker and try again.
		return fmt.Errorf("metadata.FromContext: %v", err)
	}

	// Ignore events that are too old.
	expiration := meta.Timestamp.Add(60 * time.Second)
	if time.Now().After(expiration) {
		log.Printf("event timeout: halting retries for expired event '%q'", meta.EventID)
		return nil
	}

	fullPath := strings.Split(e.Value.Name, "/documents/")[1]
	pathParts := strings.Split(fullPath, "/")
	doc := strings.Join(pathParts[1:], "/")
	fileType := e.Value.Fields["type"]["stringValue"]
	src := e.Value.Fields["src"]["stringValue"]

	// Let's check what type of file it is and callt he appropriate program
	switch fileType {
	case "video": // Since this is a video file, we're going to make a call to ffmpeg and update the documnet with the probe data
		probeData, err := probe.ProbeMedia(src)
		if err != nil {
			return fmt.Errorf("Error probing file: %v", err)
		}

		videoStream := probeData.FirstVideoStream()

		myDate, err := time.Parse("2022-06-16T18:01:44.000000Z", videoStream.Tags.CreationTime)
		if err != nil {
			return fmt.Errorf("Couldnt parse date from %v: %v", doc, err)
		}

		docUpdates := &[]firestore.Update{
			{
				Path:  "orientation",
				Value: videoStream.DisplayAspectRatio,
			}, {
				Path:  "creation_Date",
				Value: myDate,
			},
		}

		firebase.UpdateDoc(doc, docUpdates)
	case "image":

		// Here we generate the labes for our image
		labels, err := label.LabelImage(src)
		if err != nil {
			return fmt.Errorf("Error generating labels: %v", err)
		}

		docUpdates := &[]firestore.Update{
			{
				Path:  "labels",
				Value: labels,
			},
		}

		firebase.UpdateDoc(doc, docUpdates)

	default:
		log.Panicf("The supplied documnet does not contain either a video or image file: %v", doc)
		return nil
	}

	return nil
}
