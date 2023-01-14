package convert

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"cloud.google.com/go/storage"
)

// Global API clients used across function invocations.
var (
	storageClient  *storage.Client
	outputImageExt string = "webp"
)

func init() {
	// Declare a separate err variable to avoid shadowing the client variables.
	var err error

	storageClient, err = storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}

}

func ConvertImage(source string, filename string) (outputStream []byte, outputFile *string, err error) {

	var stderr bytes.Buffer

	// Just a simple GET request to the image URL
	// We get back a *Response, and an error
	res, err := http.Get(source)

	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}

	baseFileName := strings.Split(filename, ".")[0]
	outputFileName := fmt.Sprintf("%s.%s", baseFileName, outputImageExt)

	cmd := exec.Command("convert", "-", "webp:-")
	cmd.Stdin = res.Body
	cmd.Stderr = &stderr
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatalf("Could not convert %s : %s", source, stderr.String())
		return nil, nil, err

	}
	return stdout, &outputFileName, nil

}
