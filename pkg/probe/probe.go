package probe

import (
	"context"
	"log"
	"time"

	ffprobe "gopkg.in/vansante/go-ffprobe.v2"
)

// Takes a file url, creates a readstream and probes for metadata,
func ProbeMedia(fileUrl string) (*ffprobe.ProbeData, error) {

	ctx, cancelFn := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, fileUrl)
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}

	return data, nil
}
