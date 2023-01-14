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
		Name:     "test.jpg",
		Tags:     []string{"landscape", "dogs"},
	}

	ms, err := New("dev")
	if err != nil {
		t.Error(err)
	}

	_, err = ms.Create(testMedia)

	if err != nil {
		t.Error(err)
	}

	t.Run("Should return items by filetype", func(t *testing.T) {
		items, err := ms.Query(&MediaQueryInput{
			FileType: testMedia.FileType,
		})

		if err != nil {
			t.Error(err)
		}

		got := len(items)
		want := 1

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("Should return items by tags", func(t *testing.T) {
		items, err := ms.Query(&MediaQueryInput{
			Tags: testMedia.Tags,
		})

		if err != nil {
			t.Error(err)
		}

		got := len(items)
		want := 1

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
}
