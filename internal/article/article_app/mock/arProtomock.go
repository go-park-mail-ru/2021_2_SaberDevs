// Code generated by MockGen. DO NOT EDIT.
// Source: articles.pb.go

// Package mock_article_server is a generated GoMock package.
package mock_article_server

import (
	context "context"
	reflect "reflect"

	article_server "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockArticleDeliveryClient is a mock of ArticleDeliveryClient interface.
type MockArticleDeliveryClient struct {
	ctrl     *gomock.Controller
	recorder *MockArticleDeliveryClientMockRecorder
}

// MockArticleDeliveryClientMockRecorder is the mock recorder for MockArticleDeliveryClient.
type MockArticleDeliveryClientMockRecorder struct {
	mock *MockArticleDeliveryClient
}

// NewMockArticleDeliveryClient creates a new mock instance.
func NewMockArticleDeliveryClient(ctrl *gomock.Controller) *MockArticleDeliveryClient {
	mock := &MockArticleDeliveryClient{ctrl: ctrl}
	mock.recorder = &MockArticleDeliveryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleDeliveryClient) EXPECT() *MockArticleDeliveryClientMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockArticleDeliveryClient) Delete(ctx context.Context, in *article_server.Id, opts ...grpc.CallOption) (*article_server.Nothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*article_server.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockArticleDeliveryClientMockRecorder) Delete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArticleDeliveryClient)(nil).Delete), varargs...)
}

// Fetch mocks base method.
func (m *MockArticleDeliveryClient) Fetch(ctx context.Context, in *article_server.Chunk, opts ...grpc.CallOption) (*article_server.Repview, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Fetch", varargs...)
	ret0, _ := ret[0].(*article_server.Repview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockArticleDeliveryClientMockRecorder) Fetch(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockArticleDeliveryClient)(nil).Fetch), varargs...)
}

// FindArticles mocks base method.
func (m *MockArticleDeliveryClient) FindArticles(ctx context.Context, in *article_server.Queries, opts ...grpc.CallOption) (*article_server.Repview, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindArticles", varargs...)
	ret0, _ := ret[0].(*article_server.Repview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindArticles indicates an expected call of FindArticles.
func (mr *MockArticleDeliveryClientMockRecorder) FindArticles(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindArticles", reflect.TypeOf((*MockArticleDeliveryClient)(nil).FindArticles), varargs...)
}

// FindAuthors mocks base method.
func (m *MockArticleDeliveryClient) FindAuthors(ctx context.Context, in *article_server.Queries, opts ...grpc.CallOption) (*article_server.AView, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAuthors", varargs...)
	ret0, _ := ret[0].(*article_server.AView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAuthors indicates an expected call of FindAuthors.
func (mr *MockArticleDeliveryClientMockRecorder) FindAuthors(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAuthors", reflect.TypeOf((*MockArticleDeliveryClient)(nil).FindAuthors), varargs...)
}

// FindByTag mocks base method.
func (m *MockArticleDeliveryClient) FindByTag(ctx context.Context, in *article_server.Queries, opts ...grpc.CallOption) (*article_server.Repview, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByTag", varargs...)
	ret0, _ := ret[0].(*article_server.Repview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTag indicates an expected call of FindByTag.
func (mr *MockArticleDeliveryClientMockRecorder) FindByTag(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTag", reflect.TypeOf((*MockArticleDeliveryClient)(nil).FindByTag), varargs...)
}

// GetByAuthor mocks base method.
func (m *MockArticleDeliveryClient) GetByAuthor(ctx context.Context, in *article_server.Authors, opts ...grpc.CallOption) (*article_server.Repview, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByAuthor", varargs...)
	ret0, _ := ret[0].(*article_server.Repview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAuthor indicates an expected call of GetByAuthor.
func (mr *MockArticleDeliveryClientMockRecorder) GetByAuthor(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAuthor", reflect.TypeOf((*MockArticleDeliveryClient)(nil).GetByAuthor), varargs...)
}

// GetByCategory mocks base method.
func (m *MockArticleDeliveryClient) GetByCategory(ctx context.Context, in *article_server.Categories, opts ...grpc.CallOption) (*article_server.Repview, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByCategory", varargs...)
	ret0, _ := ret[0].(*article_server.Repview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCategory indicates an expected call of GetByCategory.
func (mr *MockArticleDeliveryClientMockRecorder) GetByCategory(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCategory", reflect.TypeOf((*MockArticleDeliveryClient)(nil).GetByCategory), varargs...)
}

// GetByID mocks base method.
func (m *MockArticleDeliveryClient) GetByID(ctx context.Context, in *article_server.Id, opts ...grpc.CallOption) (*article_server.FullArticle, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByID", varargs...)
	ret0, _ := ret[0].(*article_server.FullArticle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockArticleDeliveryClientMockRecorder) GetByID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockArticleDeliveryClient)(nil).GetByID), varargs...)
}

// GetByTag mocks base method.
func (m *MockArticleDeliveryClient) GetByTag(ctx context.Context, in *article_server.Tags, opts ...grpc.CallOption) (*article_server.Repview, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByTag", varargs...)
	ret0, _ := ret[0].(*article_server.Repview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTag indicates an expected call of GetByTag.
func (mr *MockArticleDeliveryClientMockRecorder) GetByTag(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTag", reflect.TypeOf((*MockArticleDeliveryClient)(nil).GetByTag), varargs...)
}

// Store mocks base method.
func (m *MockArticleDeliveryClient) Store(ctx context.Context, in *article_server.Create, opts ...grpc.CallOption) (*article_server.Created, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Store", varargs...)
	ret0, _ := ret[0].(*article_server.Created)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockArticleDeliveryClientMockRecorder) Store(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockArticleDeliveryClient)(nil).Store), varargs...)
}

// Update mocks base method.
func (m *MockArticleDeliveryClient) Update(ctx context.Context, in *article_server.ArticleUpdate, opts ...grpc.CallOption) (*article_server.Nothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*article_server.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockArticleDeliveryClientMockRecorder) Update(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticleDeliveryClient)(nil).Update), varargs...)
}

// MockArticleDeliveryServer is a mock of ArticleDeliveryServer interface.
type MockArticleDeliveryServer struct {
	ctrl     *gomock.Controller
	recorder *MockArticleDeliveryServerMockRecorder
}

// MockArticleDeliveryServerMockRecorder is the mock recorder for MockArticleDeliveryServer.
type MockArticleDeliveryServerMockRecorder struct {
	mock *MockArticleDeliveryServer
}

// NewMockArticleDeliveryServer creates a new mock instance.
func NewMockArticleDeliveryServer(ctrl *gomock.Controller) *MockArticleDeliveryServer {
	mock := &MockArticleDeliveryServer{ctrl: ctrl}
	mock.recorder = &MockArticleDeliveryServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleDeliveryServer) EXPECT() *MockArticleDeliveryServerMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockArticleDeliveryServer) Delete(arg0 context.Context, arg1 *article_server.Id) (*article_server.Nothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*article_server.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockArticleDeliveryServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArticleDeliveryServer)(nil).Delete), arg0, arg1)
}

