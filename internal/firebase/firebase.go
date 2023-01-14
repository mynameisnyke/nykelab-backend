package firebase

import (
	"context"

	"cloud.google.com/go/firestore"
)

// -- FIRESTORE -- //

func CreateClient() (*firestore.Client, error) {
	ctx := context.Background()

	// Get a client
	client, err := firestore.NewClient(ctx, "nykelab")
	if err != nil {
		panic(err)
	}

	return client, nil
}

func CreateDoc(doc *VideoAssetProbeUpdate) error {

	ctx := context.Background()

	// Get a client
	client, err := firestore.NewClient(ctx, "nykelab")
	if err != nil {
		panic(err)
	}

	docRef := client.Collection("assets").Doc((*doc).ID)

	_, err2 := docRef.Create(ctx, *doc)

	if err2 != nil {
		panic(err2)
	}

	return nil
}

func UpdateDoc(id string, updates *[]firestore.Update) error {

	ctx := context.Background()

	// Get a client
	client, err := firestore.NewClient(ctx, "nykelab")
	if err != nil {
		panic(err)
	}

	docRef := client.Collection("assets").Doc(id)

	_, err2 := docRef.Update(ctx, *updates)

	if err2 != nil {
		panic(err2)
	}

	return nil
}
