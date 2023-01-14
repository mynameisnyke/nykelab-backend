package probe_function

import (
	"log"
	"testing"

	"github.com/mynameisnyke/nykelab-backend/pkg/convert"
	"github.com/mynameisnyke/nykelab-backend/pkg/storage"
)

func TestConvert(t *testing.T) {

	stdout, outputFile, err := convert.ConvertImage("https://firebasestorage.googleapis.com/v0/b/nykelab.appspot.com/o/_7M30503.jpg?alt=media&token=a4a575ba-b6ca-448c-bf5f-811833dc2374", "_7M30503.jpg")
	if err != nil {
		log.Panicf("%v", err)
	}

	storage.WriteFileToGCS(stdout, *outputFile)

}
