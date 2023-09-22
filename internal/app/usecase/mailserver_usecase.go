package usecase

import (
	"context"
	"strings"

	"github.com/gibalmeida/mailservermngr/internal/app/apperror"
	"github.com/gibalmeida/mailservermngr/internal/domain"
	"github.com/gibalmeida/mailservermngr/internal/port"
)

type mailServerUseCase struct {
	repo port.MailServerRepository
}

func NewMailServerUseCase(repo port.MailServerRepository) port.MailServerUseCase {
	return &mailServerUseCase{
		repo: repo,
	}
}

func (u *mailServerUseCase) CreateAccount(ctx context.Context, newAccount domain.NewAccount) error {
	ok, err := u.domainExist(ctx, newAccount.Domain)
	if err != nil {
		return err
	} else if !ok {
		return apperror.ErrDomainNotExist
	}

	_, err = u.repo.GetAccount(ctx, newAccount.Name, newAccount.Domain)

	if err == nil {
		return apperror.ErrAccountAlreadyExist
	}

	return u.repo.CreateAccount(ctx, newAccount)
}

func (u *mailServerUseCase) GetAccount(ctx context.Context, name, emailDomain string) (*domain.Account, error) {

	return u.repo.GetAccount(ctx, name, emailDomain)
}

func (u *mailServerUseCase) UpdateAccountPassword(ctx context.Context, name, domain, clearTextPassword string) error {
	_, err := u.repo.GetAccount(ctx, name, domain)
	if err != nil {
		return err
	}

	return u.repo.UpdateAccountPassword(ctx, name, domain, clearTextPassword)
}

func (u *mailServerUseCase) DeleteAccount(ctx context.Context, name string, emailDomain string) error {
	_, err := u.repo.GetAccount(ctx, name, emailDomain)
	if err != nil {
		return err
	}

	if err = u.repo.DeleteAccount(ctx, name, emailDomain); err != nil {
		return err
	}

	// TODO: We still need to implement a way to delete the account folder on the email server. Otherwise, only the account is deleted, but the folder containing the emails will continue to exist on the server.

	return nil
}

func (u *mailServerUseCase) GetAccounts(ctx context.Context) ([]*domain.Account, error) {
	return u.repo.GetAccounts(ctx)
}

func (u *mailServerUseCase) GetAccountsByDomain(ctx context.Context, emailDomain string) ([]*domain.Account, error) {
	return u.repo.GetAccountsByDomain(ctx, emailDomain)
}

func (u *mailServerUseCase) CreateAddressAlias(ctx context.Context, alias, addresses string) error {
	_, err := u.repo.GetAddressAlias(ctx, alias)
	if err == nil {
		return apperror.ErrAddressAliasAlreadyExist
	}

	splittedEmailAddress := strings.Split(alias, "@")

	domainExist, err := u.domainOrDomainAliasExist(ctx, splittedEmailAddress[1])
	if err != nil {
		return err
	} else if !domainExist {
		return apperror.ErrDomainOrDomainAliasNotExist
	}

	_, err = u.repo.GetAccount(ctx, splittedEmailAddress[0], splittedEmailAddress[1])
	if err == nil {
		return apperror.ErrAddressAliasAlreadyExistAsRegular
	} else if err != apperror.ErrAccountNotExist {
		return err
	}

	return u.repo.CreateAddressAlias(ctx, alias, addresses)
}

func (u *mailServerUseCase) UpdateAddressAlias(ctx context.Context, addressAlias domain.AddressAlias) (*domain.AddressAlias, error) {
	_, err := u.repo.GetAddressAlias(ctx, addressAlias.Alias)
	if err != nil {
		return &domain.AddressAlias{}, err
	}

	return u.repo.UpdateAddressAlias(ctx, addressAlias)
}

func (u *mailServerUseCase) DeleteAddressAlias(ctx context.Context, alias string) error {
	_, err := u.repo.GetAddressAlias(ctx, alias)
	if err != nil {
		return err
	}

	return u.repo.DeleteAddressAlias(ctx, alias)
}

