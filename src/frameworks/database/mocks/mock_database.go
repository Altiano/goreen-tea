// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	database "gitlab.com/altiano/golang-boilerplate/src/frameworks/database"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

// MockDb is a mock of Db interface.
type MockDb struct {
	ctrl     *gomock.Controller
	recorder *MockDbMockRecorder
}

// MockDbMockRecorder is the mock recorder for MockDb.
type MockDbMockRecorder struct {
	mock *MockDb
}

// NewMockDb creates a new mock instance.
func NewMockDb(ctrl *gomock.Controller) *MockDb {
	mock := &MockDb{ctrl: ctrl}
	mock.recorder = &MockDbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDb) EXPECT() *MockDbMockRecorder {
	return m.recorder
}

// Collection mocks base method.
func (m *MockDb) Collection(name string, opts ...*options.CollectionOptions) database.Coll {
	m.ctrl.T.Helper()
	varargs := []interface{}{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Collection", varargs...)
	ret0, _ := ret[0].(database.Coll)
	return ret0
}

// Collection indicates an expected call of Collection.
func (mr *MockDbMockRecorder) Collection(name interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collection", reflect.TypeOf((*MockDb)(nil).Collection), varargs...)
}

// MockColl is a mock of Coll interface.
type MockColl struct {
	ctrl     *gomock.Controller
	recorder *MockCollMockRecorder
}

// MockCollMockRecorder is the mock recorder for MockColl.
type MockCollMockRecorder struct {
	mock *MockColl
}

// NewMockColl creates a new mock instance.
func NewMockColl(ctrl *gomock.Controller) *MockColl {
	mock := &MockColl{ctrl: ctrl}
	mock.recorder = &MockCollMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockColl) EXPECT() *MockCollMockRecorder {
	return m.recorder
}

// Aggregate mocks base method.
func (m *MockColl) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, pipeline}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Aggregate", varargs...)
	ret0, _ := ret[0].(*mongo.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Aggregate indicates an expected call of Aggregate.
func (mr *MockCollMockRecorder) Aggregate(ctx, pipeline interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, pipeline}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aggregate", reflect.TypeOf((*MockColl)(nil).Aggregate), varargs...)
}

// BulkWrite mocks base method.
func (m *MockColl) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, models}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BulkWrite", varargs...)
	ret0, _ := ret[0].(*mongo.BulkWriteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkWrite indicates an expected call of BulkWrite.
func (mr *MockCollMockRecorder) BulkWrite(ctx, models interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, models}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkWrite", reflect.TypeOf((*MockColl)(nil).BulkWrite), varargs...)
}

// CountDocuments mocks base method.
func (m *MockColl) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountDocuments", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountDocuments indicates an expected call of CountDocuments.
func (mr *MockCollMockRecorder) CountDocuments(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountDocuments", reflect.TypeOf((*MockColl)(nil).CountDocuments), varargs...)
}

// DeleteMany mocks base method.
func (m *MockColl) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMany", varargs...)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMany indicates an expected call of DeleteMany.
func (mr *MockCollMockRecorder) DeleteMany(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMany", reflect.TypeOf((*MockColl)(nil).DeleteMany), varargs...)
}

// DeleteOne mocks base method.
func (m *MockColl) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteOne", varargs...)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockCollMockRecorder) DeleteOne(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockColl)(nil).DeleteOne), varargs...)
}

// Drop mocks base method.
func (m *MockColl) Drop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Drop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Drop indicates an expected call of Drop.
func (mr *MockCollMockRecorder) Drop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drop", reflect.TypeOf((*MockColl)(nil).Drop), ctx)
}

// Find mocks base method.
func (m *MockColl) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (database.Cursor, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(database.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockCollMockRecorder) Find(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockColl)(nil).Find), varargs...)
}

// FindByID mocks base method.
func (m *MockColl) FindByID(ctx context.Context, id, v interface{}, opts ...*options.FindOneOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id, v}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByID", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindByID indicates an expected call of FindByID.
func (mr *MockCollMockRecorder) FindByID(ctx, id, v interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id, v}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockColl)(nil).FindByID), varargs...)
}

// FindOne mocks base method.
func (m *MockColl) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) database.SingleResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(database.SingleResult)
	return ret0
}

// FindOne indicates an expected call of FindOne.
func (mr *MockCollMockRecorder) FindOne(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockColl)(nil).FindOne), varargs...)
}

// FindOneAndDelete mocks base method.
func (m *MockColl) FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) database.SingleResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOneAndDelete", varargs...)
	ret0, _ := ret[0].(database.SingleResult)
	return ret0
}

