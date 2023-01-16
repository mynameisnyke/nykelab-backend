package queue

import "testing"

func TestQueueService(t *testing.T) {

	t.Run("Should create an queue service", func(t *testing.T) {
		_, err := NewQueueService(&QueueServiceConfig{Project: "nykelab", QueueName: "transcodes-dev"})

		if err != nil {
			t.Error(err)
		}
	})

}

func TestQueueCreatTopic(t *testing.T) {

	t.Run("Should create a new queue topic", func(t *testing.T) {
		svc, err := NewQueueService(&QueueServiceConfig{Project: "nykelab", QueueName: "transcodes-dev"})

		if err != nil {
			t.Error(err)
		}
		_, err = svc.Create(&Queue{
			MediaID: "test-1",
			Status:  "created",
			Type:    "image/jpg",
		})

		if err != nil {
			t.Error(err)
		}

	})
}
