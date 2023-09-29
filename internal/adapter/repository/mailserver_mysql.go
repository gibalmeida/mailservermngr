package repository

import (
	"context"
	"database/sql"

	"github.com/gibalmeida/mailservermngr/internal/app/apperror"
	"github.com/gibalmeida/mailservermngr/internal/domain"
	"github.com/gibalmeida/mailservermngr/pkg/db/mysql"
	"github.com/gibalmeida/mailservermngr/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlMailServerRepository struct {
	db      *sql.DB
	queries *mysql.Queries
}

func NewMysqlMailServerRepository(DatabaseURI string) (*MysqlMailServerRepository, error) {

	db, err := sql.Open("mysql", DatabaseURI)
	if err != nil {
		return &MysqlMailServerRepository{}, err
	}

	queries := mysql.New(db)

	return &MysqlMailServerRepository{
		db:      db,
		queries: queries,
	}, nil
}

func (r MysqlMailServerRepository) Close() {
	if r.db != nil {
		r.db.Close()
	}
}

func (r MysqlMailServerRepository) CreateAccount(ctx context.Context, newAccount domain.NewAccount) error {

	passwordHashed, err := utils.HashPassword(newAccount.Password)
	if err != nil {
		return err
	}
	_, err = r.queries.CreateAccount(context.Background(), mysql.CreateAccountParams{
		Name:     newAccount.Name,
		Domain:   newAccount.Domain,
		Password: string(passwordHashed[:]),
		HomeDir:  sql.NullString{String: "/home/vpopmail/domains/" + newAccount.Domain + "/" + newAccount.Name, Valid: true}, // TODO: Get the directory path from the application configuration.
		Quota:    sql.NullString{String: "NOQUOTA", Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r MysqlMailServerRepository) UpdateAccountPassword(ctx context.Context, name, emailDomain, clearTextPassword string) error {

	passwordHashed, err := utils.HashPassword(clearTextPassword)
	if err != nil {
		return err
	}
	return r.queries.UpdateAccountPassword(ctx,
		mysql.UpdateAccountPasswordParams{
			Name:     name,
			Domain:   emailDomain,
			Password: passwordHashed,
		})

}

func (r MysqlMailServerRepository) DeleteAccount(ctx context.Context, name, emailDomain string) error {

	return r.queries.DeleteAccount(ctx, mysql.DeleteAccountParams{Name: name, Domain: emailDomain})

}

func (r MysqlMailServerRepository) GetAccount(ctx context.Context, name, addressDomain string) (*domain.Account, error) {

	account, err := r.queries.GetAccount(ctx, mysql.GetAccountParams{
		Name:   name,
		Domain: addressDomain,
	})

	if err == sql.ErrNoRows {
		return &domain.Account{}, apperror.ErrAccountNotExist
	} else if err != nil {
		return &domain.Account{}, err
	}

	return &domain.Account{
		Name:   account.Name,
		Domain: account.Domain,
	}, nil

}

func (r MysqlMailServerRepository) GetAccounts(ctx context.Context) ([]*domain.Account, error) {
	var result []*domain.Account

	accounts, err := r.queries.GetAccounts(ctx)

	if err != nil {
		return result, err
	}

	for _, account := range accounts {
		result = append(result,
			&domain.Account{
				Domain: account.Domain,
				Name:   account.Name,
			})
	}

	return result, nil

}

func (r MysqlMailServerRepository) GetAccountsByDomain(ctx context.Context, addressDomain string) ([]*domain.Account, error) {
	var result []*domain.Account

	accounts, err := r.queries.GetAccountsFilteredByDomain(ctx, addressDomain)

	if err != nil {
		return result, err
	}

	for _, account := range accounts {
		result = append(result,
			&domain.Account{
				Domain: account.Domain,
				Name:   account.Name,
			})
	}

	return result, nil

}

func (r MysqlMailServerRepository) CreateAddressAlias(ctx context.Context, alias string, addresses string) error {
	_, err := r.queries.CreateAddressAlias(ctx, mysql.CreateAddressAliasParams{
		Alias:     alias,
		Addresses: sql.NullString{String: addresses, Valid: true},
	})

	return err
}

func (r MysqlMailServerRepository) DeleteAddressAlias(ctx context.Context, alias string) error {

	return r.queries.DeleteAddressAlias(ctx, alias)
}

func (r MysqlMailServerRepository) UpdateAddressAlias(ctx context.Context, addressAlias domain.AddressAlias) (*domain.AddressAlias, error) {

	err := r.queries.UpdateAddressAlias(ctx, mysql.UpdateAddressAliasParams{
		Alias: addressAlias.Alias,
		Addresses: sql.NullString{
			String: addressAlias.Addresses,
			Valid:  true,
		},
	})
	if err != nil {
		return &domain.AddressAlias{}, err
	}

	result, err := r.queries.GetAddressAlias(ctx, addressAlias.Alias)

	return &domain.AddressAlias{Alias: result.Alias, Addresses: result.Addresses.String}, err
}

func (r MysqlMailServerRepository) GetAddressAlias(ctx context.Context, alias string) (*domain.AddressAlias, error) {
	var result *domain.AddressAlias

	addressAlias, err := r.queries.GetAddressAlias(ctx, alias)

	if err == sql.ErrNoRows {
		return result, apperror.ErrAddressAliasNotExist
	} else if err != nil {
		return result, err
	} else {
		result.Alias = addressAlias.Alias
		result.Addresses = addressAlias.Addresses.String
	}

	return result, nil

}

func (r MysqlMailServerRepository) GetAddressesAliases(ctx context.Context) ([]*domain.AddressAlias, error) {
	var result []*domain.AddressAlias

	addressesAliases, err := r.queries.GetAddressesAliases(ctx)

	if err != nil {
		return result, err
	}

	for _, addressAlias := range addressesAliases {
		result = append(result,
			&domain.AddressAlias{
				Alias:     addressAlias.Alias,
				Addresses: addressAlias.Addresses.String,
			})
	}

	return result, nil
}

func (r MysqlMailServerRepository) GetAddressesAliasesByDomain(ctx context.Context, addressDomain string) ([]*domain.AddressAlias, error) {
	var result []*domain.AddressAlias

	addressesAliases, err := r.queries.GetAddressesAliasesFilteredByDomain(ctx, "%"+addressDomain+"%") // The percentage signals is a trick to be used on the SQL query that are using LIKE

	if err != nil {
		return result, err
	}

	for _, addressAlias := range addressesAliases {
		result = append(result,
			&domain.AddressAlias{
				Alias:     addressAlias.Alias,
				Addresses: addressAlias.Alias,
			})
	}

	return result, nil
}

func (r MysqlMailServerRepository) CreateDomainAlias(ctx context.Context, alias string, emailDomain string) error {
	_, err := r.queries.CreateDomainAlias(ctx, mysql.CreateDomainAliasParams{
		Alias:  alias,
		Domain: emailDomain,
	})

	return err
}

func (r MysqlMailServerRepository) DeleteDomainAlias(ctx context.Context, alias string) error {

	return r.queries.DeleteDomainAlias(ctx, alias)
}

func (r MysqlMailServerRepository) GetDomainAlias(ctx context.Context, alias string) (*domain.DomainAlias, error) {
	var result *domain.DomainAlias

	emailDomainAlias, err := r.queries.GetDomainAlias(ctx, alias)

	if err == sql.ErrNoRows {
		return result, apperror.ErrDomainAliasNotExist
	} else if err != nil {
		return result, err
	} else {
		result.Alias = emailDomainAlias.Alias
		result.Domain = emailDomainAlias.Domain
	}

	return result, nil

}

func (r MysqlMailServerRepository) GetDomainsAliases(ctx context.Context) ([]*domain.DomainAlias, error) {
	var result []*domain.DomainAlias

	emailDomainsAliases, err := r.queries.GetDomainsAliases(ctx)

	if err != nil {
		return result, err
	}

	for _, emailDomainAlias := range emailDomainsAliases {
		result = append(result,
			&domain.DomainAlias{
				Alias:  emailDomainAlias.Alias,
				Domain: emailDomainAlias.Domain,
			})
	}

	return result, nil
}

func (r MysqlMailServerRepository) GetDomainsAliasesByDomain(ctx context.Context, addressDomain string) ([]*domain.DomainAlias, error) {
	var result []*domain.DomainAlias

	domainsAliases, err := r.queries.GetDomainsAliasesFilteredByDomain(ctx, addressDomain)

	if err != nil {
		return result, err
	}

	for _, domainAlias := range domainsAliases {
		result = append(result,
			&domain.DomainAlias{
				Alias:  domainAlias.Alias,
				Domain: domainAlias.Alias,
			})
	}

	return result, nil
}

func (r MysqlMailServerRepository) CreateDomain(ctx context.Context, emailDomain string) error {
	_, err := r.queries.CreateDomain(ctx, emailDomain)

	return err
}

func (r MysqlMailServerRepository) DeleteDomain(ctx context.Context, emailDomain string) error {

	return r.queries.DeleteDomain(ctx, emailDomain)
}

func (r MysqlMailServerRepository) GetDomain(ctx context.Context, addressDomain string) (*domain.Domain, error) {
	var result *domain.Domain

	emailDomainRow, err := r.queries.GetDomain(ctx, addressDomain)

	if err == sql.ErrNoRows {
		return result, apperror.ErrDomainNotExist
	} else if err != nil {
		return result, err
	} else {
		result.Domain = emailDomainRow
	}

	return result, nil

}

func (r MysqlMailServerRepository) GetDomains(ctx context.Context) ([]*domain.Domain, error) {
	var result []*domain.Domain

	emailDomains, err := r.queries.GetDomains(ctx)

	if err != nil {
		return result, err
	}

	for _, addressDomain := range emailDomains {
		result = append(result,
			&domain.Domain{
				Domain: addressDomain,
			})
	}

	return result, nil
}
