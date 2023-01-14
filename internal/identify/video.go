package identify

import (
	"context"

	"gopkg.in/vansante/go-ffprobe.v2"
)

func Probe(url string) (*ffprobe.ProbeData, error) {

	data, err := ffprobe.ProbeURL(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return data, nil
}
