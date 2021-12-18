// Code generated by MockGen. DO NOT EDIT.
// Source: session.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	gomock "github.com/golang/mock/gomock"
)

// MockSessionUsecase is a mock of SessionUsecase interface.
type MockSessionUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockSessionUsecaseMockRecorder
}

// MockSessionUsecaseMockRecorder is the mock recorder for MockSessionUsecase.
type MockSessionUsecaseMockRecorder struct {
	mock *MockSessionUsecase
}

// NewMockSessionUsecase creates a new mock instance.
func NewMockSessionUsecase(ctrl *gomock.Controller) *MockSessionUsecase {
	mock := &MockSessionUsecase{ctrl: ctrl}
	mock.recorder = &MockSessionUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionUsecase) EXPECT() *MockSessionUsecaseMockRecorder {
	return m.recorder
}

// IsSession mocks base method.
func (m *MockSessionUsecase) IsSession(ctx context.Context, cookie string) (models.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSession", ctx, cookie)
	ret0, _ := ret[0].(models.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSession indicates an expected call of IsSession.
func (mr *MockSessionUsecaseMockRecorder) IsSession(ctx, cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSession", reflect.TypeOf((*MockSessionUsecase)(nil).IsSession), ctx, cookie)
}

// MockSessionRepository is a mock of SessionRepository interface.
type MockSessionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSessionRepositoryMockRecorder
}

// MockSessionRepositoryMockRecorder is the mock recorder for MockSessionRepository.
type MockSessionRepositoryMockRecorder struct {
	mock *MockSessionRepository
}

// NewMockSessionRepository creates a new mock instance.
func NewMockSessionRepository(ctrl *gomock.Controller) *MockSessionRepository {
	mock := &MockSessionRepository{ctrl: ctrl}
	mock.recorder = &MockSessionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionRepository) EXPECT() *MockSessionRepositoryMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockSessionRepository) CreateSession(ctx context.Context, login string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, login)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockSessionRepositoryMockRecorder) CreateSession(ctx, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockSessionRepository)(nil).CreateSession), ctx, login)
}

// DeleteSession mocks base method.
func (m *MockSessionRepository) DeleteSession(ctx context.Context, cookieValue string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", ctx, cookieValue)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockSessionRepositoryMockRecorder) DeleteSession(ctx, cookieValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockSessionRepository)(nil).DeleteSession), ctx, cookieValue)
}

// GetSessionLogin mocks base method.
func (m *MockSessionRepository) GetSessionLogin(ctx context.Context, cookie string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionLogin", ctx, cookie)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionLogin indicates an expected call of GetSessionLogin.
func (mr *MockSessionRepositoryMockRecorder) GetSessionLogin(ctx, cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionLogin", reflect.TypeOf((*MockSessionRepository)(nil).GetSessionLogin), ctx, cookie)
}