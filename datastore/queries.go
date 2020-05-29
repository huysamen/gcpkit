package datastore

import (
	"context"
	"errors"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

func KeysQuery(ctx context.Context, client *datastore.Client, query *datastore.Query) ([]*datastore.Key, *datastore.Cursor, error) {
	it := client.Run(ctx, query)
	keys := make([]*datastore.Key, 0)

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

	cursor, err := it.Cursor()
	if err != nil {
		return nil, nil, err
	}

	return keys, &cursor, nil
}

func ListQuery(
	ctx context.Context,
	client *datastore.Client,
	query *datastore.Query,
	generator func() interface{},
) ([]interface{}, *datastore.Cursor, error) {
	it := client.Run(ctx, query)

	entities := make([]interface{}, 0)

	for {
		entity := generator()

		_, err := it.Next(entity)
		if errors.Is(err, iterator.Done) {
			break
		}

		if _, ok := err.(*datastore.ErrFieldMismatch); ok {
			entities = append(entities, entity)
			continue
		}

		if err != nil {
			return nil, nil, err
		}

		entities = append(entities, entity)
	}

	cursor, err := it.Cursor()
	if err != nil {
		return nil, nil, err
	}

	return entities, &cursor, nil
}

func Exists(ctx context.Context, client *datastore.Client, namespace string, kind string, key *datastore.Key) (bool, error) {
	query := datastore.
		NewQuery(kind).
		KeysOnly().
		Filter("__key__ =", key).
		Limit(1)

	if namespace != "" {
		query = query.Namespace(namespace)
	}

	keys, _, err := KeysQuery(ctx, client, query)

	if err != nil {
		return false, err
	}

	return len(keys) > 0, nil
}
