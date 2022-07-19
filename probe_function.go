package probe_function

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/mynameisnyke/nykelab-backend/pkg/firebase"
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
			log.Panicf("Error probing file: %v", err)
		}

		videoStream := probeData.FirstVideoStream()

		docUpdates := &[]firestore.Update{
			{
				Path:  "Orientation",
				Value: videoStream.DisplayAspectRatio,
			}, {
				Path:  "Creation_Date",
				Value: videoStream.Tags.CreationTime,
			},
			{
				Path:  "Testing",
				Value: "Test worked from main",
			},
		}

		firebase.UpdateDoc(doc, docUpdates)
	default:
		log.Panicf("The supplied documnet does not contain either a video or image file: %v", doc)
	}

	return nil
}
