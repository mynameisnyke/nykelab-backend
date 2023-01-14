package services

import (
	"testing"

	"cloud.google.com/go/firestore"
)

func TestMediaClient(t *testing.T) {
	_, err := New("dev")
	if err != nil {
		t.Error(err)
	}
}

func TestCreate(t *testing.T) {

	ms, err := New("dev")
	if err != nil {
		t.Error(err)
	}

	_, err = ms.Create(&Media{
		FileType: "image/jpg",
		Name:     "test.jpg",
	})

	if err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {

	updates := []firestore.Update{
		{Path: "FileType", Value: 10},
	}

	testMedia := &Media{
		FileType: "image/jpg",
		Name:     "test.jpg",
	}

	ms, err := New("dev")
	if err != nil {
		t.Error(err)
	}

	id, err := ms.Create(testMedia)

	if err != nil {
		t.Error(err)
	}

	err = ms.Update(*id, &updates)

	if err != nil {
		t.Error(err)
	}
}

func TestQuery(t *testing.T) {

	testMedia := &Media{
		FileType: "image/test-query",
		Name:     "query.jpg",
		Tags:     []string{"landscape", "dogs"},
	}

	ms, err := New("dev")
	if err != nil {
		t.Error(err)
	}

	// Create documents to query
	for i := 0; i < 10; i++ {

		_, err = ms.Create(testMedia)

		if err != nil {
			t.Error(err)
		}
	}

	t.Run("Should return items by filetype", func(t *testing.T) {
		output, err := ms.Query(&MediaQueryInput{
			FileType: testMedia.FileType,
		})

		if err != nil {
			t.Error(err)
		}

		got := len(output.Items)
		want := 5

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("Should return items by tags", func(t *testing.T) {
		output, err := ms.Query(&MediaQueryInput{
			Tags: testMedia.Tags,
		})

		if err != nil {
			t.Error(err)
		}

		got := len(output.Items)
		want := 5

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("Should return all 10 test items", func(t *testing.T) {
		output, err := ms.Query(&MediaQueryInput{
			Tags:  testMedia.Tags,
			Limit: 5,
		})

		if err != nil {
			t.Error(err)
		}

		output2, err := ms.Query(&MediaQueryInput{
			Tags:   testMedia.Tags,
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
