package taskmanager

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

// XML to JSON conversion parameters
type XMLToJSONParams struct {
	Content string
	Bucket  string
	Key     string
}

// Example XML to JSON task
func ConvertXMLToJSONTask(ctx context.Context, params XMLToJSONParams) (string, error) {
	var data map[string]interface{}
	if err := xml.Unmarshal([]byte(params.Content), &data); err != nil {
		return "", fmt.Errorf("xml unmarshal error: %w", err)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("json marshal error: %w", err)
	}

	// Simulate S3 upload
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(3 * time.Second):
	}
	return fmt.Sprintf("Content: %s\nStored to s3://%s/%s", jsonData, params.Bucket, params.Key), nil
}
