package storage

import (
	"context"
	"hash/crc32"

	"cloud.google.com/go/storage"
	"golang.org/x/sync/errgroup"
)

func Write(ctx context.Context, client *storage.Client, bucket string, files ...*FileUpload) error {
	var g errgroup.Group

	for _, f := range files {
		f := f

		g.Go(func() error {
			wc := client.Bucket(bucket).Object(f.filename()).NewWriter(ctx)
			wc.CRC32C = crc32.Checksum(f.Data, crc32.MakeTable(crc32.Castagnoli))
			wc.SendCRC32C = true

			if f.MimeType != "" {
				wc.ContentType = f.MimeType
			}

			if _, err := wc.Write(f.Data); err != nil {
				return err
			}

			return wc.Close()
		})
	}

	return g.Wait()
}
