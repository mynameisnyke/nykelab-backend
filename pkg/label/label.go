package label

import (
	"context"
	"fmt"

	vision "cloud.google.com/go/vision/apiv1"
)

type Labels []string

func LabelImage(uri string) (Labels, error) {

	var visionLabels Labels

	// This is how we init the client
	ctx := context.Background()
	c, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		// TODO: Handle error.
		fmt.Errorf(" Could not create annotation client: %v", err)
		return visionLabels, nil
	}
	defer c.Close()

	img := vision.NewImageFromURI(uri)

	// Here we actually make the request for the labels
	label_entities, err := c.DetectLabels(ctx, img, nil, 10)
	if err != nil {
		fmt.Errorf(" Could not generate labels for %v: %v", uri, err)
		return visionLabels, err
	}

	if len(label_entities) < 2 {
		return visionLabels, fmt.Errorf("Could not generate labels")
	} else {
		for _, annotation := range label_entities {
			// We'll append the labells to our array
			visionLabels = append(visionLabels, annotation.Description)
		}
	}

	return visionLabels, nil

}
