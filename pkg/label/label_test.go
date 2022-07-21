package label

import (
	"fmt"
	"testing"
)

func testLabel(t *testing.T) {

	labels, err := LabelImage("gs://nykelab.appspot.com/top-travel-photographer-los-angeles-california-editorial-advertising-1.jpeg")

	if err != nil {
		fmt.Errorf("Did not work!")
	}

	if len(labels) == 0 {
		fmt.Print("Did not work")
	}

}
