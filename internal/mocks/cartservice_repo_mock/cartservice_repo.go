// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go
//
// Generated by this command:
//
//	mockgen -source=./service.go -destination=../../mocks/cartservice_repo_mock/cartservice_repo.go -package=cartservicerepomock .
//

// Package cartservicerepomock is a generated GoMock package.
package cartservicerepomock

import (
	context "context"
	entity "interview/internal/entity"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateCart mocks base method.
func (m *MockRepository) CreateCart(ctx context.Context, cart entity.Cart) (entity.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", ctx, cart)
	ret0, _ := ret[0].(entity.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockRepositoryMockRecorder) CreateCart(ctx, cart any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockRepository)(nil).CreateCart), ctx, cart)
}

// CreateCartItem mocks base method.
func (m *MockRepository) CreateCartItem(ctx context.Context, cartItem entity.CartItem) (entity.CartItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCartItem", ctx, cartItem)
	ret0, _ := ret[0].(entity.CartItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCartItem indicates an expected call of CreateCartItem.
func (mr *MockRepositoryMockRecorder) CreateCartItem(ctx, cartItem any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCartItem", reflect.TypeOf((*MockRepository)(nil).CreateCartItem), ctx, cartItem)
}

// DeleteCartItemByID mocks base method.
func (m *MockRepository) DeleteCartItemByID(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCartItemByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCartItemByID indicates an expected call of DeleteCartItemByID.
func (mr *MockRepositoryMockRecorder) DeleteCartItemByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCartItemByID", reflect.TypeOf((*MockRepository)(nil).DeleteCartItemByID), ctx, id)
}

// FindCartItemByID mocks base method.
func (m *MockRepository) FindCartItemByID(ctx context.Context, id uint) (entity.CartItem, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCartItemByID", ctx, id)
	ret0, _ := ret[0].(entity.CartItem)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindCartItemByID indicates an expected call of FindCartItemByID.
func (mr *MockRepositoryMockRecorder) FindCartItemByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCartItemByID", reflect.TypeOf((*MockRepository)(nil).FindCartItemByID), ctx, id)
}

// FindCartItemByProduct mocks base method.
func (m *MockRepository) FindCartItemByProduct(ctx context.Context, cartID uint, product string) (entity.CartItem, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCartItemByProduct", ctx, cartID, product)
	ret0, _ := ret[0].(entity.CartItem)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindCartItemByProduct indicates an expected call of FindCartItemByProduct.
func (mr *MockRepositoryMockRecorder) FindCartItemByProduct(ctx, cartID, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCartItemByProduct", reflect.TypeOf((*MockRepository)(nil).FindCartItemByProduct), ctx, cartID, product)
}

// FindCartItemsByCartID mocks base method.
func (m *MockRepository) FindCartItemsByCartID(ctx context.Context, cartID uint) ([]entity.CartItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCartItemsByCartID", ctx, cartID)
	ret0, _ := ret[0].([]entity.CartItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCartItemsByCartID indicates an expected call of FindCartItemsByCartID.
func (mr *MockRepositoryMockRecorder) FindCartItemsByCartID(ctx, cartID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCartItemsByCartID", reflect.TypeOf((*MockRepository)(nil).FindCartItemsByCartID), ctx, cartID)
}

// FindOpenCartBySessionID mocks base method.
func (m *MockRepository) FindOpenCartBySessionID(ctx context.Context, sessionID string) (entity.Cart, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOpenCartBySessionID", ctx, sessionID)
	ret0, _ := ret[0].(entity.Cart)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindOpenCartBySessionID indicates an expected call of FindOpenCartBySessionID.
func (mr *MockRepositoryMockRecorder) FindOpenCartBySessionID(ctx, sessionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOpenCartBySessionID", reflect.TypeOf((*MockRepository)(nil).FindOpenCartBySessionID), ctx, sessionID)
}

// UpdateCart mocks base method.
func (m *MockRepository) UpdateCart(ctx context.Context, cart entity.Cart) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", ctx, cart)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockRepositoryMockRecorder) UpdateCart(ctx, cart any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockRepository)(nil).UpdateCart), ctx, cart)
}

// UpdateCartItem mocks base method.
func (m *MockRepository) UpdateCartItem(ctx context.Context, cartItem entity.CartItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartItem", ctx, cartItem)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCartItem indicates an expected call of UpdateCartItem.
func (mr *MockRepositoryMockRecorder) UpdateCartItem(ctx, cartItem any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartItem", reflect.TypeOf((*MockRepository)(nil).UpdateCartItem), ctx, cartItem)
}
