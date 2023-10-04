package usecase_test

import (
	"context"
	"testing"

	"github.com/gibalmeida/mailservermngr/internal/adapter/repository"
	"github.com/gibalmeida/mailservermngr/internal/app/apperror"
	"github.com/gibalmeida/mailservermngr/internal/app/usecase"
	"github.com/gibalmeida/mailservermngr/internal/domain"
	"github.com/gibalmeida/mailservermngr/internal/port"

	"github.com/stretchr/testify/assert"
)

func setupMailServerTestCase(t *testing.T) (context.Context, port.MailServerUseCase, func(t *testing.T)) {
	t.Log("setup test case")

	ctx := context.TODO()
	repo := repository.NewMemMailServerRepository()
	uc := usecase.NewMailServerUseCase(repo)

	_ = uc.CreateDomain(ctx, "test.com")
	_ = uc.CreateDomain(ctx, "test2.com")
	_ = uc.CreateDomain(ctx, "test3.com")
	_ = uc.CreateDomainAlias(ctx, "alias.com", "test.com")
	_ = uc.CreateDomainAlias(ctx, "alias2.com", "test2.com")
	_ = uc.CreateDomainAlias(ctx, "alias3.com", "test3.com")
	_ = uc.CreateAccount(ctx, domain.NewAccount{Name: "exist", Domain: "test.com", Password: "pass"})
	_ = uc.CreateAccount(ctx, domain.NewAccount{Name: "email", Domain: "test2.com", Password: "pass"})
	_ = uc.CreateAccount(ctx, domain.NewAccount{Name: "email", Domain: "test3.com", Password: "pass"})
	_ = uc.CreateAddressAlias(ctx, domain.AddressAlias{Alias: "alias@test.com", Addresses: []domain.EmailAddress{"exist@test.com"}})
	_ = uc.CreateAddressAlias(ctx, domain.AddressAlias{Alias: "alias@test2.com", Addresses: []domain.EmailAddress{"email@test2.com"}})
	_ = uc.CreateAddressAlias(ctx, domain.AddressAlias{Alias: "alias@test3.com", Addresses: []domain.EmailAddress{"email@test3.com"}})

	return ctx, uc, func(t *testing.T) {
		t.Log("teardown test case")

		_ = uc.DeleteDomain(ctx, "test.com")
		_ = uc.DeleteDomain(ctx, "test2.com")
		_ = uc.DeleteDomain(ctx, "test3.com")
		_ = uc.DeleteDomainAlias(ctx, "alias.com")
		_ = uc.DeleteDomainAlias(ctx, "alias2.com")
		_ = uc.DeleteDomainAlias(ctx, "alias3.com")
		_ = uc.DeleteAccount(ctx, "exist", "test.com")
		_ = uc.DeleteAccount(ctx, "email", "test2.com")
		_ = uc.DeleteAccount(ctx, "email", "test3.com")
		_ = uc.DeleteAddressAlias(ctx, "alias@test.com")
		_ = uc.DeleteAddressAlias(ctx, "alias@test2.com")
		_ = uc.DeleteAddressAlias(ctx, "alias@test3.com")
	}
}

