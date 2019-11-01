package storage

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func NewClient(ctx context.Context, options ...option.ClientOption) (*storage.Client, error) {
	var opts []option.ClientOption

	if credentialsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"); credentialsFile != "" {
		opts = append(opts, option.WithCredentialsFile(credentialsFile))
	}

	opts = append(opts, options...)

	return storage.NewClient(ctx, opts...)
}
