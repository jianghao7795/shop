// Code generated by MockGen. DO NOT EDIT.
// Source: user/internal/biz (interfaces: UserRepo)

// Package mrepo is a generated GoMock package.
package mrepo

import (
	context "context"
	reflect "reflect"
	biz "user/internal/biz"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// CheckPassword implements biz.UserRepo.
func (m *MockUserRepo) CheckPassword(ctx context.Context, password string, encryptedPassword string) (bool, error) {
	panic("unimplemented")
}

// GetUserById implements biz.UserRepo.
func (m *MockUserRepo) GetUserById(ctx context.Context, id int64) (*biz.User, error) {
	panic("unimplemented")
}

// ListUser implements biz.UserRepo.
func (m *MockUserRepo) ListUser(ctx context.Context, pageNum int, pageSize int) ([]*biz.User, int, error) {
	panic("unimplemented")
}

// UpdateUser implements biz.UserRepo.
func (m *MockUserRepo) UpdateUser(context.Context, *biz.User) (bool, error) {
	panic("unimplemented")
}

// UserByMobile implements biz.UserRepo.
func (m *MockUserRepo) UserByMobile(ctx context.Context, mobile string) (*biz.User, error) {
	panic("unimplemented")
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepo) CreateUser(arg0 context.Context, arg1 *biz.User) (*biz.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*biz.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepoMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepo)(nil).CreateUser), arg0, arg1)
}
