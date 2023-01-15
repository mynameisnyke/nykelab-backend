package services

import (
	"context"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
)

type QueueService struct {
	client    *cloudtasks.Client
	queuePath string
}

func NewQueueService(queue string) (*QueueService, error) {
	ctx := context.Background()
	c, err := cloudtasks.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &QueueService{client: c, queuePath: queue}, nil
}

func (qs *QueueService) AddTask(t *cloudtaskspb.Task) error {
	_, err := qs.client.CreateTask(context.Background(), &cloudtaskspb.CreateTaskRequest{
		Parent: qs.queuePath,
		Task:   t,
	})

	if err != nil {
		return err
	}

	return nil
}
