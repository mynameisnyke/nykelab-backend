package services

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
)

const (
	devCollection  = "media_dev"
	prodCollection = "meida_prod"
)

type Media struct {
	Name               string   `json:"name" firestore:"name"`
	FileType           string   `json:"filetype" firestore:"filetype"`
	FileSize           string   `json:"filesize" firestore:"filesize"`
	RequiredTranscodes []string `json:"RequiredTranscodes" firestore:"required_transcodes"`
	Tags               []string `json:"tags" firestore:"tags"`
	ID                 string   `json:"id" firestore:"id"`
	Galleries          []string `json:"gallery_ids" firestore:"gallery_ids"`
}

type MediaService struct {
	db         *firestore.Client
	collection string
}

type MediaQueryInput struct {
	Tags      []string `url:"tags"`
	FileType  string   `url:"filetype"`
	Galleries []string `url:"galleries"`
}

func New(stage string) (*MediaService, error) {

	ctx := context.Background()

	// Get a client
	db, err := firestore.NewClient(ctx, "nykelab")
	if err != nil {
		return nil, err
	}

	if stage == "dev" {
		return &MediaService{db, devCollection}, nil
	}
	return &MediaService{db, prodCollection}, nil
}

// Creates a record and returns the ID
func (ms *MediaService) Create(m *Media) (*string, error) {

	// ID id generated by backend
	m.ID = uuid.New().String()

	docRef := ms.db.Collection(ms.collection).Doc(m.ID)

	_, err := docRef.Create(context.Background(), m)

	if err != nil {
		return nil, err
	}

	return &m.ID, nil
}

// Given the ID, this will apply the passed in updates
func (ms *MediaService) Update(id string, updates *[]firestore.Update) error {

	docRef := ms.db.Collection(ms.collection).Doc(id)

	_, err := docRef.Update(context.Background(), *updates)

	if err != nil {
		return err
	}

	return nil
}

// Given the ID, this will apply the passed in updates
func (ms *MediaService) Query(q *MediaQueryInput) ([]*firestore.DocumentSnapshot, error) {

	var query firestore.Query

	collRef := ms.db.Collection(ms.collection)

	fmt.Printf("QUERY:%v\n", query)

	// Check to see if filetype were supplied
	if q.FileType != "" {
		query = collRef.Where("filetype", "==", q.FileType)
	}

	// Check to see if tags were suppleid
	if q.Tags != nil {
		if q.FileType != "" {
			query = query.Where("tags", "array-contains-any", q.Tags)
		} else {
			query = collRef.Where("tags", "array-contains-any", q.Tags)
		}
	}

	// Check to see if galleries were supplied
	if q.Galleries != nil {
		if q.FileType != "" || q.Tags != nil {
			query = query.Where("galleries", "array-contains-any", q.Galleries)
		} else {
			query = collRef.Where("galleries", "array-contains-any", q.Galleries)
		}
	}

	// Run the Query
	queryResults, err := query.Documents(context.Background()).GetAll()

	if err != nil {
		return nil, err
	}

	if len(queryResults) == 0 {
		return nil, fmt.Errorf("Received 0 results for query\nTYPE:%s\nTAGS:%v\n", q.FileType, q.Tags)
	}

	return queryResults, nil
}
