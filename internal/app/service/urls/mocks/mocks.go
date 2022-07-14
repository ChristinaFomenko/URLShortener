// Code generated by MockGen. DO NOT EDIT.
// Source: urls.go

// Package mock_urls is a generated GoMock package.
package mock_urls

import (
	context "context"
	reflect "reflect"

	models "github.com/ChristinaFomenko/shortener/internal/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockurlRepository is a mock of urlRepository interface.
type MockurlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockurlRepositoryMockRecorder
}

// MockurlRepositoryMockRecorder is the mock recorder for MockurlRepository.
type MockurlRepositoryMockRecorder struct {
	mock *MockurlRepository
}

// NewMockurlRepository creates a new mock instance.
func NewMockurlRepository(ctrl *gomock.Controller) *MockurlRepository {
	mock := &MockurlRepository{ctrl: ctrl}
	mock.recorder = &MockurlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockurlRepository) EXPECT() *MockurlRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockurlRepository) Add(ctx context.Context, urlID, url, userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, urlID, url, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockurlRepositoryMockRecorder) Add(ctx, urlID, url, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockurlRepository)(nil).Add), ctx, urlID, url, userID)
}

// AddBatch mocks base method.
func (m *MockurlRepository) AddBatch(ctx context.Context, urls []models.UserURL, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBatch", ctx, urls, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBatch indicates an expected call of AddBatch.
func (mr *MockurlRepositoryMockRecorder) AddBatch(ctx, urls, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBatch", reflect.TypeOf((*MockurlRepository)(nil).AddBatch), ctx, urls, userID)
}

// FetchURLs mocks base method.
func (m *MockurlRepository) FetchURLs(ctx context.Context, userID string) ([]models.UserURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchURLs", ctx, userID)
	ret0, _ := ret[0].([]models.UserURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchURLs indicates an expected call of FetchURLs.
func (mr *MockurlRepositoryMockRecorder) FetchURLs(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchURLs", reflect.TypeOf((*MockurlRepository)(nil).FetchURLs), ctx, userID)
}

// Get mocks base method.
func (m *MockurlRepository) Get(ctx context.Context, urlID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, urlID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockurlRepositoryMockRecorder) Get(ctx, urlID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockurlRepository)(nil).Get), ctx, urlID)
}

// Mockgenerator is a mock of generator interface.
type Mockgenerator struct {
	ctrl     *gomock.Controller
	recorder *MockgeneratorMockRecorder
}

// MockgeneratorMockRecorder is the mock recorder for Mockgenerator.
type MockgeneratorMockRecorder struct {
	mock *Mockgenerator
}

// NewMockgenerator creates a new mock instance.
func NewMockgenerator(ctrl *gomock.Controller) *Mockgenerator {
	mock := &Mockgenerator{ctrl: ctrl}
	mock.recorder = &MockgeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockgenerator) EXPECT() *MockgeneratorMockRecorder {
	return m.recorder
}

// Letters mocks base method.
func (m *Mockgenerator) Letters(n int64) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Letters", n)
	ret0, _ := ret[0].(string)
	return ret0
}

// Letters indicates an expected call of Letters.
func (mr *MockgeneratorMockRecorder) Letters(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Letters", reflect.TypeOf((*Mockgenerator)(nil).Letters), n)
}
