package storage

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
)

var storageClient *storage.Client

func init() {
	// Declare a separate err variable to avoid shadowing the client variables.
	var err error

	storageClient, err = storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}

}

func WriteFileToGCS(stream []byte, key string) (err error) {

	wc := storageClient.Bucket("mdm-archive").Object(key).NewWriter(context.Background())

	_, err = wc.Write(stream)
	if err != nil {
		log.Fatalf("Could not print upload %s to cloud storage", key)
		return err
	}

	if err := wc.Close(); err != nil {
		log.Fatalf("Could not close out upload for %s", key)
		return err
	}

	return nil

}
