// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/huysamen/gcpkit/datastore (interfaces: CRUD)
//
// Generated by this command:
//
//	mockgen -destination datastore/mock/crud.go -package mock github.com/huysamen/gcpkit/datastore CRUD
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	datastore "cloud.google.com/go/datastore"
	gomock "go.uber.org/mock/gomock"
)

// MockCRUD is a mock of CRUD interface.
type MockCRUD[E any] struct {
	ctrl     *gomock.Controller
	recorder *MockCRUDMockRecorder[E]
	isgomock struct{}
}

// MockCRUDMockRecorder is the mock recorder for MockCRUD.
type MockCRUDMockRecorder[E any] struct {
	mock *MockCRUD[E]
}

// NewMockCRUD creates a new mock instance.
func NewMockCRUD[E any](ctrl *gomock.Controller) *MockCRUD[E] {
	mock := &MockCRUD[E]{ctrl: ctrl}
	mock.recorder = &MockCRUDMockRecorder[E]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCRUD[E]) EXPECT() *MockCRUDMockRecorder[E] {
	return m.recorder
}

// Client mocks base method.
func (m *MockCRUD[E]) Client() *datastore.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(*datastore.Client)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockCRUDMockRecorder[E]) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockCRUD[E])(nil).Client))
}

// Count mocks base method.
func (m *MockCRUD[E]) Count(ctx context.Context, kind string, ancestor *datastore.Key) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, kind, ancestor)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockCRUDMockRecorder[E]) Count(ctx, kind, ancestor any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockCRUD[E])(nil).Count), ctx, kind, ancestor)
}

// Create mocks base method.
func (m *MockCRUD[E]) Create(ctx context.Context, key *datastore.Key, entity *E) (*datastore.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, key, entity)
	ret0, _ := ret[0].(*datastore.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCRUDMockRecorder[E]) Create(ctx, key, entity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCRUD[E])(nil).Create), ctx, key, entity)
}

// CreateBatch mocks base method.
func (m *MockCRUD[E]) CreateBatch(ctx context.Context, keys []*datastore.Key, entities []*E) ([]*datastore.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBatch", ctx, keys, entities)
	ret0, _ := ret[0].([]*datastore.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBatch indicates an expected call of CreateBatch.
func (mr *MockCRUDMockRecorder[E]) CreateBatch(ctx, keys, entities any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBatch", reflect.TypeOf((*MockCRUD[E])(nil).CreateBatch), ctx, keys, entities)
}

// CreateBatchTx mocks base method.
func (m *MockCRUD[E]) CreateBatchTx(tx *datastore.Transaction, keys []*datastore.Key, entities []*E) ([]*datastore.PendingKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBatchTx", tx, keys, entities)
	ret0, _ := ret[0].([]*datastore.PendingKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBatchTx indicates an expected call of CreateBatchTx.
func (mr *MockCRUDMockRecorder[E]) CreateBatchTx(tx, keys, entities any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBatchTx", reflect.TypeOf((*MockCRUD[E])(nil).CreateBatchTx), tx, keys, entities)
}

// CreateTx mocks base method.
func (m *MockCRUD[E]) CreateTx(tx *datastore.Transaction, key *datastore.Key, entity *E) (*datastore.PendingKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTx", tx, key, entity)
	ret0, _ := ret[0].(*datastore.PendingKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockCRUDMockRecorder[E]) CreateTx(tx, key, entity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockCRUD[E])(nil).CreateTx), tx, key, entity)
}

// Delete mocks base method.
func (m *MockCRUD[E]) Delete(ctx context.Context, key *datastore.Key) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCRUDMockRecorder[E]) Delete(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCRUD[E])(nil).Delete), ctx, key)
}

// DeleteBatch mocks base method.
func (m *MockCRUD[E]) DeleteBatch(ctx context.Context, keys []*datastore.Key) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBatch", ctx, keys)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBatch indicates an expected call of DeleteBatch.
func (mr *MockCRUDMockRecorder[E]) DeleteBatch(ctx, keys any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBatch", reflect.TypeOf((*MockCRUD[E])(nil).DeleteBatch), ctx, keys)
}

// DeleteBatchTx mocks base method.
func (m *MockCRUD[E]) DeleteBatchTx(tx *datastore.Transaction, keys []*datastore.Key) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBatchTx", tx, keys)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBatchTx indicates an expected call of DeleteBatchTx.
func (mr *MockCRUDMockRecorder[E]) DeleteBatchTx(tx, keys any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBatchTx", reflect.TypeOf((*MockCRUD[E])(nil).DeleteBatchTx), tx, keys)
}

// DeleteTx mocks base method.
func (m *MockCRUD[E]) DeleteTx(tx *datastore.Transaction, key *datastore.Key) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTx", tx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTx indicates an expected call of DeleteTx.
func (mr *MockCRUDMockRecorder[E]) DeleteTx(tx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTx", reflect.TypeOf((*MockCRUD[E])(nil).DeleteTx), tx, key)
}

