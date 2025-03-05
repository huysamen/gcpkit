package datastore

import (
	"context"
	"errors"

	"cloud.google.com/go/datastore"
)

type CRUD[E any] interface {
	Client() *datastore.Client
	Exists(ctx context.Context, key *datastore.Key) (exists bool, err error)
	Create(ctx context.Context, key *datastore.Key, entity *E) (*datastore.Key, error)
	CreateTx(tx *datastore.Transaction, key *datastore.Key, entity *E) (*datastore.PendingKey, error)
	CreateBatch(ctx context.Context, keys []*datastore.Key, entities []*E) ([]*datastore.Key, error)
	CreateBatchTx(tx *datastore.Transaction, keys []*datastore.Key, entities []*E) ([]*datastore.PendingKey, error)
	Read(ctx context.Context, key *datastore.Key) (entity *E, err error)
	ReadTx(tx *datastore.Transaction, key *datastore.Key) (entity *E, err error)
	ReadBatch(ctx context.Context, keys []*datastore.Key) (entities []*E, err error)
	ReadBatchTx(tx *datastore.Transaction, keys []*datastore.Key) (entities []*E, err error)
	ListAll(ctx context.Context, kind string, ancestor *datastore.Key) (entities []*E, err error)
	ListAllKeys(ctx context.Context, kind string, ancestor *datastore.Key) (keys []*datastore.Key, err error)
	Count(ctx context.Context, kind string, ancestor *datastore.Key) (count int, err error)
	Update(ctx context.Context, key *datastore.Key, entity *E) error
	UpdateTx(tx *datastore.Transaction, key *datastore.Key, entity *E) error
	UpdateBatch(ctx context.Context, keys []*datastore.Key, entities []*E) error
	UpdateBatchTx(tx *datastore.Transaction, keys []*datastore.Key, entities []*E) error
	Delete(ctx context.Context, key *datastore.Key) error
	DeleteTx(tx *datastore.Transaction, key *datastore.Key) error
	DeleteBatch(ctx context.Context, keys []*datastore.Key) error
	DeleteBatchTx(tx *datastore.Transaction, keys []*datastore.Key) error
}

type CRUDRepo[E any] struct {
	client *datastore.Client
}

func NewCRUDRepo[E any](client *datastore.Client) CRUD[E] {
	return &CRUDRepo[E]{client: client}
}

func (r *CRUDRepo[E]) Client() *datastore.Client {
	return r.client
}

func (r *CRUDRepo[E]) Exists(ctx context.Context, key *datastore.Key) (exists bool, err error) {
	return Exists(ctx, r.client, key)
}

func (r *CRUDRepo[E]) Create(ctx context.Context, key *datastore.Key, entity *E) (*datastore.Key, error) {
	return r.client.Put(ctx, key, entity)
}

func (r *CRUDRepo[E]) CreateTx(tx *datastore.Transaction, key *datastore.Key, entity *E) (*datastore.PendingKey, error) {
	return tx.Put(key, entity)
}

func (r *CRUDRepo[E]) CreateBatch(ctx context.Context, keys []*datastore.Key, entities []*E) ([]*datastore.Key, error) {
	return r.client.PutMulti(ctx, keys, entities)
}

func (r *CRUDRepo[E]) CreateBatchTx(tx *datastore.Transaction, keys []*datastore.Key, entities []*E) ([]*datastore.PendingKey, error) {
	return tx.PutMulti(keys, entities)
}

func (r *CRUDRepo[E]) Read(ctx context.Context, key *datastore.Key) (entity *E, err error) {
	entity = new(E)
	err = r.client.Get(ctx, key, entity)

	if errors.Is(err, datastore.ErrNoSuchEntity) {
		return nil, nil
	}

	return
}

func (r *CRUDRepo[E]) ReadTx(tx *datastore.Transaction, key *datastore.Key) (entity *E, err error) {
	entity = new(E)
	err = tx.Get(key, entity)

	if errors.Is(err, datastore.ErrNoSuchEntity) {
		return nil, nil
	}

	return
}

func (r *CRUDRepo[E]) ReadBatch(ctx context.Context, keys []*datastore.Key) (entities []*E, err error) {
	entities = make([]*E, len(keys))
	err = r.client.GetMulti(ctx, keys, entities)

	return
}

func (r *CRUDRepo[E]) ReadBatchTx(tx *datastore.Transaction, keys []*datastore.Key) (entities []*E, err error) {
	entities = make([]*E, len(keys))
	err = tx.GetMulti(keys, entities)

	return
}

func (r *CRUDRepo[E]) ListAll(ctx context.Context, kind string, ancestor *datastore.Key) (entities []*E, err error) {
	query := datastore.NewQuery(kind)

	if ancestor != nil {
		query = query.Ancestor(ancestor)
	}

	return Query[E](ctx, r.client, query)
}

func (r *CRUDRepo[E]) ListAllKeys(ctx context.Context, kind string, ancestor *datastore.Key) (keys []*datastore.Key, err error) {
	query := datastore.NewQuery(kind)

	if ancestor != nil {
		query = query.Ancestor(ancestor)
	}

	keys, _, err = QueryKeys(ctx, r.client, query)

	return keys, err
}

func (r *CRUDRepo[E]) Count(ctx context.Context, kind string, ancestor *datastore.Key) (int, error) {
	query := datastore.NewQuery(kind)

	if ancestor != nil {
		query = query.Ancestor(ancestor)
	}

	return CountForQuery(ctx, r.client, query)
}

func (r *CRUDRepo[E]) Update(ctx context.Context, key *datastore.Key, entity *E) error {
	_, err := r.client.Put(ctx, key, entity)

	return err
}

func (r *CRUDRepo[E]) UpdateTx(tx *datastore.Transaction, key *datastore.Key, entity *E) error {
	_, err := tx.Put(key, entity)

	return err
}

func (r *CRUDRepo[E]) UpdateBatch(ctx context.Context, keys []*datastore.Key, entities []*E) error {
	_, err := r.client.PutMulti(ctx, keys, entities)

	return err
}

func (r *CRUDRepo[E]) UpdateBatchTx(tx *datastore.Transaction, keys []*datastore.Key, entities []*E) error {
	_, err := tx.PutMulti(keys, entities)

	return err
}

func (r *CRUDRepo[E]) Delete(ctx context.Context, key *datastore.Key) error {
	return r.client.Delete(ctx, key)
}

func (r *CRUDRepo[E]) DeleteTx(tx *datastore.Transaction, key *datastore.Key) error {
	return tx.Delete(key)
}

func (r *CRUDRepo[E]) DeleteBatch(ctx context.Context, keys []*datastore.Key) error {
	return r.client.DeleteMulti(ctx, keys)
}

func (r *CRUDRepo[E]) DeleteBatchTx(tx *datastore.Transaction, keys []*datastore.Key) error {
	return tx.DeleteMulti(keys)
}