// Fetch mocks base method.
func (m *MockArticleDeliveryServer) Fetch(arg0 context.Context, arg1 *article_server.Chunk) (*article_server.Repview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0, arg1)
	ret0, _ := ret[0].(*article_server.Repview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockArticleDeliveryServerMockRecorder) Fetch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockArticleDeliveryServer)(nil).Fetch), arg0, arg1)
}

// FindArticles mocks base method.
func (m *MockArticleDeliveryServer) FindArticles(arg0 context.Context, arg1 *article_server.Queries) (*article_server.Repview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindArticles", arg0, arg1)
	ret0, _ := ret[0].(*article_server.Repview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindArticles indicates an expected call of FindArticles.
func (mr *MockArticleDeliveryServerMockRecorder) FindArticles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindArticles", reflect.TypeOf((*MockArticleDeliveryServer)(nil).FindArticles), arg0, arg1)
}

// FindAuthors mocks base method.
func (m *MockArticleDeliveryServer) FindAuthors(arg0 context.Context, arg1 *article_server.Queries) (*article_server.AView, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAuthors", arg0, arg1)
	ret0, _ := ret[0].(*article_server.AView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAuthors indicates an expected call of FindAuthors.
func (mr *MockArticleDeliveryServerMockRecorder) FindAuthors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAuthors", reflect.TypeOf((*MockArticleDeliveryServer)(nil).FindAuthors), arg0, arg1)
}

// FindByTag mocks base method.
func (m *MockArticleDeliveryServer) FindByTag(arg0 context.Context, arg1 *article_server.Queries) (*article_server.Repview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTag", arg0, arg1)
	ret0, _ := ret[0].(*article_server.Repview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTag indicates an expected call of FindByTag.
func (mr *MockArticleDeliveryServerMockRecorder) FindByTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTag", reflect.TypeOf((*MockArticleDeliveryServer)(nil).FindByTag), arg0, arg1)
}

// GetByAuthor mocks base method.
func (m *MockArticleDeliveryServer) GetByAuthor(arg0 context.Context, arg1 *article_server.Authors) (*article_server.Repview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAuthor", arg0, arg1)
	ret0, _ := ret[0].(*article_server.Repview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAuthor indicates an expected call of GetByAuthor.
func (mr *MockArticleDeliveryServerMockRecorder) GetByAuthor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAuthor", reflect.TypeOf((*MockArticleDeliveryServer)(nil).GetByAuthor), arg0, arg1)
}

// GetByCategory mocks base method.
func (m *MockArticleDeliveryServer) GetByCategory(arg0 context.Context, arg1 *article_server.Categories) (*article_server.Repview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCategory", arg0, arg1)
	ret0, _ := ret[0].(*article_server.Repview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCategory indicates an expected call of GetByCategory.
func (mr *MockArticleDeliveryServerMockRecorder) GetByCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCategory", reflect.TypeOf((*MockArticleDeliveryServer)(nil).GetByCategory), arg0, arg1)
}

// GetByID mocks base method.
func (m *MockArticleDeliveryServer) GetByID(arg0 context.Context, arg1 *article_server.Id) (*article_server.FullArticle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*article_server.FullArticle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockArticleDeliveryServerMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockArticleDeliveryServer)(nil).GetByID), arg0, arg1)
}

// GetByTag mocks base method.
func (m *MockArticleDeliveryServer) GetByTag(arg0 context.Context, arg1 *article_server.Tags) (*article_server.Repview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTag", arg0, arg1)
	ret0, _ := ret[0].(*article_server.Repview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTag indicates an expected call of GetByTag.
func (mr *MockArticleDeliveryServerMockRecorder) GetByTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTag", reflect.TypeOf((*MockArticleDeliveryServer)(nil).GetByTag), arg0, arg1)
}

// Store mocks base method.
func (m *MockArticleDeliveryServer) Store(arg0 context.Context, arg1 *article_server.Create) (*article_server.Created, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(*article_server.Created)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockArticleDeliveryServerMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockArticleDeliveryServer)(nil).Store), arg0, arg1)
}

// Update mocks base method.
func (m *MockArticleDeliveryServer) Update(arg0 context.Context, arg1 *article_server.ArticleUpdate) (*article_server.Nothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*article_server.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockArticleDeliveryServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticleDeliveryServer)(nil).Update), arg0, arg1)
}
