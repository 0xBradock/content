package main

import (
	"context"
	"log"

	"github.com/0xBradock/go-filestorage/domain"
	"github.com/0xBradock/go-filestorage/filestorage"
)

func main() {
	ctx := context.Background()

	storage, err := filestorage.New(ctx, "s3")
	if err != nil {
		log.Fatal("failed to open storage")
	}

	domain.New(ctx, storage)
}
