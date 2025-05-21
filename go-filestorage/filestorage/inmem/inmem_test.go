package storage_inmem_test

import (
	"context"
	"testing"

	"github.com/0xBradock/go-filestorage/filestorage"
	"github.com/stretchr/testify/require"
)

func TestInmem_UploadDownload(t *testing.T) {
	ctx := context.TODO()
	storage, err := filestorage.New(ctx, "inmem")
	require.NoError(t, err)

	fileName := "filename.json"
	filePath := "test/file/path"
	content := "content"
	err = storage.Upload(ctx, fileName, filePath, []byte(content))
	require.NoError(t, err)

	res, err := storage.Download(ctx, fileName, filePath)
	require.NoError(t, err)
	require.Equal(t, content, string(res))
}
