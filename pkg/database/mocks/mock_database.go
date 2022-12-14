// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	model "beneburg/pkg/database/model"
	context "context"
	gomock "github.com/golang/mock/gomock"
	gen "gorm.io/gen"
	reflect "reflect"
)

// MockDatabase is a mock of Database interface
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// AutoMigrate mocks base method
func (m *MockDatabase) AutoMigrate(models ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range models {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AutoMigrate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AutoMigrate indicates an expected call of AutoMigrate
func (mr *MockDatabaseMockRecorder) AutoMigrate(models ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoMigrate", reflect.TypeOf((*MockDatabase)(nil).AutoMigrate), models...)
}

// CreateUser mocks base method
func (m *MockDatabase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockDatabaseMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDatabase)(nil).CreateUser), ctx, user)
}

// UpdateOrCreateUser mocks base method
func (m *MockDatabase) UpdateOrCreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrCreateUser", ctx, user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrCreateUser indicates an expected call of UpdateOrCreateUser
func (mr *MockDatabaseMockRecorder) UpdateOrCreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrCreateUser", reflect.TypeOf((*MockDatabase)(nil).UpdateOrCreateUser), ctx, user)
}

// CreateToken mocks base method
func (m *MockDatabase) CreateToken(ctx context.Context, telegramID int64) (*model.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", ctx, telegramID)
	ret0, _ := ret[0].(*model.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken
func (mr *MockDatabaseMockRecorder) CreateToken(ctx, telegramID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockDatabase)(nil).CreateToken), ctx, telegramID)
}

// GetAllUsers mocks base method
func (m *MockDatabase) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", ctx)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers
func (mr *MockDatabaseMockRecorder) GetAllUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockDatabase)(nil).GetAllUsers), ctx)
}

// GetUserByID mocks base method
func (m *MockDatabase) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID
func (mr *MockDatabaseMockRecorder) GetUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockDatabase)(nil).GetUserByID), ctx, id)
}

// GetUserByTelegramID mocks base method
func (m *MockDatabase) GetUserByTelegramID(ctx context.Context, telegramID int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByTelegramID", ctx, telegramID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByTelegramID indicates an expected call of GetUserByTelegramID
func (mr *MockDatabaseMockRecorder) GetUserByTelegramID(ctx, telegramID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByTelegramID", reflect.TypeOf((*MockDatabase)(nil).GetUserByTelegramID), ctx, telegramID)
}

// UpdateUserByID mocks base method
func (m *MockDatabase) UpdateUserByID(ctx context.Context, id uint, user *model.User, updateFieldNames ...string) (*gen.ResultInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id, user}
	for _, a := range updateFieldNames {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUserByID", varargs...)
	ret0, _ := ret[0].(*gen.ResultInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserByID indicates an expected call of UpdateUserByID
func (mr *MockDatabaseMockRecorder) UpdateUserByID(ctx, id, user interface{}, updateFieldNames ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id, user}, updateFieldNames...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserByID", reflect.TypeOf((*MockDatabase)(nil).UpdateUserByID), varargs...)
}

// GetUserIDByToken mocks base method
func (m *MockDatabase) GetUserIDByToken(ctx context.Context, token string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIDByToken", ctx, token)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIDByToken indicates an expected call of GetUserIDByToken
func (mr *MockDatabaseMockRecorder) GetUserIDByToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIDByToken", reflect.TypeOf((*MockDatabase)(nil).GetUserIDByToken), ctx, token)
}
