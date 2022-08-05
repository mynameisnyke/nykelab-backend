package probe_function

import (
	"log"
	"testing"

	"github.com/mynameisnyke/nykelab-backend/pkg/convert"
	"github.com/mynameisnyke/nykelab-backend/pkg/storage"
)

func TestConvert(t *testing.T) {

	stdout, outputFile, err := convert.ConvertImage("/Volumes/Whiskey/Images/Meems/watsonpain.png")
	if err != nil {
		log.Panicf("%v", err)
	}

	storage.WriteFileToGCS(stdout, *outputFile)

}
