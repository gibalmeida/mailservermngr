// Code generated by MockGen. DO NOT EDIT.
// Source: internal/port/mailserver.go
//
// Generated by this command:
//
//	mockgen -source=internal/port/mailserver.go -package mock -write_generate_directive -destination=internal/app/mock/mailserver_mock.go
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	domain "github.com/gibalmeida/mailservermngr/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

//go:generate mockgen -source=internal/port/mailserver.go -package mock -write_generate_directive -destination=internal/app/mock/mailserver_mock.go

// MockMailServerRepository is a mock of MailServerRepository interface.
type MockMailServerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMailServerRepositoryMockRecorder
}

// MockMailServerRepositoryMockRecorder is the mock recorder for MockMailServerRepository.
type MockMailServerRepositoryMockRecorder struct {
	mock *MockMailServerRepository
}

// NewMockMailServerRepository creates a new mock instance.
func NewMockMailServerRepository(ctrl *gomock.Controller) *MockMailServerRepository {
	mock := &MockMailServerRepository{ctrl: ctrl}
	mock.recorder = &MockMailServerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMailServerRepository) EXPECT() *MockMailServerRepositoryMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockMailServerRepository) CreateAccount(ctx context.Context, newAccount domain.NewAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, newAccount)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockMailServerRepositoryMockRecorder) CreateAccount(ctx, newAccount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockMailServerRepository)(nil).CreateAccount), ctx, newAccount)
}

