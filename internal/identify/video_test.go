package identify

import (
	"testing"
)

func TestProbe(t *testing.T) {

	url := "https://storage.googleapis.com/nl-video-welcome-full-proxies/20201105_080255.webm"

	t.Run("Should return metatadata", func(t *testing.T) {
		_, err := Probe(url)

		if err != nil {
			t.Error(err)
		}

	})

	t.Run("Should not return metadata", func(t *testing.T) {
		_, err := Probe("https://test.com")

		if err == nil {
			t.Errorf("Expected error but got nil")
		}
	})
}
