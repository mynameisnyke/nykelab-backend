package transcode

import "testing"

func TestTranscodeService(t *testing.T) {
	t.Run("Should create service", func(t *testing.T) {
		_, err := NewTranscodeService(&TranscodeServiceConfig{
			ParentPath: "projects/nykelab/locations/us-east1",
		})
		if err != nil {
			t.Error(err)
		}
	})
}
func TestTranscodeJobCreate(t *testing.T) {
	svc, err := NewTranscodeService(&TranscodeServiceConfig{
		ParentPath: "projects/nykelab/locations/us-east1",
	})
	if err != nil {
		t.Error(err)
	}
	t.Run("Should create job", func(t *testing.T) {
		_, err := svc.CreateJob(&CreateJobInput{
			InputUri:  "gs://nykelab-uploads-dev/fakeid/hc.mp4",
			OutputUri: "gs://nykelab-transcodes-dev/fakeid/",
			Preset:    "h264",
		})

		if err != nil {
			t.Error(err)
		}
	})
}

func TestTranscodeGetJob(t *testing.T) {
	svc, err := NewTranscodeService(&TranscodeServiceConfig{
		ParentPath: "projects/nykelab/locations/us-east1",
	})
	if err != nil {
		t.Error(err)
	}

	job, err := svc.CreateJob(&CreateJobInput{
		InputUri:  "gs://nykelab-uploads-dev/fakeid/hc.mp4",
		OutputUri: "gs://nykelab-transcodes-dev/fakeid/",
		Preset:    "h264",
	})

	if err != nil {
		t.Error(err)
	}
	t.Run("Should create and get job details", func(t *testing.T) {

		_, err := svc.GetJob(job.Name)

		if err != nil {
			t.Error(err)
		}
	})
}
