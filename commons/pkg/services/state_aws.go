package services

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// AwsStateStorer implements the StateStorer interface
type AwsStateStorer struct {
	Address  string
	Filename string
}

// NewAwsStateStorer returns a StateStorer implementation for AWS
func NewAwsStateStorer(address, filename string) StateStorer {
	return &AwsStateStorer{
		Address:  address,
		Filename: filename,
	}
}

// GetState retrieves the state from the S3 bucket
func (as *AwsStateStorer) GetState() (State, error) {
	sess, err := NewAwsSession()
	if err != nil {
		return nil, err
	}
	downloader := s3manager.NewDownloader(sess)
	buffer := &aws.WriteAtBuffer{}
	_, err = downloader.Download(buffer,
		&s3.GetObjectInput{
			Bucket: aws.String(as.Address),
			Key:    aws.String(as.Filename),
		})
	if err != nil {
		return nil, err
	}
	state := make(map[string]int64)
	err = json.Unmarshal(buffer.Bytes(), &state)
	if err != nil {
		return nil, err
	}
	return state, nil
}

// SetState writes the state to the S3 bucket
func (as *AwsStateStorer) SetState(state State) error {
	json, err := json.Marshal(state)
	if err != nil {
		return err
	}
	sess, err := NewAwsSession()
	if err != nil {
		return err
	}
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(as.Address),
		Key:    aws.String(as.Filename),
		Body:   strings.NewReader(string(json)),
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteState writes the state to the S3 bucket
func (as *AwsStateStorer) DeleteState() error {
	sess, err := NewAwsSession()
	if err != nil {
		return err
	}
	s3service := s3.New(sess)
	_, err = s3service.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(as.Address),
		Key:    aws.String(as.Filename)})
	if err != nil {
		return err
	}
	err = s3service.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(as.Address),
		Key:    aws.String(as.Filename),
	})
	if err != nil {
		return err
	}
	return nil
}
