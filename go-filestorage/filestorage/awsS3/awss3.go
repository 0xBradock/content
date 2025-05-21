package awsS3

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/0xBradock/go-filestorage/domain"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct {
	client *s3.Client
	bucket string
}

func newClient(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     "AccessKeyID",
				SecretAccessKey: "SecretAccessKey",
			},
		}))
	if err != nil {
		return nil, errors.New("failed to create s3 client")
	}

	client := s3.NewFromConfig(cfg)

	return client, nil
}

// Download implements domain.FileStorage.
func (s *S3) Download(ctx context.Context, fileName string, filePath string) ([]byte, error) {
	key := fmt.Sprintf("%s/%s", filePath, fileName)

	object, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	defer object.Body.Close()

	return io.ReadAll(object.Body)
}

// Upload implements domain.FileStorage.
func (s *S3) Upload(ctx context.Context, fileName string, filePath string, content []byte) error {
	key := fmt.Sprintf("%s/%s", filePath, fileName)

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
		Body:   bytes.NewReader(content),
	})
	if err != nil {
		return err
	}

	return nil
}

// New returns an implementation for the domain.FileStorage with a S3 as storage.
func New(ctx context.Context) (domain.FileStorage, error) {
	client, err := newClient(ctx)
	if err != nil {
		return nil, err
	}

	return &S3{
		client: client,
		bucket: "bradock-erase",
	}, nil
}
