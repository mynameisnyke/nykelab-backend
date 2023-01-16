package services

import (
	"context"
	"fmt"

	transcoder "cloud.google.com/go/video/transcoder/apiv1"
	"cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
)

const (
	transcoderDevProj      = "nykelab"
	transcoderProdProj     = "nykelab"
	transcoderDevLocation  = "us-east1"
	transcoderProdLocation = "us-east1"
)

type TranscoderService struct {
	client     *transcoder.Client
	parentPath string
}

type CreateJobInput struct {
	InputUri  string `json:"input_uri"`
	OutputUri string `json:"output_uri"`
	Preset    string `json:"preset"`
}

type TranscodeServiceConfig struct {
	ParentPath string
}

func NewTranscodeService(config *TranscodeServiceConfig) (*TranscoderService, error) {

	ctx := context.Background()
	client, err := transcoder.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewClient: %v", err)
	}

	return &TranscoderService{client: client, parentPath: config.ParentPath}, nil
}

// createJobFromPreset creates a job based on a given preset template. See
// https://cloud.google.com/transcoder/docs/how-to/jobs#create_jobs_presets
// for more information.
func (ts *TranscoderService) CreateJob(jobInput *CreateJobInput) (*transcoderpb.Job, error) {

	req := &transcoderpb.CreateJobRequest{
		Parent: ts.parentPath,
		Job: &transcoderpb.Job{
			InputUri:  jobInput.InputUri,
			OutputUri: jobInput.OutputUri,
			JobConfig: &transcoderpb.Job_TemplateId{
				TemplateId: jobInput.Preset,
			},
		},
	}
	// Creates the job, Jobs take a variable amount of time to run.
	// You can query for the job state.
	res, err := ts.client.CreateJob(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("createJobFromPreset: %v", err)
	}

	return res, nil
}

func (ts *TranscoderService) GetJob(name string) (*transcoderpb.Job, error) {

	req := &transcoderpb.GetJobRequest{
		Name: name,
	}
	res, err := ts.client.GetJob(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
