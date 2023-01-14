package services

import (
	"testing"
)

func TestGalleryClient(t *testing.T) {
	_, err := NewGalleryService("dev")
	if err != nil {
		t.Error(err)
	}
}

func TestGalleryCreate(t *testing.T) {

	ms, err := NewGalleryService("dev")
	if err != nil {
		t.Error(err)
	}

	_, err = ms.Create(&Gallery{
		Name: "Test_Gallery_Create",
	})

	if err != nil {
		t.Error(err)
	}
}

func TestGallkeryQuery(t *testing.T) {

	testGallery := &Gallery{
		Name:   "testing-gallery-query",
		Tags:   []string{"landscape", "dogs"},
		Aspect: "square",
	}

	ms, err := NewGalleryService("dev")
	if err != nil {
		t.Error(err)
	}

	// Create documents to query
	for i := 0; i < 10; i++ {

		_, err = ms.Create(testGallery)

		if err != nil {
			t.Error(err)
		}
	}

	t.Run("Should return 10 items by tags", func(t *testing.T) {
		output, err := ms.Query(&GalleryQueryInput{
			Tags:  testGallery.Tags,
			Limit: 10,
		})

		if err != nil {
			t.Error(err)
		}

		got := len(output.Items)
		want := 10

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("Should return 10 items by aspect", func(t *testing.T) {
		output, err := ms.Query(&GalleryQueryInput{
			Aspect: testGallery.Aspect,
			Limit:  10,
		})

		if err != nil {
			t.Error(err)
		}

		got := len(output.Items)
		want := 10

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("Should return 10 items using pagination", func(t *testing.T) {
		output, err := ms.Query(&GalleryQueryInput{
			Tags:  testGallery.Tags,
			Limit: 5,
		})

		if err != nil {
			t.Error(err)
		}

		output2, err := ms.Query(&GalleryQueryInput{
			Tags:   testGallery.Tags,
			Limit:  5,
			Offset: output.Offset,
		})

		if err != nil {
			t.Error(err)
		}

		output.Items = append(output.Items, output2.Items...)

		if output.Items[0].ID == output.Items[7].ID {
			t.Error("Duplicated items")
		}

		got := len(output.Items)
		want := 10

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
}
