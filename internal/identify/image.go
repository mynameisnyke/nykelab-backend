package identify

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
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

type MagickResult []MagickInfo

type MagickInfo struct {
	Image Image `json:"image"`
}

type Image struct {
	Name              string            `json:"name"`
	Format            string            `json:"format"`
	FormatDescription string            `json:"formatDescription"`
	MIMEType          string            `json:"mimeType"`
	Class             string            `json:"class"`
	Geometry          Geometry          `json:"geometry"`
	Resolution        PrintSize         `json:"resolution"`
	PrintSize         PrintSize         `json:"printSize"`
	Units             string            `json:"units"`
	Type              string            `json:"type"`
	Endianess         string            `json:"endianess"`
	Colorspace        string            `json:"colorspace"`
	Depth             int64             `json:"depth"`
	BaseDepth         int64             `json:"baseDepth"`
	Pixels            int64             `json:"pixels"`
	RenderingIntent   string            `json:"renderingIntent"`
	Gamma             float64           `json:"gamma"`
	BackgroundColor   string            `json:"backgroundColor"`
	BorderColor       string            `json:"borderColor"`
	MatteColor        string            `json:"matteColor"`
	TransparentColor  string            `json:"transparentColor"`
	Interlace         string            `json:"interlace"`
	Intensity         string            `json:"intensity"`
	Compose           string            `json:"compose"`
	PageGeometry      Geometry          `json:"pageGeometry"`
	Dispose           string            `json:"dispose"`
	Iterations        int64             `json:"iterations"`
	Compression       string            `json:"compression"`
	Quality           int64             `json:"quality"`
	Orientation       string            `json:"orientation"`
	Properties        map[string]string `json:"properties"`
	Tainted           bool              `json:"tainted"`
	Filesize          string            `json:"filesize"`
	NumberPixels      string            `json:"numberPixels"`
	PixelsPerSecond   string            `json:"pixelsPerSecond"`
	UserTime          string            `json:"userTime"`
	ElapsedTime       string            `json:"elapsedTime"`
	Version           string            `json:"version"`
}

type Geometry struct {
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
	X      int64 `json:"x"`
	Y      int64 `json:"y"`
}

type PrintSize struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func IdentifyImage(url string) (*MagickResult, error) {

	resp, err := http.Get(url)
	if err != nil {
		// handle error
		return nil, err
	}
	defer resp.Body.Close()

	// Use - as input and output to use stdin and stdout.
	cmd := exec.Command("convert", "-", "json:")
	cmd.Stdin = resp.Body
	stdout, err := cmd.Output()

	var metadata *MagickResult

	if err != nil {
		log.Panicf("cmd.Run: %v", err)
	}

	err = json.Unmarshal(stdout, &metadata)

	if err != nil {
		log.Panicf("Could not unmarshall json: %v", err)
	}

	return metadata, nil

}
