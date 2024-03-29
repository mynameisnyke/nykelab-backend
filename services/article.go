package services

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
)

const (
	articleDevCollection  = "article_dev"
	articleProdCollection = "article_prod"
)

type Article struct {
	Headline    string   `json:"name" firestore:"name"`
	Type        string   `json:"type" firestore:"type"`
	Tags        []string `json:"tags" firestore:"tags"`
	ID          string   `json:"id" firestore:"id"`
	PublishDate string   `json:"publish_date" firestore:"publish_date"`
	CreateDate  string   `json:"create_date" firestore:"create_date"`
	Published   bool     `json:"published" firestore:"published"`
}

type ArticleService struct {
	db         *firestore.Client
	collection string
}

type ArticleQueryInput struct {
	Tags     []string `url:"tags"`
	FileType string   `url:"filetype"`
}

func NewArticleService(stage string) (*ArticleService, error) {

	ctx := context.Background()

	// Get a client
	db, err := firestore.NewClient(ctx, "nykelab")
	if err != nil {
		return nil, err
	}

	if stage == "dev" {
		return &ArticleService{db, articleDevCollection}, nil
	}
	return &ArticleService{db, articleProdCollection}, nil
}

// Creates a record and returns the ID
func (as *ArticleService) Create(m *Article) (*string, error) {

	// ID id generated by backend
	m.ID = uuid.New().String()

	docRef := as.db.Collection(as.collection).Doc(m.ID)

	_, err := docRef.Create(context.Background(), m)

	if err != nil {
		return nil, err
	}

	return &m.ID, nil
}

// Given the ID, this will apply the passed in updates
func (as *ArticleService) Update(id string, updates *[]firestore.Update) error {

	docRef := as.db.Collection(as.collection).Doc(id)

	_, err := docRef.Update(context.Background(), *updates)

	if err != nil {
		return err
	}

	return nil
}

// Given the ID, this will apply the passed in updates
func (as *ArticleService) Query(q *ArticleQueryInput) ([]*firestore.DocumentSnapshot, error) {

	var query firestore.Query

	collRef := as.db.Collection(as.collection)

	if q.FileType != "" {
		query = collRef.Where("filetype", "==", q.FileType)
	}

	if q.Tags != nil {
		if q.FileType != "" {
			query = query.Where("tags", "array-contains-any", q.Tags)
		} else {
			query = collRef.Where("tags", "array-contains-any", q.Tags)
		}
	}

	queryResults, err := query.Documents(context.Background()).GetAll()

	if err != nil {
		return nil, err
	}

	if len(queryResults) == 0 {
		return nil, fmt.Errorf("Received 0 results for query\nTYPE:%s\nTAGS:%v\n", q.FileType, q.Tags)
	}

	return queryResults, nil
}
