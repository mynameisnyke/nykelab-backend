package services

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
)

const (
	mediaDevCollection  = "media_dev"
	mediaProdCollection = "meida_prod"
	defaultMediaLimit   = 5
)

type Media struct {
	Name               string    `json:"name" firestore:"name"`
	FileType           string    `json:"filetype" firestore:"filetype"`
	FileSize           string    `json:"filesize" firestore:"filesize"`
	RequiredTranscodes []string  `json:"required_transcodes" firestore:"required_transcodes"`
	Tags               []string  `json:"tags" firestore:"tags"`
	ID                 string    `json:"id" firestore:"id"`
	Galleries          []string  `json:"gallery_ids" firestore:"gallery_ids"`
	DateAdded          time.Time `json:"date_added" firestore:"date_added"`
	DateCreated        time.Time `json:"date_created" firestore:"date_created,omitempty"`
	Status             string    `json:"status" firestore:"status"`
}

type MediaService struct {
	db         *firestore.Client
	collection string
}

type MediaQueryInput struct {
	Tags      []string  `url:"tags"`
	FileType  string    `url:"filetype"`
	Galleries []string  `url:"galleries"`
	Limit     int       `url:"limi"`
	Offset    time.Time `url:"offset"`
}

type MediaQueryOutput struct {
	Offset time.Time
	Items  []*Media
}

func NewMediaService(stage string) (*MediaService, error) {

	ctx := context.Background()

	// Get a client
	db, err := firestore.NewClient(ctx, "nykelab")
	if err != nil {
		return nil, err
	}

	if stage == "dev" {
		return &MediaService{db, mediaDevCollection}, nil
	}
	return &MediaService{db, mediaProdCollection}, nil
}

// Creates a record and returns the ID
func (ms *MediaService) Create(m *Media) (*string, error) {

	// ID id generated by backend
	m.ID = uuid.New().String()
	m.DateAdded = time.Now()

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
func (ms *MediaService) Query(q *MediaQueryInput) (*MediaQueryOutput, error) {

	var mediaItems []*Media

	// Check for pagination token
	if !q.Offset.IsZero() {
		fmt.Print("OFFSET PROVIDED\n\n")
		//  Check to see if limit was supplied, if not will apply default
		if q.Limit == 0 {
			q.Limit = defaultMediaLimit
		}
		var query firestore.Query

		collRef := ms.db.Collection(ms.collection)

		// Check to see if filetype were supplied
		if q.FileType != "" {
			query = collRef.Where("filetype", "==", q.FileType).Limit(q.Limit).StartAfter(q.Offset).OrderBy("date_added", firestore.Desc)
		}

		// Check to see if tags were suppleid
		if q.Tags != nil {
			if q.FileType != "" {
				query = query.Where("tags", "array-contains-any", q.Tags).Limit(q.Limit).StartAfter(q.Offset).OrderBy("date_added", firestore.Desc)
			} else {
				query = collRef.Where("tags", "array-contains-any", q.Tags).Limit(q.Limit).StartAfter(q.Offset).OrderBy("date_added", firestore.Desc)
			}
		}

		// Check to see if galleries were supplied
		if q.Galleries != nil {
			if q.FileType != "" || q.Tags != nil {
				query = query.Where("galleries", "array-contains-any", q.Galleries).StartAfter(q.Offset).OrderBy("date_added", firestore.Desc)
			} else {
				query = collRef.Where("galleries", "array-contains-any", q.Galleries).StartAfter(q.Offset).OrderBy("date_added", firestore.Desc)
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

		// Convert the snapshots into media structs
		for _, v := range queryResults {
			var m Media
			v.DataTo(&m)
			mediaItems = append(mediaItems, &m)
		}

		// Grab the timestamp of the last retuned doc for offset
		offset := mediaItems[len(mediaItems)-1].DateAdded

		return &MediaQueryOutput{Offset: offset, Items: mediaItems}, nil
	}

	//  Check to see if limit was supplied, if not will apply default
	if q.Limit == 0 {
		q.Limit = defaultMediaLimit
	}
	var query firestore.Query

	collRef := ms.db.Collection(ms.collection)

	// Check to see if filetype were supplied
	if q.FileType != "" {
		query = collRef.Where("filetype", "==", q.FileType).Limit(q.Limit)
	}

	// Check to see if tags were suppleid
	if q.Tags != nil {
		if q.FileType != "" {
			query = query.Where("tags", "array-contains-any", q.Tags).Limit(q.Limit)
		} else {
			query = collRef.Where("tags", "array-contains-any", q.Tags).Limit(q.Limit)
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

	// Convert the snapshots into media structs
	for _, v := range queryResults {
		var m Media
		v.DataTo(&m)
		mediaItems = append(mediaItems, &m)
	}

	// Grab the timestamp of the last retuned doc for offset
	offset := mediaItems[len(mediaItems)-1].DateAdded

	return &MediaQueryOutput{Offset: offset, Items: mediaItems}, nil
}
