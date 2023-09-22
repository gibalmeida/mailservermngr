package repository

import (
	"context"
	"strings"
	"sync"

	"github.com/gibalmeida/mailservermngr/internal/app/apperror"
	"github.com/gibalmeida/mailservermngr/internal/domain"
)

type MemMailServerRepository struct {
	Accounts         map[string]*Account
	AddressesAliases map[string]*domain.AddressAlias
	Domains          map[string]*domain.Domain
	DomainsAliases   map[string]*domain.DomainAlias
	lock             *sync.RWMutex
}

type Account struct {
	Domain   string
	Name     string
	Password string
}

func NewMemMailServerRepository() *MemMailServerRepository {

	return &MemMailServerRepository{
		Accounts:         make(map[string]*Account),
		AddressesAliases: make(map[string]*domain.AddressAlias),
		Domains:          make(map[string]*domain.Domain),
		DomainsAliases:   make(map[string]*domain.DomainAlias),
		lock:             &sync.RWMutex{},
	}

}

func (r MemMailServerRepository) Close() {

}
func (r MemMailServerRepository) CreateAccount(ctx context.Context, newAccount domain.NewAccount) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	index := newAccount.Name + "@" + newAccount.Domain
	_, exist := r.Accounts[index]

	if exist {
		return apperror.ErrAccountAlreadyExist
	}

	r.Accounts[index] = &Account{Domain: newAccount.Domain, Name: newAccount.Name, Password: newAccount.Password}

	return nil
}

func (r MemMailServerRepository) UpdateAccountPassword(ctx context.Context, name, emailDomain, clearTextPassword string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	account, exist := r.Accounts[name+"@"+emailDomain]

	if !exist {
		return apperror.ErrAccountNotExist
	}

	account.Password = clearTextPassword

	return nil
}

func (r MemMailServerRepository) DeleteAccount(ctx context.Context, name, emailDomain string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	_, exist := r.Accounts[name+"@"+emailDomain]

	if !exist {
		return apperror.ErrAccountNotExist
	}

	delete(r.Accounts, name+"@"+emailDomain)

	return nil

}

func (r MemMailServerRepository) GetAccount(ctx context.Context, name, emailDomain string) (*domain.Account, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	account, exist := r.Accounts[name+"@"+emailDomain]

	if !exist {
		return nil, apperror.ErrAccountNotExist
	}
	return &domain.Account{Domain: account.Domain, Name: account.Name}, nil

}

func (r MemMailServerRepository) GetAccounts(ctx context.Context) ([]*domain.Account, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	accounts := []*domain.Account{}

	for _, a := range r.Accounts {
		accounts = append(accounts, &domain.Account{Domain: a.Domain, Name: a.Name})
	}
	return accounts, nil

}

func (r MemMailServerRepository) GetAccountsByDomain(ctx context.Context, emailDomain string) ([]*domain.Account, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	accounts := []*domain.Account{}

	for _, a := range r.Accounts {
		if a.Domain == emailDomain {
			accounts = append(accounts, &domain.Account{Domain: a.Domain, Name: a.Name})
		}
	}
	return accounts, nil

}

func (r MemMailServerRepository) CreateAddressAlias(ctx context.Context, alias string, addresses string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	_, exist := r.AddressesAliases[alias]

	if exist {
		return apperror.ErrAddressAliasAlreadyExist
	}

	r.AddressesAliases[alias] = &domain.AddressAlias{Alias: alias, Addresses: addresses}

	return nil

}

func (r MemMailServerRepository) DeleteAddressAlias(ctx context.Context, alias string) error {

	r.lock.Lock()
	defer r.lock.Unlock()

	_, exist := r.AddressesAliases[alias]

	if !exist {
		return apperror.ErrAddressAliasNotExist
	}

	delete(r.AddressesAliases, alias)

	return nil

}

func (r MemMailServerRepository) UpdateAddressAlias(ctx context.Context, addressAlias domain.AddressAlias) (*domain.AddressAlias, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	alias := addressAlias.Alias
	_, exist := r.AddressesAliases[alias]

	if !exist {
		return &domain.AddressAlias{}, apperror.ErrAddressAliasNotExist
	}

	r.AddressesAliases[alias] = &addressAlias

	return r.AddressesAliases[alias], nil
}

func (r MemMailServerRepository) GetAddressAlias(ctx context.Context, alias string) (*domain.AddressAlias, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	addressAlias, exist := r.AddressesAliases[alias]

	if !exist {
		return nil, apperror.ErrAddressAliasNotExist
	}
	return addressAlias, nil

}

func (r MemMailServerRepository) GetAddressesAliases(ctx context.Context) ([]*domain.AddressAlias, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	addressesAliases := []*domain.AddressAlias{}

	for _, a := range r.AddressesAliases {
		addressesAliases = append(addressesAliases, a)
	}
	return addressesAliases, nil

}

func (r MemMailServerRepository) GetAddressesAliasesByDomain(ctx context.Context, emailDomain string) ([]*domain.AddressAlias, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	addressesAliases := []*domain.AddressAlias{}

	for _, a := range r.AddressesAliases {
		if strings.Contains(a.Addresses, emailDomain) {
			addressesAliases = append(addressesAliases, a)
		}
	}
	return addressesAliases, nil
}

func (r MemMailServerRepository) CreateDomainAlias(ctx context.Context, alias string, emailDomain string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	_, exist := r.DomainsAliases[alias]

	if exist {
		return apperror.ErrDomainAliasAlreadyExist
	}

	r.DomainsAliases[alias] = &domain.DomainAlias{Alias: alias, Domain: emailDomain}

	return nil
}

func (r MemMailServerRepository) DeleteDomainAlias(ctx context.Context, alias string) error {

	r.lock.Lock()
	defer r.lock.Unlock()

	_, exist := r.DomainsAliases[alias]

	if !exist {
		return apperror.ErrDomainNotExist
	}

	delete(r.DomainsAliases, alias)

	return nil
}

func (r MemMailServerRepository) GetDomainAlias(ctx context.Context, alias string) (*domain.DomainAlias, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	domainAlias, exist := r.DomainsAliases[alias]

	if !exist {
		return nil, apperror.ErrDomainAliasNotExist
	}
	return domainAlias, nil

}

func (r MemMailServerRepository) GetDomainsAliases(ctx context.Context) ([]*domain.DomainAlias, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	domainsAliases := []*domain.DomainAlias{}

	for _, a := range r.DomainsAliases {
		domainsAliases = append(domainsAliases, a)
	}
	return domainsAliases, nil
}

func (r MemMailServerRepository) CreateDomain(ctx context.Context, emailDomain string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	_, exist := r.Domains[emailDomain]

	if exist {
		return apperror.ErrDomainAlreadyExist
	}

	r.Domains[emailDomain] = &domain.Domain{Domain: emailDomain}

	return nil
}

func (r MemMailServerRepository) DeleteDomain(ctx context.Context, emailDomain string) error {

	r.lock.Lock()
	defer r.lock.Unlock()

	_, exist := r.Domains[emailDomain]

	if !exist {
		return apperror.ErrDomainNotExist
	}

	delete(r.Domains, emailDomain)

	return nil
}

func (r MemMailServerRepository) GetDomain(ctx context.Context, emailDomain string) (*domain.Domain, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	result, exist := r.Domains[emailDomain]

	if !exist {
		return nil, apperror.ErrDomainNotExist
	}
	return result, nil

}

func (r MemMailServerRepository) GetDomains(ctx context.Context) ([]*domain.Domain, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	domains := []*domain.Domain{}

	for _, a := range r.Domains {
		domains = append(domains, a)
	}
	return domains, nil
}
