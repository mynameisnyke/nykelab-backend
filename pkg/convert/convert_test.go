package convert

import (
	"testing"
)

func TestConvert(t *testing.T) {

	stdout, outputFile, err := ConvertImage("/Volumes/Whiskey/Images/Meems/asukabait.gif")

	WriteFileToGCS()

}
