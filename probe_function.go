package probe_function

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// 	"time"

// 	"cloud.google.com/go/firestore"
// 	"cloud.google.com/go/functions/metadata"
// 	"github.com/mynameisnyke/nykelab-backend/internal/label"
// 	"github.com/mynameisnyke/nykelab-backend/internal/storage"
// )

// // FirestoreEvent is the payload of a Firestore event.
// type FirestoreEvent struct {
// 	OldValue   FirestoreValue `json:"oldValue"`
// 	Value      FirestoreValue `json:"value"`
// 	UpdateMask struct {
// 		FieldPaths []string `json:"fieldPaths"`
// 	} `json:"updateMask"`
// }

// // FirestoreValue holds Firestore fields.
// type FirestoreValue struct {
// 	CreateTime time.Time                    `json:"createTime"`
// 	Fields     map[string]map[string]string `json:"fields"`
// 	Name       string                       `json:"name"`
// 	UpdateTime time.Time                    `json:"updateTime"`
// }

// // GOOGLE_CLOUD_PROJECT is automatically set by the Cloud Functions runtime.
// var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

// func ConvertFile(ctx context.Context, e FirestoreEvent) error {

// 	meta, err := metadata.FromContext(ctx)
// 	if err != nil {
// 		// Assume an error on the function invoker and try again.
// 		return fmt.Errorf("metadata.FromContext: %v", err)
// 	}

// 	// Ignore events that are too old.
// 	expiration := meta.Timestamp.Add(60 * time.Second)
// 	if time.Now().After(expiration) {
// 		log.Printf("event timeout: halting retries for expired event '%q'", meta.EventID)
// 		return nil
// 	}

// 	fileType := e.Value.Fields["type"]["stringValue"]
// 	src := e.Value.Fields["src"]["stringValue"]
// 	filename := e.Value.Fields["name"]["stringValue"]

// 	switch fileType {
// 	case "image":
// 		stdout, outputFile, err := convert.ConvertImage(src, filename)
// 		if err != nil {
// 			log.Panicf("%v", err)
// 		}

// 		err = storage.WriteFileToGCS(stdout, *outputFile)
// 		if err != nil {
// 			log.Panicf("%v", err)
// 		}
// 	}

// 	return nil

// }

// func ProbeFile(ctx context.Context, e FirestoreEvent) error {

// 	meta, err := metadata.FromContext(ctx)
// 	if err != nil {
// 		// Assume an error on the function invoker and try again.
// 		return fmt.Errorf("metadata.FromContext: %v", err)
// 	}

// 	// Ignore events that are too old.
// 	expiration := meta.Timestamp.Add(60 * time.Second)
// 	if time.Now().After(expiration) {
// 		log.Printf("event timeout: halting retries for expired event '%q'", meta.EventID)
// 		return nil
// 	}

// 	fullPath := strings.Split(e.Value.Name, "/documents/")[1]
// 	pathParts := strings.Split(fullPath, "/")
// 	doc := strings.Join(pathParts[1:], "/")
// 	fileType := e.Value.Fields["type"]["stringValue"]
// 	src := e.Value.Fields["src"]["stringValue"]
// 	uri := e.Value.Fields["uri"]["stringValue"]
// 	log.Printf("Here is the uri value that's being parsed: %v", uri)
// 	bucket, err := identify.ParseBucket(uri)
// 	if err != nil {
// 		log.Panicf("Could not parse bucket for %v: %v", src, err)
// 	}

// 	obj, err := identify.ParseObject(uri)
// 	if err != nil {
// 		log.Panicf("Could not parse object for %v: %v", src, err)
// 	}

// 	log.Printf("This is your bucket %s, and this is your string %s", *bucket, *obj)

// 	// Let's check what type of file it is and call the appropriate program
// 	switch fileType {
// 	case "video": // Since this is a video file, we're going to make a call to ffmpeg and update the documentt with the probe data
// 		probeData, err := probe.ProbeMedia(src)
// 		if err != nil {
// 			return fmt.Errorf("Error probing file: %v", err)
// 		}

// 		videoStream := probeData.FirstVideoStream()

// 		creationDate, err := time.Parse("2022-06-16T18:01:44.000000Z", videoStream.Tags.CreationTime)
// 		if err != nil {
// 			return fmt.Errorf("Couldnt parse date from %v: %v", doc, err)
// 		}

// 		docUpdates := &[]firestore.Update{
// 			{
// 				Path:  "orientation",
// 				Value: videoStream.DisplayAspectRatio,
// 			}, {
// 				Path:  "creation_Date",
// 				Value: creationDate,
// 			},
// 		}

// 		firebase.UpdateDoc(doc, docUpdates)
// 	case "image":

// 		// Here we generate the labes for our image
// 		labels, err := label.LabelImage(src)
// 		if err != nil {
// 			return fmt.Errorf("Error generating labels: %v", err)
// 		}

// 		// Let's find the create time using exif dataa
// 		imageMeta, err := identify.IdentifyImage(ctx, *bucket, *obj)
// 		if err != nil {
// 			return fmt.Errorf("Error parsing metada for image %v: %v", src, err)
// 		}

// 		createTime, err := time.Parse("2006:01:02 15:04:05", (*imageMeta).Image.Properties["exif:DateTimeOriginal"])
// 		if err != nil {
// 			return fmt.Errorf("Could not parse create time from the image %v: %v", src, err)
// 		}
// 		log.Printf("Here is the image create time: %v", createTime)

// 		// Let's use the width and height from the metadata to determine if this image is vertical or horizontal
// 		vertical := identify.DetermineOrientation(int(imageMeta.Image.Geometry.Width), int(imageMeta.Image.Geometry.Height))

// 		docUpdates := &[]firestore.Update{
// 			{
// 				Path:  "labels",
// 				Value: labels,
// 			},
// 			{
// 				Path:  "create_time",
// 				Value: createTime,
// 			},
// 			{
// 				Path:  "vertical",
// 				Value: vertical,
// 			},
// 		}

// 		firebase.UpdateDoc(doc, docUpdates)

// 	default:
// 		log.Panicf("The supplied documnet does not contain either a video or image file: %v", doc)
// 		return nil
// 	}

// 	return nil
// }