// FindOneAndDelete indicates an expected call of FindOneAndDelete.
func (mr *MockCollMockRecorder) FindOneAndDelete(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneAndDelete", reflect.TypeOf((*MockColl)(nil).FindOneAndDelete), varargs...)
}

// FindOneAndUpdate mocks base method.
func (m *MockColl) FindOneAndUpdate(ctx context.Context, filter, update interface{}, opts ...*options.FindOneAndUpdateOptions) database.SingleResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter, update}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOneAndUpdate", varargs...)
	ret0, _ := ret[0].(database.SingleResult)
	return ret0
}

// FindOneAndUpdate indicates an expected call of FindOneAndUpdate.
func (mr *MockCollMockRecorder) FindOneAndUpdate(ctx, filter, update interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter, update}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneAndUpdate", reflect.TypeOf((*MockColl)(nil).FindOneAndUpdate), varargs...)
}

// InsertMany mocks base method.
func (m *MockColl) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, documents}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertMany", varargs...)
	ret0, _ := ret[0].(*mongo.InsertManyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMany indicates an expected call of InsertMany.
func (mr *MockCollMockRecorder) InsertMany(ctx, documents interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, documents}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMany", reflect.TypeOf((*MockColl)(nil).InsertMany), varargs...)
}

// InsertOne mocks base method.
func (m *MockColl) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, document}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertOne", varargs...)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockCollMockRecorder) InsertOne(ctx, document interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, document}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockColl)(nil).InsertOne), varargs...)
}

// UpdateMany mocks base method.
func (m *MockColl) UpdateMany(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter, update}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMany", varargs...)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMany indicates an expected call of UpdateMany.
func (mr *MockCollMockRecorder) UpdateMany(ctx, filter, update interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter, update}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMany", reflect.TypeOf((*MockColl)(nil).UpdateMany), varargs...)
}

// UpdateOne mocks base method.
func (m *MockColl) UpdateOne(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter, update}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOne", varargs...)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOne indicates an expected call of UpdateOne.
func (mr *MockCollMockRecorder) UpdateOne(ctx, filter, update interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter, update}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockColl)(nil).UpdateOne), varargs...)
}

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockClient) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockClientMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockClient)(nil).Connect))
}

// Database mocks base method.
func (m *MockClient) Database(arg0 string) database.Db {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Database", arg0)
	ret0, _ := ret[0].(database.Db)
	return ret0
}

// Database indicates an expected call of Database.
func (mr *MockClientMockRecorder) Database(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Database", reflect.TypeOf((*MockClient)(nil).Database), arg0)
}

// StartSession mocks base method.
func (m *MockClient) StartSession() (mongo.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartSession")
	ret0, _ := ret[0].(mongo.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartSession indicates an expected call of StartSession.
func (mr *MockClientMockRecorder) StartSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSession", reflect.TypeOf((*MockClient)(nil).StartSession))
}

// MockSingleResult is a mock of SingleResult interface.
type MockSingleResult struct {
	ctrl     *gomock.Controller
	recorder *MockSingleResultMockRecorder
}

// MockSingleResultMockRecorder is the mock recorder for MockSingleResult.
type MockSingleResultMockRecorder struct {
	mock *MockSingleResult
}

// NewMockSingleResult creates a new mock instance.
func NewMockSingleResult(ctrl *gomock.Controller) *MockSingleResult {
	mock := &MockSingleResult{ctrl: ctrl}
	mock.recorder = &MockSingleResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSingleResult) EXPECT() *MockSingleResultMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockSingleResult) Decode(v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode.
func (mr *MockSingleResultMockRecorder) Decode(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockSingleResult)(nil).Decode), v)
}

// Err mocks base method.
func (m *MockSingleResult) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockSingleResultMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockSingleResult)(nil).Err))
}

// MockCursor is a mock of Cursor interface.
type MockCursor struct {
	ctrl     *gomock.Controller
	recorder *MockCursorMockRecorder
}

// MockCursorMockRecorder is the mock recorder for MockCursor.
type MockCursorMockRecorder struct {
	mock *MockCursor
}

// NewMockCursor creates a new mock instance.
func NewMockCursor(ctrl *gomock.Controller) *MockCursor {
	mock := &MockCursor{ctrl: ctrl}
	mock.recorder = &MockCursorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCursor) EXPECT() *MockCursorMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockCursor) All(ctx context.Context, results interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", ctx, results)
	ret0, _ := ret[0].(error)
	return ret0
}

// All indicates an expected call of All.
func (mr *MockCursorMockRecorder) All(ctx, results interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockCursor)(nil).All), ctx, results)
}