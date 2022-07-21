package identify

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestFibInt(t *testing.T) {

	ctx := context.Background()
	meta, err := IdentifyImage(ctx, "nykelab.appspot.com", "_7M36157.jpg")

	if err != nil {
		fmt.Errorf("%v", meta)

	}

	time, err := time.Parse("2022:04:04 16:13:40", (*meta)[0].Image.Properties["exif:DateTimeDigitized"])

	if err != nil {
		fmt.Errorf("%v", err)
	}
	t.Log(time)
}
