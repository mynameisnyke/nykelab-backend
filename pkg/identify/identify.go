package identify

import (
	"context"
	"encoding/json"
	"fmt"
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

func IdentifyImage(ctx context.Context, inputBucket string, name string) (*MagickResult, error) {

	// Here we map out input blob to a reader
	inputBlob := storageClient.Bucket(inputBucket).Object(name)
	r, err := inputBlob.NewReader(ctx)

	if err != nil {
		fmt.Errorf("NewReader: %v", err)
	}
	// Use - as input and output to use stdin and stdout.
	cmd := exec.Command("magick", "convert", "-", "json:")
	cmd.Stdin = r
	stdout, err := cmd.Output()

	var metadata MagickResult

	if err != nil {
		fmt.Errorf("cmd.Run: %v", err)
	}

	err2 := json.Unmarshal(stdout, &metadata)

	if err != nil {
		fmt.Errorf("Could not unmarshall json: %v", err2)
	}

	fmt.Print()

	return &metadata, nil
}
