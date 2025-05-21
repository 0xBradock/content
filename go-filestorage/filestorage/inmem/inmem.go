package storage_inmem

import (
	"context"
	"fmt"

	"github.com/0xBradock/go-filestorage/domain"
)

type Inmem struct {
	store map[string][]byte
}

// Download implements domain.FileStorage.
func (i *Inmem) Download(ctx context.Context, fileName string, filePath string) ([]byte, error) {
	key := fmt.Sprintf("%s/%s", filePath, fileName)
	val, exist := i.store[key]
	if !exist {
		return []byte{}, domain.ErrFileNotFound
	}

	return val, nil
}

// Upload implements domain.FileStorage.
func (i *Inmem) Upload(ctx context.Context, fileName string, filePath string, content []byte) error {
	key := fmt.Sprintf("%s/%s", filePath, fileName)
	i.store[key] = content

	return nil
}

func New(ctx context.Context) (domain.FileStorage, error) {
	return &Inmem{store: map[string][]byte{}}, nil
}
