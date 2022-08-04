package convert

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

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

func ConvertImage(source string) (outputStream []byte, err error) {

	ctx := context.Background()

	baseFileName := strings.Split(source, ".")[0]
	// cmd := exec.Command("convert", source, fmt.Sprintf("%s.webp", baseFileName))
	cmd := exec.Command("convert", source, "webp:-")
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatalf("Could not covert %s", source)
		return nil, err
	}

	wc := storageClient.Bucket("mdm-archive").Object(fmt.Sprintf("%s.webp", baseFileName)).NewWriter(ctx)

	_, err = wc.Write(stdout)
	if err != nil {
		log.Fatalf("Could not print upload %s to cloud storage", source)
		return nil, err
	}

	if err := wc.Close(); err != nil {
		log.Fatalf("Could not close out upload for %s", source)
		return nil, err
	}

	return stdout, nil
}

// func IdentifyImage(ctx context.Context, inputBucket string, name string) (*MagickInfo, error) {

// 	// Here we map out input blob to a reader
// 	readInputBlob := storageClient.Bucket(inputBucket).Object(name)
// 	r, err := readInputBlob.NewReader(ctx)

// 	if err != nil {
// 		log.Panicf("NewReader: %v", err)
// 	}
// 	// Use - as input and output to use stdin and stdout.
// 	cmd := exec.Command("convert", "", "json:")
// 	cmd.Stdin = r
// 	stdout, err := cmd.Output()

// 	var metadata *MagickInfo

// 	if err != nil {
// 		log.Panicf("cmd.Run: %v", err)
// 	}

// 	err = json.Unmarshal(stdout, &metadata)

// 	if err != nil {
// 		log.Panicf("Could not unmarshall json: %v", err)
// 	}

// 	return metadata, nil

// }
