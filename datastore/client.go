package datastore

import (
	"context"
	"os"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"
)

func NewClient(ctx context.Context, projectID string, options ...option.ClientOption) (*datastore.Client, error) {
	var opts []option.ClientOption

	if os.Getenv("DATASTORE_EMULATOR_HOST") == "" {
		if credentialsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"); credentialsFile != "" {
			opts = append(opts, option.WithCredentialsFile(credentialsFile))
		}
	}

	opts = append(opts, options...)

	return datastore.NewClient(ctx, projectID, opts...)
}
