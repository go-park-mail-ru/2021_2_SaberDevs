// Code generated by MockGen. DO NOT EDIT.
// Source: articles.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	models "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockArticleUsecase is a mock of ArticleUsecase interface
type MockArticleUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockArticleUsecaseMockRecorder
}

// MockArticleUsecaseMockRecorder is the mock recorder for MockArticleUsecase
type MockArticleUsecaseMockRecorder struct {
	mock *MockArticleUsecase
}

// NewMockArticleUsecase creates a new mock instance
func NewMockArticleUsecase(ctrl *gomock.Controller) *MockArticleUsecase {
	mock := &MockArticleUsecase{ctrl: ctrl}
	mock.recorder = &MockArticleUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockArticleUsecase) EXPECT() *MockArticleUsecaseMockRecorder {
	return m.recorder
}

// Fetch mocks base method
func (m *MockArticleUsecase) Fetch(ctx context.Context, idLastLoaded string, chunkSize int) ([]models.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, idLastLoaded, chunkSize)
	ret0, _ := ret[0].([]models.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockArticleUsecaseMockRecorder) Fetch(ctx, idLastLoaded, chunkSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockArticleUsecase)(nil).Fetch), ctx, idLastLoaded, chunkSize)
}

// GetByID mocks base method
func (m *MockArticleUsecase) GetByID(ctx context.Context, id int64) (models.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(models.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockArticleUsecaseMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockArticleUsecase)(nil).GetByID), ctx, id)
}

// GetByTag mocks base method
func (m *MockArticleUsecase) GetByTag(ctx context.Context, tag, idLastLoaded string, chunkSize int) ([]models.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTag", ctx, tag, idLastLoaded, chunkSize)
	ret0, _ := ret[0].([]models.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTag indicates an expected call of GetByTag
func (mr *MockArticleUsecaseMockRecorder) GetByTag(ctx, tag, idLastLoaded, chunkSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTag", reflect.TypeOf((*MockArticleUsecase)(nil).GetByTag), ctx, tag, idLastLoaded, chunkSize)
}

// GetByAuthor mocks base method
func (m *MockArticleUsecase) GetByAuthor(ctx context.Context, author, idLastLoaded string, chunkSize int) ([]models.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAuthor", ctx, author, idLastLoaded, chunkSize)
	ret0, _ := ret[0].([]models.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAuthor indicates an expected call of GetByAuthor
func (mr *MockArticleUsecaseMockRecorder) GetByAuthor(ctx, author, idLastLoaded, chunkSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAuthor", reflect.TypeOf((*MockArticleUsecase)(nil).GetByAuthor), ctx, author, idLastLoaded, chunkSize)
}

// Update mocks base method
func (m *MockArticleUsecase) Update(ctx context.Context, a *models.ArticleUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockArticleUsecaseMockRecorder) Update(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticleUsecase)(nil).Update), ctx, a)
}

// Store mocks base method
func (m *MockArticleUsecase) Store(ctx context.Context, c *http.Cookie, a *models.ArticleCreate) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, c, a)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store
func (mr *MockArticleUsecaseMockRecorder) Store(ctx, c, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockArticleUsecase)(nil).Store), ctx, c, a)
}

// Delete mocks base method
func (m *MockArticleUsecase) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockArticleUsecaseMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArticleUsecase)(nil).Delete), ctx, id)
}

// MockArticleRepository is a mock of ArticleRepository interface
type MockArticleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockArticleRepositoryMockRecorder
}

// MockArticleRepositoryMockRecorder is the mock recorder for MockArticleRepository
type MockArticleRepositoryMockRecorder struct {
	mock *MockArticleRepository
}

// NewMockArticleRepository creates a new mock instance
func NewMockArticleRepository(ctrl *gomock.Controller) *MockArticleRepository {
	mock := &MockArticleRepository{ctrl: ctrl}
	mock.recorder = &MockArticleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockArticleRepository) EXPECT() *MockArticleRepositoryMockRecorder {
	return m.recorder
}

// Fetch mocks base method
func (m *MockArticleRepository) Fetch(ctx context.Context, from, chunkSize int) ([]models.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, from, chunkSize)
	ret0, _ := ret[0].([]models.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockArticleRepositoryMockRecorder) Fetch(ctx, from, chunkSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockArticleRepository)(nil).Fetch), ctx, from, chunkSize)
}

// GetByID mocks base method
func (m *MockArticleRepository) GetByID(ctx context.Context, id int64) (models.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(models.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockArticleRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockArticleRepository)(nil).GetByID), ctx, id)
}

// GetByTag mocks base method
func (m *MockArticleRepository) GetByTag(ctx context.Context, tag string, from, chunkSize int) ([]models.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTag", ctx, tag, from, chunkSize)
	ret0, _ := ret[0].([]models.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTag indicates an expected call of GetByTag
func (mr *MockArticleRepositoryMockRecorder) GetByTag(ctx, tag, from, chunkSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTag", reflect.TypeOf((*MockArticleRepository)(nil).GetByTag), ctx, tag, from, chunkSize)
}

// GetByAuthor mocks base method
func (m *MockArticleRepository) GetByAuthor(ctx context.Context, author string, from, chunkSize int) ([]models.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAuthor", ctx, author, from, chunkSize)
	ret0, _ := ret[0].([]models.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAuthor indicates an expected call of GetByAuthor
func (mr *MockArticleRepositoryMockRecorder) GetByAuthor(ctx, author, from, chunkSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAuthor", reflect.TypeOf((*MockArticleRepository)(nil).GetByAuthor), ctx, author, from, chunkSize)
}

// Update mocks base method
func (m *MockArticleRepository) Update(ctx context.Context, a *models.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockArticleRepositoryMockRecorder) Update(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticleRepository)(nil).Update), ctx, a)
}

// Store mocks base method
func (m *MockArticleRepository) Store(ctx context.Context, a *models.Article) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, a)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store
func (mr *MockArticleRepositoryMockRecorder) Store(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockArticleRepository)(nil).Store), ctx, a)
}

// Delete mocks base method
func (m *MockArticleRepository) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockArticleRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArticleRepository)(nil).Delete), ctx, id)
}
