package pubsub

import (
	"context"
	"os"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func NewClient(ctx context.Context, projectID string, options ...option.ClientOption) (*pubsub.Client, error) {
	var opts []option.ClientOption

	if os.Getenv("PUBSUB_EMULATOR_HOST") == "" {
		if credentialsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"); credentialsFile != "" {
			opts = append(opts, option.WithCredentialsFile(credentialsFile))
		}
	}

	opts = append(opts, options...)

	return pubsub.NewClient(ctx, projectID, opts...)
}
