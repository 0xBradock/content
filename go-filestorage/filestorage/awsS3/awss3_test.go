package awsS3_test

import (
	"context"
	"os"
	"testing"

	"github.com/0xBradock/go-filestorage/filestorage"
	"github.com/stretchr/testify/require"
)

func TestAWSS3_UploadDownload(t *testing.T) {
	ctx := context.TODO()
	os.Setenv("AWS_REGION", "eu-west-1")

	storage, err := filestorage.New(ctx, "s3")
	require.NoError(t, err)

	fileName := "a-file-name"
	filePath := "testdata/some/path"
	content := "content"

	err = storage.Upload(ctx, fileName, filePath, []byte(content))
	require.NoError(t, err)

	res, err := storage.Download(ctx, fileName, filePath)
	require.NoError(t, err)
	require.Equal(t, content, string(res))
}
