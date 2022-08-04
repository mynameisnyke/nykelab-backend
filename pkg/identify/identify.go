package identify

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"

	"cloud.google.com/go/storage"
)

// Global API clients used across function invocations.
var (
	storageClient *storage.Client
)

func init() {
	// Declare a separate err variable to avoid shadowing the client variables.
	var err error

	storageClient, err = storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}

}

func IdentifyImage(ctx context.Context, inputBucket string, name string) (*MagickInfo, error) {

	// Here we map out input blob to a reader
	readInputBlob := storageClient.Bucket(inputBucket).Object(name)
	r, err := readInputBlob.NewReader(ctx)

	if err != nil {
		log.Panicf("NewReader: %v", err)
	}
	// Use - as input and output to use stdin and stdout.
	cmd := exec.Command("convert", "-", "json:")
	cmd.Stdin = r
	stdout, err := cmd.Output()

	var metadata *MagickInfo

	if err != nil {
		log.Panicf("cmd.Run: %v", err)
	}

	err = json.Unmarshal(stdout, &metadata)

	if err != nil {
		log.Panicf("Could not unmarshall json: %v", err)
	}

	return metadata, nil

}
