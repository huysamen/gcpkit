package datastore

import (
	"context"
	"errors"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

func Exists(ctx context.Context, client *datastore.Client, key *datastore.Key) (exists bool, err error) {
	query := datastore.NewQuery(key.Kind).
		FilterField("__key__", "=", key).
		KeysOnly().
		Limit(1)

	count, err := CountForQuery(ctx, client, query)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func Query[E any](ctx context.Context, client *datastore.Client, query *datastore.Query) (entities []*E, err error) {
	it := client.Run(ctx, query)
	entities = make([]*E, 0)

	for {
		e := new(E)

		_, err = it.Next(e)
		if errors.Is(err, iterator.Done) {
			err = nil

			break
		}

		var errFieldMismatch *datastore.ErrFieldMismatch

		if errors.As(err, &errFieldMismatch) {
			entities = append(entities, e)

			continue
		}

		if err != nil {
			return nil, err
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func QueryOne[E any](ctx context.Context, client *datastore.Client, query *datastore.Query) (entity *E, err error) {
	entities, err := Query[E](ctx, client, query)
	if err != nil {
		return nil, err
	}

	if len(entities) == 0 {
		return nil, nil
	}

	return entities[0], nil
}

func QueryKeys(ctx context.Context, client *datastore.Client, query *datastore.Query) (keys []*datastore.Key, cursor *datastore.Cursor, err error) {
	query = query.KeysOnly()
	it := client.Run(ctx, query)
	keys = make([]*datastore.Key, 0)

	for {
		key, err := it.Next(nil)
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return nil, nil, err
		}

		keys = append(keys, key)
	}

	c, err := it.Cursor()
	if err != nil {
		return nil, nil, err
	}

	cursor = &c

	return keys, cursor, nil
}

func CountForQuery(ctx context.Context, client *datastore.Client, query *datastore.Query) (count int, err error) {
	keys, _, err := QueryKeys(ctx, client, query)
	if err != nil {
		return
	}

	count = len(keys)

	return
}