func (u *mailServerUseCase) GetAddressAlias(ctx context.Context, alias string) (*domain.AddressAlias, error) {
	return u.repo.GetAddressAlias(ctx, alias)
}

func (u *mailServerUseCase) GetAddressesAliases(ctx context.Context) ([]*domain.AddressAlias, error) {
	var addressesAliases []*domain.AddressAlias

	addressesAliases, err := u.repo.GetAddressesAliases(ctx)

	if err != nil {
		return addressesAliases, err
	}

	return addressesAliases, nil
}

func (u *mailServerUseCase) GetAddressesAliasesByDomain(ctx context.Context, emailDomain string) ([]*domain.AddressAlias, error) {
	return u.repo.GetAddressesAliasesByDomain(ctx, emailDomain)
}

func (u *mailServerUseCase) CreateDomain(ctx context.Context, emailDomain string) error {
	domainAlrearyRegistered, err := u.domainOrDomainAliasExist(ctx, emailDomain)
	if err != nil {
		return err
	} else if domainAlrearyRegistered {
		return apperror.ErrDomainOrDomainAliasAlreadyExist
	}

	return u.repo.CreateDomain(ctx, emailDomain)

}

func (u *mailServerUseCase) DeleteDomain(ctx context.Context, emailDomain string) error {
	_, err := u.repo.GetDomain(ctx, emailDomain)
	if err != nil {
		return err
	}

	return u.repo.DeleteDomain(ctx, emailDomain)
}

func (u *mailServerUseCase) GetDomain(ctx context.Context, emailDomain string) (*domain.Domain, error) {
	return u.repo.GetDomain(ctx, emailDomain)
}

func (u *mailServerUseCase) GetDomains(ctx context.Context) ([]*domain.Domain, error) {
	return u.repo.GetDomains(ctx)
}

func (u *mailServerUseCase) CreateDomainAlias(ctx context.Context, alias, emailDomain string) error {
	domainAlrearyRegistered, err := u.domainOrDomainAliasExist(ctx, alias)
	if err != nil {
		return err
	} else if domainAlrearyRegistered {
		return apperror.ErrDomainOrDomainAliasAlreadyExist
	}

	domainExist, err := u.domainExist(ctx, emailDomain)
	if err != nil {
		return err
	} else if !domainExist {
		return apperror.ErrDomainNotExist
	}

	return u.repo.CreateDomainAlias(ctx, alias, emailDomain)

}

func (u *mailServerUseCase) DeleteDomainAlias(ctx context.Context, alias string) error {
	_, err := u.repo.GetDomainAlias(ctx, alias)
	if err != nil {
		return err
	}

	return u.repo.DeleteDomainAlias(ctx, alias)
}

func (u *mailServerUseCase) GetDomainAlias(ctx context.Context, alias string) (*domain.DomainAlias, error) {
	return u.repo.GetDomainAlias(ctx, alias)
}

func (u *mailServerUseCase) GetDomainsAliases(ctx context.Context) ([]*domain.DomainAlias, error) {
	return u.repo.GetDomainsAliases(ctx)
}

// Checks if the domain or domain is already registered
func (u *mailServerUseCase) domainOrDomainAliasExist(ctx context.Context, emailDomain string) (bool, error) {
	_, err := u.repo.GetDomain(ctx, emailDomain)
	if err == apperror.ErrDomainNotExist {
		_, err = u.repo.GetDomainAlias(ctx, emailDomain)
		if err == apperror.ErrDomainAliasNotExist {
			return false, nil
		}
	} else if err != nil {
		return false, err
	}
	return true, nil
}

// Checks if the domain is already registered
func (u *mailServerUseCase) domainExist(ctx context.Context, emailDomain string) (bool, error) {
	_, err := u.repo.GetDomain(ctx, emailDomain)

	switch err {
	case nil:
		return true, nil
	case apperror.ErrDomainNotExist:
		return false, nil
	default:
		return false, err
	}
}