// Exists mocks base method.
func (m *MockCRUD[E]) Exists(ctx context.Context, key *datastore.Key) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, key)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockCRUDMockRecorder[E]) Exists(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockCRUD[E])(nil).Exists), ctx, key)
}

// ListAll mocks base method.
func (m *MockCRUD[E]) ListAll(ctx context.Context, kind string, ancestor *datastore.Key) ([]*E, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", ctx, kind, ancestor)
	ret0, _ := ret[0].([]*E)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockCRUDMockRecorder[E]) ListAll(ctx, kind, ancestor any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockCRUD[E])(nil).ListAll), ctx, kind, ancestor)
}

// ListAllKeys mocks base method.
func (m *MockCRUD[E]) ListAllKeys(ctx context.Context, kind string, ancestor *datastore.Key) ([]*datastore.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllKeys", ctx, kind, ancestor)
	ret0, _ := ret[0].([]*datastore.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllKeys indicates an expected call of ListAllKeys.
func (mr *MockCRUDMockRecorder[E]) ListAllKeys(ctx, kind, ancestor any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllKeys", reflect.TypeOf((*MockCRUD[E])(nil).ListAllKeys), ctx, kind, ancestor)
}

// Read mocks base method.
func (m *MockCRUD[E]) Read(ctx context.Context, key *datastore.Key) (*E, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, key)
	ret0, _ := ret[0].(*E)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockCRUDMockRecorder[E]) Read(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockCRUD[E])(nil).Read), ctx, key)
}

// ReadBatch mocks base method.
func (m *MockCRUD[E]) ReadBatch(ctx context.Context, keys []*datastore.Key) ([]*E, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadBatch", ctx, keys)
	ret0, _ := ret[0].([]*E)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadBatch indicates an expected call of ReadBatch.
func (mr *MockCRUDMockRecorder[E]) ReadBatch(ctx, keys any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadBatch", reflect.TypeOf((*MockCRUD[E])(nil).ReadBatch), ctx, keys)
}

// ReadBatchTx mocks base method.
func (m *MockCRUD[E]) ReadBatchTx(tx *datastore.Transaction, keys []*datastore.Key) ([]*E, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadBatchTx", tx, keys)
	ret0, _ := ret[0].([]*E)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadBatchTx indicates an expected call of ReadBatchTx.
func (mr *MockCRUDMockRecorder[E]) ReadBatchTx(tx, keys any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadBatchTx", reflect.TypeOf((*MockCRUD[E])(nil).ReadBatchTx), tx, keys)
}

// ReadTx mocks base method.
func (m *MockCRUD[E]) ReadTx(tx *datastore.Transaction, key *datastore.Key) (*E, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadTx", tx, key)
	ret0, _ := ret[0].(*E)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadTx indicates an expected call of ReadTx.
func (mr *MockCRUDMockRecorder[E]) ReadTx(tx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadTx", reflect.TypeOf((*MockCRUD[E])(nil).ReadTx), tx, key)
}

// Update mocks base method.
func (m *MockCRUD[E]) Update(ctx context.Context, key *datastore.Key, entity *E) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, key, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCRUDMockRecorder[E]) Update(ctx, key, entity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCRUD[E])(nil).Update), ctx, key, entity)
}

// UpdateBatch mocks base method.
func (m *MockCRUD[E]) UpdateBatch(ctx context.Context, keys []*datastore.Key, entities []*E) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBatch", ctx, keys, entities)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBatch indicates an expected call of UpdateBatch.
func (mr *MockCRUDMockRecorder[E]) UpdateBatch(ctx, keys, entities any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBatch", reflect.TypeOf((*MockCRUD[E])(nil).UpdateBatch), ctx, keys, entities)
}

// UpdateBatchTx mocks base method.
func (m *MockCRUD[E]) UpdateBatchTx(tx *datastore.Transaction, keys []*datastore.Key, entities []*E) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBatchTx", tx, keys, entities)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBatchTx indicates an expected call of UpdateBatchTx.
func (mr *MockCRUDMockRecorder[E]) UpdateBatchTx(tx, keys, entities any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBatchTx", reflect.TypeOf((*MockCRUD[E])(nil).UpdateBatchTx), tx, keys, entities)
}

// UpdateTx mocks base method.
func (m *MockCRUD[E]) UpdateTx(tx *datastore.Transaction, key *datastore.Key, entity *E) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTx", tx, key, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTx indicates an expected call of UpdateTx.
func (mr *MockCRUDMockRecorder[E]) UpdateTx(tx, key, entity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTx", reflect.TypeOf((*MockCRUD[E])(nil).UpdateTx), tx, key, entity)
}
