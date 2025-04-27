package taskmanager

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type EmailTasks struct {
	SendEmail TaskLauncher[EmailParams]
}

// Email sending parameters
type EmailParams struct {
	To      string
	From    string
	Subject string
	Content string
}

// Example email task
func SendEmail(ctx context.Context, params EmailParams) (string, error) {
	// Simulate sending email
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(2 * time.Second):
	}
	if params.To == "" {
		return "", errors.New("missing recipient email")
	}
	return fmt.Sprintf("Email sent to %s", params.To), nil
}
