package filestorage

import (
	"context"
	"fmt"

	"github.com/0xBradock/go-filestorage/domain"
	"github.com/0xBradock/go-filestorage/filestorage/awsS3"
	storage_inmem "github.com/0xBradock/go-filestorage/filestorage/inmem"
)

func New(ctx context.Context, storage string) (domain.FileStorage, error) {
	switch storage {
	case "s3":
		return awsS3.New(ctx)
	case "inmem":
		return storage_inmem.New(ctx)
	}

	return nil, fmt.Errorf("file storage %s does not exist", storage)
}
