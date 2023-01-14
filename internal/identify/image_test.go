package identify

import (
	"testing"
)

func TestImageProbe(t *testing.T) {
	url := "https://storage.googleapis.com/nl-pics/adulthood.jpg"

	_, err := IdentifyImage(url)

	if err != nil {
		t.Error(err)
	}
}
