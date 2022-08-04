package identify

import (
	"strings"
)

func ParseBucket(uri string) (b *string, err error) {

	var bucket string
	fullUri := strings.Split(uri, "gs://")[1]
	if len(fullUri) <= 0 {
		return nil, err
	}
	uriParts := strings.Split(fullUri, "/")
	bucket = uriParts[0]

	return &bucket, nil
}
func ParseObject(uri string) (o *string, err error) {

	var object string
	fullUri := strings.Split(uri, "gs://")[1]
	if len(fullUri) <= 0 {
		return nil, err
	}
	uriParts := strings.Split(fullUri, "/")
	object = strings.Join(uriParts[1:], "/")

	return &object, nil
}
