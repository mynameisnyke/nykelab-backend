package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/functions/metadata"
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
	Fields     interface{} `json:"fields"`
	Name       string      `json:"name"`
	UpdateTime time.Time   `json:"updateTime"`
}

// MyData represents a value from Firestore. The type definition depends on the
// format of your database.
type MyData struct {
	Original struct {
		StringValue string `json:"stringValue"`
	} `json:"original"`
}

// GOOGLE_CLOUD_PROJECT is automatically set by the Cloud Functions runtime.
var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

// MakeUpperCase is triggered by a change to a Firestore document. It updates
// the `original` value of the document to upper case.
func ProbeFile(ctx context.Context, e FirestoreEvent) error {
	meta, err := metadata.FromContext(ctx)

	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	fullPath := strings.Split(e.Value.Name, "/documents/")[1]
	pathParts := strings.Split(fullPath, "/")
	doc := strings.Join(pathParts[1:], "/")

	log.Printf("Doc %v", doc)
	log.Printf("Function triggered by change to: %v", meta.Resource)
	log.Printf("Old value: %+v", e.OldValue)
	log.Printf("Fields: %+v", e.Value.Fields)

	// probeData, err := probe.ProbeMedia("https://firebasestorage.googleapis.com/v0/b/nykelab.appspot.com/o/C0002.MP4?alt=media&token=81541462-2bf1-4779-a802-5da8f1bd3ce9")
	// if err != nil {
	// 	log.Panicf("Error probing file: %v", err)
	// }

	// videoStream := probeData.FirstVideoStream()

	// docUpdates := &[]firestore.Update{
	// 	{
	// 		Path:  "Orientation",
	// 		Value: videoStream.DisplayAspectRatio,
	// 	}, {
	// 		Path:  "Creation_Date",
	// 		Value: videoStream.Tags.CreationTime,
	// 	},
	// 	{
	// 		Path:  "Testing",
	// 		Value: "Test worked from main",
	// 	},
	// }

	// firebase.UpdateDoc(doc, docUpdates)

	// data := map[string]string{"original": newValue}

	// //
	// _, err := client.Collection(collection).Doc(doc).Set(ctx, data)
	// if err != nil {
	// 	return fmt.Errorf("Set: %v", err)
	// }
	return nil
}