// CreateAddressAlias mocks base method.
func (m *MockMailServerRepository) CreateAddressAlias(ctx context.Context, alias, addresses string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAddressAlias", ctx, alias, addresses)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAddressAlias indicates an expected call of CreateAddressAlias.
func (mr *MockMailServerRepositoryMockRecorder) CreateAddressAlias(ctx, alias, addresses any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAddressAlias", reflect.TypeOf((*MockMailServerRepository)(nil).CreateAddressAlias), ctx, alias, addresses)
}

// CreateDomain mocks base method.
func (m *MockMailServerRepository) CreateDomain(ctx context.Context, emailDomain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDomain", ctx, emailDomain)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDomain indicates an expected call of CreateDomain.
func (mr *MockMailServerRepositoryMockRecorder) CreateDomain(ctx, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDomain", reflect.TypeOf((*MockMailServerRepository)(nil).CreateDomain), ctx, emailDomain)
}

// CreateDomainAlias mocks base method.
func (m *MockMailServerRepository) CreateDomainAlias(ctx context.Context, alias, emailDomain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDomainAlias", ctx, alias, emailDomain)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDomainAlias indicates an expected call of CreateDomainAlias.
func (mr *MockMailServerRepositoryMockRecorder) CreateDomainAlias(ctx, alias, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDomainAlias", reflect.TypeOf((*MockMailServerRepository)(nil).CreateDomainAlias), ctx, alias, emailDomain)
}

// DeleteAccount mocks base method.
func (m *MockMailServerRepository) DeleteAccount(ctx context.Context, name, emailDomain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", ctx, name, emailDomain)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockMailServerRepositoryMockRecorder) DeleteAccount(ctx, name, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockMailServerRepository)(nil).DeleteAccount), ctx, name, emailDomain)
}

// DeleteAddressAlias mocks base method.
func (m *MockMailServerRepository) DeleteAddressAlias(ctx context.Context, alias string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddressAlias", ctx, alias)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddressAlias indicates an expected call of DeleteAddressAlias.
func (mr *MockMailServerRepositoryMockRecorder) DeleteAddressAlias(ctx, alias any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddressAlias", reflect.TypeOf((*MockMailServerRepository)(nil).DeleteAddressAlias), ctx, alias)
}

// DeleteDomain mocks base method.
func (m *MockMailServerRepository) DeleteDomain(ctx context.Context, emailDomain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDomain", ctx, emailDomain)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDomain indicates an expected call of DeleteDomain.
func (mr *MockMailServerRepositoryMockRecorder) DeleteDomain(ctx, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDomain", reflect.TypeOf((*MockMailServerRepository)(nil).DeleteDomain), ctx, emailDomain)
}

// DeleteDomainAlias mocks base method.
func (m *MockMailServerRepository) DeleteDomainAlias(ctx context.Context, alias string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDomainAlias", ctx, alias)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDomainAlias indicates an expected call of DeleteDomainAlias.
func (mr *MockMailServerRepositoryMockRecorder) DeleteDomainAlias(ctx, alias any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDomainAlias", reflect.TypeOf((*MockMailServerRepository)(nil).DeleteDomainAlias), ctx, alias)
}

// GetAccount mocks base method.
func (m *MockMailServerRepository) GetAccount(ctx context.Context, name, emailDomain string) (*domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, name, emailDomain)
	ret0, _ := ret[0].(*domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockMailServerRepositoryMockRecorder) GetAccount(ctx, name, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockMailServerRepository)(nil).GetAccount), ctx, name, emailDomain)
}

// GetAccounts mocks base method.
func (m *MockMailServerRepository) GetAccounts(ctx context.Context) ([]*domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", ctx)
	ret0, _ := ret[0].([]*domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockMailServerRepositoryMockRecorder) GetAccounts(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockMailServerRepository)(nil).GetAccounts), ctx)
}

// GetAccountsByDomain mocks base method.
func (m *MockMailServerRepository) GetAccountsByDomain(ctx context.Context, emailDomain string) ([]*domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountsByDomain", ctx, emailDomain)
	ret0, _ := ret[0].([]*domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountsByDomain indicates an expected call of GetAccountsByDomain.
func (mr *MockMailServerRepositoryMockRecorder) GetAccountsByDomain(ctx, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountsByDomain", reflect.TypeOf((*MockMailServerRepository)(nil).GetAccountsByDomain), ctx, emailDomain)
}

// GetAddressAlias mocks base method.
func (m *MockMailServerRepository) GetAddressAlias(ctx context.Context, alias string) (*domain.AddressAlias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressAlias", ctx, alias)
	ret0, _ := ret[0].(*domain.AddressAlias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressAlias indicates an expected call of GetAddressAlias.
func (mr *MockMailServerRepositoryMockRecorder) GetAddressAlias(ctx, alias any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressAlias", reflect.TypeOf((*MockMailServerRepository)(nil).GetAddressAlias), ctx, alias)
}

// GetAddressesAliases mocks base method.
func (m *MockMailServerRepository) GetAddressesAliases(ctx context.Context) ([]*domain.AddressAlias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressesAliases", ctx)
	ret0, _ := ret[0].([]*domain.AddressAlias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressesAliases indicates an expected call of GetAddressesAliases.
func (mr *MockMailServerRepositoryMockRecorder) GetAddressesAliases(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressesAliases", reflect.TypeOf((*MockMailServerRepository)(nil).GetAddressesAliases), ctx)
}

// GetAddressesAliasesByDomain mocks base method.
func (m *MockMailServerRepository) GetAddressesAliasesByDomain(ctx context.Context, emailDomain string) ([]*domain.AddressAlias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressesAliasesByDomain", ctx, emailDomain)
	ret0, _ := ret[0].([]*domain.AddressAlias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressesAliasesByDomain indicates an expected call of GetAddressesAliasesByDomain.
func (mr *MockMailServerRepositoryMockRecorder) GetAddressesAliasesByDomain(ctx, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressesAliasesByDomain", reflect.TypeOf((*MockMailServerRepository)(nil).GetAddressesAliasesByDomain), ctx, emailDomain)
}

// GetDomain mocks base method.
func (m *MockMailServerRepository) GetDomain(ctx context.Context, emailDomain string) (*domain.Domain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomain", ctx, emailDomain)
	ret0, _ := ret[0].(*domain.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomain indicates an expected call of GetDomain.
func (mr *MockMailServerRepositoryMockRecorder) GetDomain(ctx, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomain", reflect.TypeOf((*MockMailServerRepository)(nil).GetDomain), ctx, emailDomain)
}

// GetDomainAlias mocks base method.
func (m *MockMailServerRepository) GetDomainAlias(ctx context.Context, alias string) (*domain.DomainAlias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainAlias", ctx, alias)
	ret0, _ := ret[0].(*domain.DomainAlias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainAlias indicates an expected call of GetDomainAlias.
func (mr *MockMailServerRepositoryMockRecorder) GetDomainAlias(ctx, alias any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainAlias", reflect.TypeOf((*MockMailServerRepository)(nil).GetDomainAlias), ctx, alias)
}

// GetDomains mocks base method.
func (m *MockMailServerRepository) GetDomains(ctx context.Context) ([]*domain.Domain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomains", ctx)
	ret0, _ := ret[0].([]*domain.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomains indicates an expected call of GetDomains.
func (mr *MockMailServerRepositoryMockRecorder) GetDomains(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomains", reflect.TypeOf((*MockMailServerRepository)(nil).GetDomains), ctx)
}

// GetDomainsAliases mocks base method.
func (m *MockMailServerRepository) GetDomainsAliases(ctx context.Context) ([]*domain.DomainAlias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainsAliases", ctx)
	ret0, _ := ret[0].([]*domain.DomainAlias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainsAliases indicates an expected call of GetDomainsAliases.
func (mr *MockMailServerRepositoryMockRecorder) GetDomainsAliases(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainsAliases", reflect.TypeOf((*MockMailServerRepository)(nil).GetDomainsAliases), ctx)
}

// UpdateAccountPassword mocks base method.
func (m *MockMailServerRepository) UpdateAccountPassword(ctx context.Context, name, domain, clearTextPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccountPassword", ctx, name, domain, clearTextPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccountPassword indicates an expected call of UpdateAccountPassword.
func (mr *MockMailServerRepositoryMockRecorder) UpdateAccountPassword(ctx, name, domain, clearTextPassword any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountPassword", reflect.TypeOf((*MockMailServerRepository)(nil).UpdateAccountPassword), ctx, name, domain, clearTextPassword)
}

// UpdateAddressAlias mocks base method.
func (m *MockMailServerRepository) UpdateAddressAlias(ctx context.Context, addressAlias domain.AddressAlias) (*domain.AddressAlias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddressAlias", ctx, addressAlias)
	ret0, _ := ret[0].(*domain.AddressAlias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAddressAlias indicates an expected call of UpdateAddressAlias.
func (mr *MockMailServerRepositoryMockRecorder) UpdateAddressAlias(ctx, addressAlias any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddressAlias", reflect.TypeOf((*MockMailServerRepository)(nil).UpdateAddressAlias), ctx, addressAlias)
}

// MockMailServerUseCase is a mock of MailServerUseCase interface.
type MockMailServerUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockMailServerUseCaseMockRecorder
}

// MockMailServerUseCaseMockRecorder is the mock recorder for MockMailServerUseCase.
type MockMailServerUseCaseMockRecorder struct {
	mock *MockMailServerUseCase
}

// NewMockMailServerUseCase creates a new mock instance.
func NewMockMailServerUseCase(ctrl *gomock.Controller) *MockMailServerUseCase {
	mock := &MockMailServerUseCase{ctrl: ctrl}
	mock.recorder = &MockMailServerUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMailServerUseCase) EXPECT() *MockMailServerUseCaseMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockMailServerUseCase) CreateAccount(ctx context.Context, newAccount domain.NewAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, newAccount)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockMailServerUseCaseMockRecorder) CreateAccount(ctx, newAccount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockMailServerUseCase)(nil).CreateAccount), ctx, newAccount)
}

// CreateAddressAlias mocks base method.
func (m *MockMailServerUseCase) CreateAddressAlias(ctx context.Context, alias, addresses string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAddressAlias", ctx, alias, addresses)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAddressAlias indicates an expected call of CreateAddressAlias.
func (mr *MockMailServerUseCaseMockRecorder) CreateAddressAlias(ctx, alias, addresses any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAddressAlias", reflect.TypeOf((*MockMailServerUseCase)(nil).CreateAddressAlias), ctx, alias, addresses)
}

// CreateDomain mocks base method.
func (m *MockMailServerUseCase) CreateDomain(ctx context.Context, emailDomain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDomain", ctx, emailDomain)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDomain indicates an expected call of CreateDomain.
func (mr *MockMailServerUseCaseMockRecorder) CreateDomain(ctx, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDomain", reflect.TypeOf((*MockMailServerUseCase)(nil).CreateDomain), ctx, emailDomain)
}

// CreateDomainAlias mocks base method.
func (m *MockMailServerUseCase) CreateDomainAlias(ctx context.Context, alias, emailDomain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDomainAlias", ctx, alias, emailDomain)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDomainAlias indicates an expected call of CreateDomainAlias.
func (mr *MockMailServerUseCaseMockRecorder) CreateDomainAlias(ctx, alias, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDomainAlias", reflect.TypeOf((*MockMailServerUseCase)(nil).CreateDomainAlias), ctx, alias, emailDomain)
}

// DeleteAccount mocks base method.
func (m *MockMailServerUseCase) DeleteAccount(ctx context.Context, name, emailDomain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", ctx, name, emailDomain)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockMailServerUseCaseMockRecorder) DeleteAccount(ctx, name, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockMailServerUseCase)(nil).DeleteAccount), ctx, name, emailDomain)
}

// DeleteAddressAlias mocks base method.
func (m *MockMailServerUseCase) DeleteAddressAlias(ctx context.Context, alias string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddressAlias", ctx, alias)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddressAlias indicates an expected call of DeleteAddressAlias.
func (mr *MockMailServerUseCaseMockRecorder) DeleteAddressAlias(ctx, alias any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddressAlias", reflect.TypeOf((*MockMailServerUseCase)(nil).DeleteAddressAlias), ctx, alias)
}

// DeleteDomain mocks base method.
func (m *MockMailServerUseCase) DeleteDomain(ctx context.Context, emailDomain string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDomain", ctx, emailDomain)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDomain indicates an expected call of DeleteDomain.
func (mr *MockMailServerUseCaseMockRecorder) DeleteDomain(ctx, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDomain", reflect.TypeOf((*MockMailServerUseCase)(nil).DeleteDomain), ctx, emailDomain)
}

// DeleteDomainAlias mocks base method.
func (m *MockMailServerUseCase) DeleteDomainAlias(ctx context.Context, alias string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDomainAlias", ctx, alias)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDomainAlias indicates an expected call of DeleteDomainAlias.
func (mr *MockMailServerUseCaseMockRecorder) DeleteDomainAlias(ctx, alias any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDomainAlias", reflect.TypeOf((*MockMailServerUseCase)(nil).DeleteDomainAlias), ctx, alias)
}

// GetAccount mocks base method.
func (m *MockMailServerUseCase) GetAccount(ctx context.Context, name, emailDomain string) (*domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, name, emailDomain)
	ret0, _ := ret[0].(*domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockMailServerUseCaseMockRecorder) GetAccount(ctx, name, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockMailServerUseCase)(nil).GetAccount), ctx, name, emailDomain)
}

// GetAccounts mocks base method.
func (m *MockMailServerUseCase) GetAccounts(ctx context.Context) ([]*domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", ctx)
	ret0, _ := ret[0].([]*domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockMailServerUseCaseMockRecorder) GetAccounts(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockMailServerUseCase)(nil).GetAccounts), ctx)
}

// GetAccountsByDomain mocks base method.
func (m *MockMailServerUseCase) GetAccountsByDomain(ctx context.Context, emailDomain string) ([]*domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountsByDomain", ctx, emailDomain)
	ret0, _ := ret[0].([]*domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountsByDomain indicates an expected call of GetAccountsByDomain.
func (mr *MockMailServerUseCaseMockRecorder) GetAccountsByDomain(ctx, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountsByDomain", reflect.TypeOf((*MockMailServerUseCase)(nil).GetAccountsByDomain), ctx, emailDomain)
}

// GetAddressAlias mocks base method.
func (m *MockMailServerUseCase) GetAddressAlias(ctx context.Context, alias string) (*domain.AddressAlias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressAlias", ctx, alias)
	ret0, _ := ret[0].(*domain.AddressAlias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressAlias indicates an expected call of GetAddressAlias.
func (mr *MockMailServerUseCaseMockRecorder) GetAddressAlias(ctx, alias any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressAlias", reflect.TypeOf((*MockMailServerUseCase)(nil).GetAddressAlias), ctx, alias)
}

// GetAddressesAliases mocks base method.
func (m *MockMailServerUseCase) GetAddressesAliases(ctx context.Context) ([]*domain.AddressAlias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressesAliases", ctx)
	ret0, _ := ret[0].([]*domain.AddressAlias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressesAliases indicates an expected call of GetAddressesAliases.
func (mr *MockMailServerUseCaseMockRecorder) GetAddressesAliases(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressesAliases", reflect.TypeOf((*MockMailServerUseCase)(nil).GetAddressesAliases), ctx)
}

// GetAddressesAliasesByDomain mocks base method.
func (m *MockMailServerUseCase) GetAddressesAliasesByDomain(ctx context.Context, emailDomain string) ([]*domain.AddressAlias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressesAliasesByDomain", ctx, emailDomain)
	ret0, _ := ret[0].([]*domain.AddressAlias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressesAliasesByDomain indicates an expected call of GetAddressesAliasesByDomain.
func (mr *MockMailServerUseCaseMockRecorder) GetAddressesAliasesByDomain(ctx, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressesAliasesByDomain", reflect.TypeOf((*MockMailServerUseCase)(nil).GetAddressesAliasesByDomain), ctx, emailDomain)
}

// GetDomain mocks base method.
func (m *MockMailServerUseCase) GetDomain(ctx context.Context, emailDomain string) (*domain.Domain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomain", ctx, emailDomain)
	ret0, _ := ret[0].(*domain.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomain indicates an expected call of GetDomain.
func (mr *MockMailServerUseCaseMockRecorder) GetDomain(ctx, emailDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomain", reflect.TypeOf((*MockMailServerUseCase)(nil).GetDomain), ctx, emailDomain)
}

// GetDomainAlias mocks base method.
func (m *MockMailServerUseCase) GetDomainAlias(ctx context.Context, alias string) (*domain.DomainAlias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainAlias", ctx, alias)
	ret0, _ := ret[0].(*domain.DomainAlias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainAlias indicates an expected call of GetDomainAlias.
func (mr *MockMailServerUseCaseMockRecorder) GetDomainAlias(ctx, alias any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainAlias", reflect.TypeOf((*MockMailServerUseCase)(nil).GetDomainAlias), ctx, alias)
}

// GetDomains mocks base method.
func (m *MockMailServerUseCase) GetDomains(ctx context.Context) ([]*domain.Domain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomains", ctx)
	ret0, _ := ret[0].([]*domain.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomains indicates an expected call of GetDomains.
func (mr *MockMailServerUseCaseMockRecorder) GetDomains(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomains", reflect.TypeOf((*MockMailServerUseCase)(nil).GetDomains), ctx)
}

// GetDomainsAliases mocks base method.
func (m *MockMailServerUseCase) GetDomainsAliases(ctx context.Context) ([]*domain.DomainAlias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainsAliases", ctx)
	ret0, _ := ret[0].([]*domain.DomainAlias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainsAliases indicates an expected call of GetDomainsAliases.
func (mr *MockMailServerUseCaseMockRecorder) GetDomainsAliases(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainsAliases", reflect.TypeOf((*MockMailServerUseCase)(nil).GetDomainsAliases), ctx)
}

// UpdateAccountPassword mocks base method.
func (m *MockMailServerUseCase) UpdateAccountPassword(ctx context.Context, name, domain, clearTextPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccountPassword", ctx, name, domain, clearTextPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccountPassword indicates an expected call of UpdateAccountPassword.
func (mr *MockMailServerUseCaseMockRecorder) UpdateAccountPassword(ctx, name, domain, clearTextPassword any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountPassword", reflect.TypeOf((*MockMailServerUseCase)(nil).UpdateAccountPassword), ctx, name, domain, clearTextPassword)
}

// UpdateAddressAlias mocks base method.
func (m *MockMailServerUseCase) UpdateAddressAlias(ctx context.Context, addressAlias domain.AddressAlias) (*domain.AddressAlias, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddressAlias", ctx, addressAlias)
	ret0, _ := ret[0].(*domain.AddressAlias)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAddressAlias indicates an expected call of UpdateAddressAlias.
func (mr *MockMailServerUseCaseMockRecorder) UpdateAddressAlias(ctx, addressAlias any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddressAlias", reflect.TypeOf((*MockMailServerUseCase)(nil).UpdateAddressAlias), ctx, addressAlias)
}
