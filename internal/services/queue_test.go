package services

import (
	"testing"

	"cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
)

func TestQueueClient(t *testing.T) {

	_, err := NewQueueService("projects/nykelab/locations/us-east1/queues/transcodes-dev")

	if err != nil {
		t.Error(err)
	}
}

func TestQueueAdd(t *testing.T) {
	svc, err := NewQueueService("projects/nykelab/locations/us-east1/queues/transcodes-dev")

	if err != nil {
		t.Error(err)
	}

	err = svc.AddTask(&cloudtaskspb.Task{
		MessageType: &cloudtaskspb.Task_HttpRequest{
			HttpRequest: &cloudtaskspb.HttpRequest{
				Url: "http://testing.com",
			},
		},
	})

	if err != nil {
		t.Error(err)
	}
}
