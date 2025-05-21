package domain

import "context"

type Service struct {
	FileStorage FileStorage
}

func New(ctx context.Context, fileStorage FileStorage) *Service {
	return &Service{
		FileStorage: fileStorage,
	}
}
