package port

import (
	"context"

	"github.com/gibalmeida/mailservermngr/internal/domain"
)

type MailServerRepository interface {
	CreateAccount(ctx context.Context, newAccount domain.NewAccount) error
	GetAccount(ctx context.Context, name, emailDomain string) (*domain.Account, error)
	UpdateAccountPassword(ctx context.Context, name, domain, clearTextPassword string) error
	DeleteAccount(ctx context.Context, name string, emailDomain string) error
	GetAccounts(ctx context.Context) ([]*domain.Account, error)
	GetAccountsByDomain(ctx context.Context, emailDomain string) ([]*domain.Account, error)
	CreateAddressAlias(ctx context.Context, alias, addresses string) error
	DeleteAddressAlias(ctx context.Context, alias string) error
	UpdateAddressAlias(ctx context.Context, addressAlias domain.AddressAlias) (*domain.AddressAlias, error)
	GetAddressAlias(ctx context.Context, alias string) (*domain.AddressAlias, error)
	GetAddressesAliases(ctx context.Context) ([]*domain.AddressAlias, error)
	GetAddressesAliasesByDomain(ctx context.Context, emailDomain string) ([]*domain.AddressAlias, error)
	CreateDomain(ctx context.Context, emailDomain string) error
	DeleteDomain(ctx context.Context, emailDomain string) error
	GetDomain(ctx context.Context, emailDomain string) (*domain.Domain, error)
	GetDomains(ctx context.Context) ([]*domain.Domain, error)
	CreateDomainAlias(ctx context.Context, alias, emailDomain string) error
	DeleteDomainAlias(ctx context.Context, alias string) error
	GetDomainAlias(ctx context.Context, alias string) (*domain.DomainAlias, error)
	GetDomainsAliases(ctx context.Context) ([]*domain.DomainAlias, error)
	GetDomainsAliasesByDomain(ctx context.Context, emailDomain string) ([]*domain.DomainAlias, error)
}

type MailServerUseCase interface {
	CreateAccount(ctx context.Context, newAccount domain.NewAccount) error
	GetAccount(ctx context.Context, name, emailDomain string) (*domain.Account, error)
	UpdateAccountPassword(ctx context.Context, name, domain, clearTextPassword string) error
	DeleteAccount(ctx context.Context, name string, emailDomain string) error
	GetAccounts(ctx context.Context) ([]*domain.Account, error)
	GetAccountsByDomain(ctx context.Context, emailDomain string) ([]*domain.Account, error)
	CreateAddressAlias(ctx context.Context, alias, addresses string) error
	DeleteAddressAlias(ctx context.Context, alias string) error
	UpdateAddressAlias(ctx context.Context, addressAlias domain.AddressAlias) (*domain.AddressAlias, error)
	GetAddressAlias(ctx context.Context, alias string) (*domain.AddressAlias, error)
	GetAddressesAliases(ctx context.Context) ([]*domain.AddressAlias, error)
	GetAddressesAliasesByDomain(ctx context.Context, emailDomain string) ([]*domain.AddressAlias, error)
	CreateDomain(ctx context.Context, emailDomain string) error
	DeleteDomain(ctx context.Context, emailDomain string) error
	GetDomain(ctx context.Context, emailDomain string) (*domain.Domain, error)
	GetDomains(ctx context.Context) ([]*domain.Domain, error)
	CreateDomainAlias(ctx context.Context, alias, emailDomain string) error
	DeleteDomainAlias(ctx context.Context, alias string) error
	GetDomainAlias(ctx context.Context, alias string) (*domain.DomainAlias, error)
	GetDomainsAliases(ctx context.Context) ([]*domain.DomainAlias, error)
	GetDomainsAliasesByDomain(ctx context.Context, emailDomain string) ([]*domain.DomainAlias, error)
}