func TestGetAccount(t *testing.T) {

	testCases := []struct {
		desc    string
		name    string
		domain  string
		account *domain.Account
		err     error
	}{
		{
			desc:   "success",
			name:   "exist",
			domain: "test.com",
			account: &domain.Account{
				Name:   "exist",
				Domain: "test.com",
			},
		},
		{
			desc: "account doesn't exist",
			err:  apperror.ErrAccountNotExist,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := uc.GetAccount(ctx, tc.name, tc.domain)
			if tc.err == nil {
				assert.Equal(t, tc.account, result)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestGetAccounts(t *testing.T) {

	testCases := []struct {
		desc   string
		length int
		err    error
	}{
		{
			desc:   "success",
			length: 3,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := uc.GetAccounts(ctx)
			if tc.err == nil {
				assert.Equal(t, tc.length, len(result))
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestGetAccountsByDomain(t *testing.T) {

	testCases := []struct {
		desc   string
		domain string
		length int
		err    error
	}{
		{
			desc:   "success",
			domain: "test.com",
			length: 1,
		},
		{
			desc:   "empty",
			domain: "nonexist.com",
			length: 0,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := uc.GetAccountsByDomain(ctx, tc.domain)
			if tc.err == nil {
				assert.Equal(t, tc.length, len(result))
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestCreateAccount(t *testing.T) {

	testCases := []struct {
		desc    string
		account *domain.NewAccount
		err     error
	}{
		{
			desc: "success",
			account: &domain.NewAccount{
				Name:     "test",
				Domain:   "test.com",
				Password: "password",
			},
		},
		{
			desc: "account already exist",
			account: &domain.NewAccount{
				Name:     "exist",
				Domain:   "test.com",
				Password: "password",
			},
			err: apperror.ErrAccountAlreadyExist,
		},
		{
			desc: "domain doesn't exist",
			account: &domain.NewAccount{
				Name:     "test",
				Domain:   "nonexist.com",
				Password: "password",
			},
			err: apperror.ErrDomainNotExist,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := uc.CreateAccount(ctx, *tc.account)
			if tc.err == nil {
				assert.NoError(t, err)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestDeleteAccount(t *testing.T) {

	testCases := []struct {
		desc    string
		account *domain.Account
		err     error
	}{
		{
			desc: "success",
			account: &domain.Account{
				Name:   "exist",
				Domain: "test.com",
			},
		},
		{
			desc: "account doesn't exist",
			account: &domain.Account{
				Name:   "nonexist",
				Domain: "test.com",
			},
			err: apperror.ErrAccountNotExist,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := uc.DeleteAccount(ctx, tc.account.Name, tc.account.Domain)
			if tc.err == nil {
				assert.NoError(t, err)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestChangeAccountPassword(t *testing.T) {

	testCases := []struct {
		desc        string
		account     *domain.Account
		newPassword string
		err         error
	}{
		{
			desc: "success",
			account: &domain.Account{
				Name:   "exist",
				Domain: "test.com",
			},
			newPassword: "new-password",
		},
		{
			desc: "account doesn't exist",
			account: &domain.Account{
				Name:   "nonexist",
				Domain: "test.com",
			},
			newPassword: "new-password",
			err:         apperror.ErrAccountNotExist,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := uc.UpdateAccountPassword(ctx, tc.account.Name, tc.account.Domain, tc.newPassword)
			if tc.err == nil {
				result, err := uc.GetAccount(ctx, tc.account.Name, tc.account.Domain)
				assert.NoError(t, err)
				assert.Equal(t, tc.account, result)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestGetAddressAlias(t *testing.T) {

	testCases := []struct {
		desc   string
		alias  string
		result *domain.AddressAlias
		err    error
	}{
		{
			desc:  "success",
			alias: "alias@test.com",
			result: &domain.AddressAlias{
				Alias:     "alias@test.com",
				Addresses: []domain.EmailAddress{"exist@test.com"},
			},
		},
		{
			desc:  "address alias doesn't exist",
			alias: "nonexist@test.com",
			err:   apperror.ErrAddressAliasNotExist,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := uc.GetAddressAlias(ctx, tc.alias)
			if tc.err == nil {
				assert.Equal(t, tc.result, result)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestGetAddressesAliases(t *testing.T) {

	testCases := []struct {
		desc   string
		length int
		err    error
	}{
		{
			desc:   "success",
			length: 3,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := uc.GetAddressesAliases(ctx)
			if tc.err == nil {
				assert.Equal(t, tc.length, len(result))
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestGetAddressesAliasesByDomain(t *testing.T) {

	testCases := []struct {
		desc   string
		domain string
		length int
		err    error
	}{
		{
			desc:   "success",
			domain: "test.com",
			length: 1,
		},
		{
			desc:   "empty",
			domain: "nonexist.com",
			length: 0,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := uc.GetAddressesAliasesByDomain(ctx, tc.domain)
			if tc.err == nil {
				assert.Equal(t, tc.length, len(result))
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestCreateAddressAlias(t *testing.T) {

	testCases := []struct {
		desc         string
		addressAlias *domain.AddressAlias
		err          error
	}{
		{
			desc: "success",
			addressAlias: &domain.AddressAlias{
				Alias:     "newalias@test.com",
				Addresses: []domain.EmailAddress{"exist@domain.com"},
			},
		},
		{
			desc: "address alias already exist",
			addressAlias: &domain.AddressAlias{
				Alias:     "alias@test.com",
				Addresses: []domain.EmailAddress{"exist@test.com"},
			},
			err: apperror.ErrAddressAliasAlreadyExist,
		},
		{
			desc: "domain or domain alias doesn't exist",
			addressAlias: &domain.AddressAlias{
				Alias:     "secondalias@nonexist.com",
				Addresses: []domain.EmailAddress{"exist@test.com"},
			},
			err: apperror.ErrDomainOrDomainAliasNotExist,
		},
		{
			desc: "already exist as a regular address",
			addressAlias: &domain.AddressAlias{
				Alias:     "email@test2.com",
				Addresses: []domain.EmailAddress{"exist@test.com"},
			},
			err: apperror.ErrAddressAliasAlreadyExistAsRegular,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := uc.CreateAddressAlias(ctx, *tc.addressAlias)
			if tc.err == nil {
				assert.NoError(t, err)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestDeleteAddressAlias(t *testing.T) {

	testCases := []struct {
		desc         string
		addressAlias string
		err          error
	}{
		{
			desc:         "success",
			addressAlias: "alias@test.com",
		},
		{
			desc:         "address alias doesn't exist",
			addressAlias: "nonexist@test.com",
			err:          apperror.ErrAddressAliasNotExist,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := uc.DeleteAddressAlias(ctx, tc.addressAlias)
			if tc.err == nil {
				assert.NoError(t, err)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestUpdateAddressAlias(t *testing.T) {

	testCases := []struct {
		desc         string
		addressAlias *domain.AddressAlias
		err          error
	}{
		{
			desc: "success",
			addressAlias: &domain.AddressAlias{
				Alias:     "alias@test.com",
				Addresses: []domain.EmailAddress{"exist@test.com", "email@test2.com", "email@test3.com"},
			},
		},
		{
			desc: "address alias doesn't exist",
			addressAlias: &domain.AddressAlias{
				Alias:     "nonexist@test.com",
				Addresses: []domain.EmailAddress{"email@test2.com"},
			},
			err: apperror.ErrAddressAliasNotExist,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := uc.UpdateAddressAlias(ctx, *tc.addressAlias)
			if tc.err == nil {
				assert.Equal(t, result, tc.addressAlias)
				assert.NoError(t, err)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestCreateDomain(t *testing.T) {

	testCases := []struct {
		desc   string
		domain *domain.Domain
		err    error
	}{
		{
			desc: "success",
			domain: &domain.Domain{
				Domain: "new.com",
			},
		},
		{
			desc: "domain already exist",
			domain: &domain.Domain{
				Domain: "test.com",
			},
			err: apperror.ErrDomainOrDomainAliasAlreadyExist,
		},
		{
			desc: "domain alias already exist",
			domain: &domain.Domain{
				Domain: "alias.com",
			},
			err: apperror.ErrDomainOrDomainAliasAlreadyExist,
		},
	}
	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := uc.CreateDomain(ctx, tc.domain.Domain)
			if tc.err == nil {
				assert.NoError(t, err)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestGetDomain(t *testing.T) {

	testCases := []struct {
		desc   string
		domain string
		result *domain.Domain
		err    error
	}{
		{
			desc:   "success",
			domain: "test.com",
			result: &domain.Domain{
				Domain: "test.com",
			},
		},
		{
			desc:   "domain doesn't exist",
			domain: "nonexist.com",
			result: &domain.Domain{},
			err:    apperror.ErrDomainNotExist,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := uc.GetDomain(ctx, tc.domain)
			if tc.err == nil {
				assert.Equal(t, tc.result, result)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestGetDomains(t *testing.T) {

	testCases := []struct {
		desc   string
		length int
		err    error
	}{
		{
			desc:   "success",
			length: 3,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := uc.GetDomains(ctx)
			if tc.err == nil {
				assert.Equal(t, tc.length, len(result))
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestGetDomainsAliasesByDomain(t *testing.T) {

	testCases := []struct {
		desc   string
		domain string
		length int
		err    error
	}{
		{
			desc:   "success",
			domain: "test.com",
			length: 1,
		},
		{
			desc:   "empty",
			domain: "nonexist.com",
			length: 0,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := uc.GetDomainsAliasesByDomain(ctx, tc.domain)
			if tc.err == nil {
				assert.Equal(t, tc.length, len(result))
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestDeleteDomain(t *testing.T) {

	testCases := []struct {
		desc   string
		domain *domain.Domain
		err    error
	}{
		{
			desc: "success",
			domain: &domain.Domain{
				Domain: "test.com",
			},
		},
		{
			desc: "domain doesn't exist",
			domain: &domain.Domain{
				Domain: "nonexist.com",
			},
			err: apperror.ErrDomainNotExist,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := uc.DeleteDomain(ctx, tc.domain.Domain)
			if tc.err == nil {
				assert.NoError(t, err)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestGetDomainAlias(t *testing.T) {

	testCases := []struct {
		desc   string
		alias  string
		result *domain.DomainAlias
		err    error
	}{
		{
			desc:  "success",
			alias: "alias.com",
			result: &domain.DomainAlias{
				Alias:  "alias.com",
				Domain: "test.com",
			},
		},
		{
			desc:  "domain alias doesn't exist",
			alias: "nonexist@test.com",
			err:   apperror.ErrDomainAliasNotExist,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := uc.GetDomainAlias(ctx, tc.alias)
			if tc.err == nil {
				assert.Equal(t, tc.result, result)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestGetDomainAliases(t *testing.T) {

	testCases := []struct {
		desc   string
		length int
		err    error
	}{
		{
			desc:   "success",
			length: 3,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := uc.GetDomainsAliases(ctx)
			if tc.err == nil {
				assert.Equal(t, tc.length, len(result))
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestCreateDomainAlias(t *testing.T) {

	testCases := []struct {
		desc        string
		domainAlias *domain.DomainAlias
		err         error
	}{
		{
			desc: "success",
			domainAlias: &domain.DomainAlias{
				Alias:  "new.com",
				Domain: "test.com",
			},
		},
		{
			desc: "domain alias already exist",
			domainAlias: &domain.DomainAlias{
				Alias:  "alias.com",
				Domain: "test.com",
			},
			err: apperror.ErrDomainOrDomainAliasAlreadyExist,
		},
		{
			desc: "domain already exist",
			domainAlias: &domain.DomainAlias{
				Alias:  "test.com",
				Domain: "test.com",
			},
			err: apperror.ErrDomainOrDomainAliasAlreadyExist,
		},
		{
			desc: "domain doesn't exist",
			domainAlias: &domain.DomainAlias{
				Alias:  "second.com",
				Domain: "nonexist.com",
			},
			err: apperror.ErrDomainNotExist,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := uc.CreateDomainAlias(ctx, tc.domainAlias.Alias, tc.domainAlias.Domain)
			if tc.err == nil {
				assert.NoError(t, err)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}

func TestDeleteDomainAlias(t *testing.T) {

	testCases := []struct {
		desc        string
		domainAlias string
		err         error
	}{
		{
			desc:        "success",
			domainAlias: "alias.com",
		},
		{
			desc:        "domain alias doesn't exist",
			domainAlias: "nonexist.com",
			err:         apperror.ErrDomainAliasNotExist,
		},
	}

	ctx, uc, teardownTestCase := setupMailServerTestCase(t)
	defer teardownTestCase(t)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := uc.DeleteDomainAlias(ctx, tc.domainAlias)
			if tc.err == nil {
				assert.NoError(t, err)
			} else if assert.Error(t, err) {
				assert.Equal(t, tc.err, err)
			}

		})
	}
}
