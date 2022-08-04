package identify

import (
	"testing"
)

// func TestFibInt(t *testing.T) {

// 	ctx := context.Background()
// 	meta, err := IdentifyImage(ctx, "nykelab.appspot.com", "_7M36157.jpg")

// 	if err != nil {
// 		fmt.Errorf("%v", meta)

// 	}

// 	time, err := time.Parse("2022:04:04 16:13:40", (*meta)[0].Image.Properties["exif:DateTimeDigitized"])

// 	if err != nil {
// 		fmt.Errorf("%v", err)
// 	}
// 	t.Log(time)
// }

func TestParseBucekt(t *testing.T) {
	var path string = "gs://nykelab.appspot.com/20210208-_7M31672.jpg"

	bucket, err := ParseBucket(path)
	if err != nil {
		t.Errorf("Did not work %v", err)
	}

	if *bucket == "nykelab.appspot.com" {
		t.Errorf("Bucket name did not match: %s", *bucket)
	}
}

// func TestParseObject(t *testing.T) {
// 	var path string = "gs://nykelab.appspot.com/20210208-_7M31672.jpg"

// 	obj, err := ParseObject(path)
// 	if err != nil {
// 		t.Errorf("Did not work %v", err)
// 	}

// 	if *obj != "20210208-_7M31672.jpg" {
// 		t.Errorf("Bucket name did not match: %s", *obj)
// 	}
// }
