package domain

import (
	"context"
	"errors"
)

var ErrFileNotFound = errors.New("file not found")

// FileStorage defines the interface to handle storage of files in different locations.
type FileStorage interface {
	// Upload stores the contents of a file in a given location and returns an error.
	Upload(ctx context.Context, fileName, filePath string, content []byte) error

	// Download retrieves the contents of a file if it exists or returns an error.
	Download(ctx context.Context, fileName, filePath string) ([]byte, error)
}
