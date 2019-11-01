package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"io/ioutil"

	"golang.org/x/sync/errgroup"
)

func Read(ctx context.Context, client *storage.Client, bucket string, files ...*FileDownload) error {
	var g errgroup.Group

	for _, f := range files {
		f := f

		g.Go(func() error {
			rc, err := client.Bucket(bucket).Object(f.filename()).NewReader(ctx)
			if err != nil {
				return err
			}

			defer func() { _ = rc.Close() }()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return err
			}

			f.Data = data

			return nil
		})
	}

	return g.Wait()
}
