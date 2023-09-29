package api_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/gibalmeida/mailservermngr/internal/adapter/http/api"
	"github.com/gibalmeida/mailservermngr/internal/adapter/http/api/model"
	"github.com/gibalmeida/mailservermngr/internal/app/mock"
	"github.com/gibalmeida/mailservermngr/internal/domain"
	"github.com/gibalmeida/mailservermngr/pkg/jwx"
	"go.uber.org/mock/gomock"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const PrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIBr4imcEmbV78LX5T6/rsZzo/YWPbtnhzYqOlF9OUL7LoAoGCCqGSM49
AwEHoUQDQgAEITmerwxXOEl9ER03yOaNqrmYFPRVSicjpmTl2wfqO1tY/mO662nJ
fy0+Ec3BO/UOyDYXLAg3Yi5iflNjdDMT1w==
-----END EC PRIVATE KEY-----`

func TestAPI(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockAuthUseCase := mock.NewMockAuthUseCase(ctrl)
	mockMailSrvUseCase := mock.NewMockMailServerUseCase(ctrl)
	defer ctrl.Finish()

	e := echo.New()

	// Log requests
	e.Use(echomiddleware.Logger())

	jwsAuth, err := jwx.NewJWSAuthenticator([]byte(PrivateKey))
	require.NoError(t, err)
	if err != nil {
		t.Logf("error creating authenticator: %s", err)
	}

	mw, err := api.CreateMiddleware(jwsAuth)
	require.NoError(t, err)
	e.Use(mw...)

	s := api.NewServer(mockAuthUseCase, mockMailSrvUseCase)

	api.RegisterHandlers(e, s)

	jwt, err := jwsAuth.CreateJWSWithClaims([]string{})
	require.NoError(t, err)

	t.Run("Should return 403 forbidden when without credentials", func(t *testing.T) {
		// GetAccounts should return 403 forbidden without credentials
		response := testutil.NewRequest().Get("/accounts").Go(t, e)
		assert.Equal(t, http.StatusForbidden, response.Code())
	})

	t.Run("Sucessfully get JWT token", func(t *testing.T) {
		token := "token"
		user := "user"
		password := "password"
		param := model.Auth{Username: user, Password: password}
		expected := model.GetTokenResponse{AccessToken: token}

		mockAuthUseCase.EXPECT().
			GetTokenUsingUsernameAndPassword(gomock.Any(), gomock.Eq(user), gomock.Eq(password)).
			Times(1).
			Return(token, nil)
		response := testutil.NewRequest().Post("/getToken").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			WithJsonBody(param).Go(t, e)

		require.Equal(t, http.StatusOK, response.Code())

		var result model.GetTokenResponse
		err = response.UnmarshalBodyToObject(&result)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, expected, result)
	})

	t.Run("Invalid format when create an account", func(t *testing.T) {
		mockMailSrvUseCase.EXPECT().
			CreateAccount(gomock.Any(), gomock.Any()).
			Times(0)
		response := testutil.NewRequest().Post("/accounts").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusBadRequest, response.Code())
	})

	t.Run("Try to get an account with a invalid email", func(t *testing.T) {
		invalidEmail := "email"

		mockMailSrvUseCase.EXPECT().
			GetAccount(gomock.Any(), gomock.Any(), gomock.Any()).
			Times(0)

		response := testutil.NewRequest().Get(fmt.Sprintf("/accounts/%s", invalidEmail)).
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusBadRequest, response.Code())

	})

	t.Run("Sucessfully creates an account", func(t *testing.T) {
		newAccount := model.NewAccount{Name: "email", Domain: "domain.com", Password: "password"}
		expectedAccount := model.Account{Name: newAccount.Name, Domain: newAccount.Domain}

		mockMailSrvUseCase.EXPECT().
			CreateAccount(gomock.Any(), gomock.Any()).
			Times(1).
			Return(nil)
		response := testutil.NewRequest().Post("/accounts").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			WithJsonBody(newAccount).Go(t, e)

		require.Equal(t, http.StatusCreated, response.Code())

		var responseAccount model.Account
		err = response.UnmarshalBodyToObject(&responseAccount)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, expectedAccount, responseAccount)

	})

	t.Run("Get a Bad Request when try create an account with a short name (1 character)", func(t *testing.T) {
		newAccount := model.NewAccount{Name: "e", Domain: "domain.com", Password: "password"}

		mockMailSrvUseCase.EXPECT().
			CreateAccount(gomock.Any(), gomock.Any()).
			Times(0).
			Return(nil)
		response := testutil.NewRequest().Post("/accounts").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			WithJsonBody(newAccount).Go(t, e)

		require.Equal(t, http.StatusBadRequest, response.Code())

	})

	t.Run("Get a Bad Request when try create an account without a domain name", func(t *testing.T) {
		newAccount := model.NewAccount{Name: "email", Domain: "d", Password: "12345"}

		mockMailSrvUseCase.EXPECT().
			CreateAccount(gomock.Any(), gomock.Any()).
			Times(0).
			Return(nil)
		response := testutil.NewRequest().Post("/accounts").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			WithJsonBody(newAccount).Go(t, e)

		require.Equal(t, http.StatusBadRequest, response.Code())

	})

	t.Run("Get a Bad Request when try create an account with a short password (less than 6 characters)", func(t *testing.T) {
		newAccount := model.NewAccount{Name: "email", Domain: "domain.com", Password: "12345"}

		mockMailSrvUseCase.EXPECT().
			CreateAccount(gomock.Any(), gomock.Any()).
			Times(0).
			Return(nil)
		response := testutil.NewRequest().Post("/accounts").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			WithJsonBody(newAccount).Go(t, e)

		require.Equal(t, http.StatusBadRequest, response.Code())

	})

	t.Run("Sucessfully get an account", func(t *testing.T) {
		name := "email"
		emailDomain := "domain.com"
		expectedAccount := model.Account{Name: name, Domain: emailDomain}

		mockMailSrvUseCase.EXPECT().
			GetAccount(gomock.Any(), name, emailDomain).
			Times(1).
			Return(&domain.Account{Name: name, Domain: emailDomain}, nil)
		response := testutil.NewRequest().Get(fmt.Sprintf("/accounts/%s", name+"@"+emailDomain)).
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusOK, response.Code())

		var responseAccount model.Account
		err = response.UnmarshalBodyToObject(&responseAccount)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, expectedAccount, responseAccount)

	})

	t.Run("Sucessfully delete an account", func(t *testing.T) {
		emailAddress := "email@domain.com"
		splittedEmailAddress := strings.Split(emailAddress, "@")

		mockMailSrvUseCase.EXPECT().
			DeleteAccount(gomock.Any(), splittedEmailAddress[0], splittedEmailAddress[1]).
			Times(1).
			Return(nil)
		response := testutil.NewRequest().Delete(fmt.Sprintf("/accounts/%s", emailAddress)).
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusNoContent, response.Code())
	})

	t.Run("Sucessfully get all accounts", func(t *testing.T) {

		domAllAccounts := []*domain.Account{{Name: "email", Domain: "domain.com"}}
		modAllAccounts := []model.Account{{Name: "email", Domain: "domain.com"}}

		mockMailSrvUseCase.EXPECT().
			GetAccounts(gomock.Any()).
			Times(1).
			Return(domAllAccounts, nil)
		response := testutil.NewRequest().Get("/accounts").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusOK, response.Code())

		var responseAllAccounts []model.Account
		err = response.UnmarshalBodyToObject(&responseAllAccounts)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, modAllAccounts, responseAllAccounts)
	})

	t.Run("Sucessfully get accounts filtered by domain", func(t *testing.T) {
		emailDomain := "domain.com"

		domAllAccounts := []*domain.Account{{Name: "email", Domain: "domain.com"}}
		modAllAccounts := []model.Account{{Name: "email", Domain: "domain.com"}}

		mockMailSrvUseCase.EXPECT().
			GetAccountsByDomain(gomock.Any(), emailDomain).
			Times(1).
			Return(domAllAccounts, nil)
		response := testutil.NewRequest().Get(fmt.Sprintf("/accounts/filterByDomain/%s", emailDomain)).
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusOK, response.Code())

		var responseAllAccounts []model.Account
		err = response.UnmarshalBodyToObject(&responseAllAccounts)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, modAllAccounts, responseAllAccounts)
	})
	t.Run("Sucessfully update an account password", func(t *testing.T) {
		name := "email"
		emailDomain := "domain.com"
		newPassword := "password"

		changeAccountPassword := model.ChangeAccountPassword{EmailAddress: name + "@" + emailDomain, NewPassword: newPassword}

		mockMailSrvUseCase.EXPECT().
			UpdateAccountPassword(gomock.Any(), name, emailDomain, newPassword).
			Times(1).
			Return(nil)
		response := testutil.NewRequest().Post("/accounts/changePassword").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			WithJsonBody(changeAccountPassword).Go(t, e)

		require.Equal(t, http.StatusNoContent, response.Code())

	})

	t.Run("Sucessfully creates an domain", func(t *testing.T) {
		newEmailDomain := "domain.com"
		newDomain := model.Domain{Domain: newEmailDomain}

		mockMailSrvUseCase.EXPECT().
			CreateDomain(gomock.Any(), newEmailDomain).
			Times(1).
			Return(nil)
		response := testutil.NewRequest().Post("/domains").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			WithJsonBody(newDomain).Go(t, e)

		require.Equal(t, http.StatusCreated, response.Code())

		var result model.Domain
		err = response.UnmarshalBodyToObject(&result)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, newDomain, result)

	})

	t.Run("Sucessfully delete an domain", func(t *testing.T) {
		emailDomain := "domain.com"

		mockMailSrvUseCase.EXPECT().
			DeleteDomain(gomock.Any(), emailDomain).
			Times(1).
			Return(nil)
		response := testutil.NewRequest().Delete(fmt.Sprintf("/domains/%s", emailDomain)).
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusNoContent, response.Code())
	})

	t.Run("Sucessfully get all domains", func(t *testing.T) {

		mockReturn := []*domain.Domain{{Domain: "domain.com"}}
		expected := []model.Domain{{Domain: "domain.com"}}

		mockMailSrvUseCase.EXPECT().
			GetDomains(gomock.Any()).
			Times(1).
			Return(mockReturn, nil)
		response := testutil.NewRequest().Get("/domains").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusOK, response.Code())

		var result []model.Domain
		err = response.UnmarshalBodyToObject(&result)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, expected, result)
	})

	t.Run("Sucessfully creates an domain alias", func(t *testing.T) {
		alias := "alias.com"
		emailDomain := "domain.com"
		param := model.DomainAlias{Alias: alias, Domain: emailDomain}
		expected := model.DomainAlias{Alias: alias, Domain: emailDomain}

		mockMailSrvUseCase.EXPECT().
			CreateDomainAlias(gomock.Any(), alias, emailDomain).
			Times(1).
			Return(nil)
		response := testutil.NewRequest().Post("/domains-aliases").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			WithJsonBody(param).Go(t, e)

		require.Equal(t, http.StatusCreated, response.Code())

		var result model.DomainAlias
		err = response.UnmarshalBodyToObject(&result)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, expected, result)

	})

	t.Run("Sucessfully delete an domain alias", func(t *testing.T) {
		alias := "alias.com"

		mockMailSrvUseCase.EXPECT().
			DeleteDomainAlias(gomock.Any(), alias).
			Times(1).
			Return(nil)
		response := testutil.NewRequest().Delete(fmt.Sprintf("/domains-aliases/%s", alias)).
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusNoContent, response.Code())
	})

	t.Run("Sucessfully get all domains aliases", func(t *testing.T) {

		mockReturn := []*domain.DomainAlias{{Domain: "alias.com"}}
		expected := []model.DomainAlias{{Domain: "alias.com"}}

		mockMailSrvUseCase.EXPECT().
			GetDomainsAliases(gomock.Any()).
			Times(1).
			Return(mockReturn, nil)
		response := testutil.NewRequest().Get("/domains-aliases").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusOK, response.Code())

		var result []model.DomainAlias
		err = response.UnmarshalBodyToObject(&result)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, expected, result)
	})

	t.Run("Sucessfully get all domain aliases filtered by domain", func(t *testing.T) {

		filter := "domain.com"
		mockReturn := []*domain.DomainAlias{{Alias: "email@alias.com", Domain: "email@domain.com"}}
		expected := []model.DomainAlias{{Alias: "email@alias.com", Domain: "email@domain.com"}}

		mockMailSrvUseCase.EXPECT().
			GetDomainsAliasesByDomain(gomock.Any(), filter).
			Times(1).
			Return(mockReturn, nil)
		response := testutil.NewRequest().Get(fmt.Sprintf("/domains-aliases/filterByDomain/%s", filter)).
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusOK, response.Code())

		var result []model.DomainAlias
		err = response.UnmarshalBodyToObject(&result)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, expected, result)
	})

	t.Run("Sucessfully creates an address alias", func(t *testing.T) {
		alias := "email@alias.com"
		addresses := "email@domain.com"
		param := model.AddressAlias{Alias: alias, Addresses: addresses}
		expected := model.AddressAlias{Alias: alias, Addresses: addresses}

		mockMailSrvUseCase.EXPECT().
			CreateAddressAlias(gomock.Any(), alias, addresses).
			Times(1).
			Return(nil)
		response := testutil.NewRequest().Post("/addresses-aliases").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			WithJsonBody(param).Go(t, e)

		require.Equal(t, http.StatusCreated, response.Code())

		var result model.AddressAlias
		err = response.UnmarshalBodyToObject(&result)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, expected, result)

	})

	t.Run("Get a Bad Request Status when try creates an address alias with wrong format at addresses list", func(t *testing.T) {
		alias := "alias@example.com"
		addresses := "email@domain.com,email2@domain.com email3@domain.com" // there is a missing comma between email2 and email3
		param := model.AddressAlias{Alias: alias, Addresses: addresses}

		mockMailSrvUseCase.EXPECT().
			CreateAddressAlias(gomock.Any(), alias, addresses).
			Times(0)
		response := testutil.NewRequest().Post("/addresses-aliases").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			WithJsonBody(param).Go(t, e)

		require.Equal(t, http.StatusBadRequest, response.Code())

	})

	t.Run("Sucessfully delete an address alias", func(t *testing.T) {
		alias := "email@alias.com"

		mockMailSrvUseCase.EXPECT().
			DeleteAddressAlias(gomock.Any(), alias).
			Times(1).
			Return(nil)
		response := testutil.NewRequest().Delete(fmt.Sprintf("/addresses-aliases/%s", alias)).
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusNoContent, response.Code())
	})

	t.Run("Sucessfully update an address alias", func(t *testing.T) {
		alias := "email@alias.com"
		addresses := "email@domain.com"
		param := model.AddressAlias{Alias: alias, Addresses: addresses}
		mockRet := &domain.AddressAlias{Alias: alias, Addresses: addresses}
		expected := param

		mockMailSrvUseCase.EXPECT().
			UpdateAddressAlias(gomock.Any(), *mockRet).
			Times(1).
			Return(mockRet, nil)
		response := testutil.NewRequest().Put(fmt.Sprintf("/addresses-aliases/%s", alias)).
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			WithJsonBody(param).
			Go(t, e)

		require.Equal(t, http.StatusOK, response.Code())

		var result model.AddressAlias
		err = response.UnmarshalBodyToObject(&result)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, expected, result)
	})

	t.Run("Sucessfully get all addresses aliases", func(t *testing.T) {

		mockReturn := []*domain.AddressAlias{{Alias: "email@alias.com", Addresses: "email@domain.com"}}
		expected := []model.AddressAlias{{Alias: "email@alias.com", Addresses: "email@domain.com"}}

		mockMailSrvUseCase.EXPECT().
			GetAddressesAliases(gomock.Any()).
			Times(1).
			Return(mockReturn, nil)
		response := testutil.NewRequest().Get("/addresses-aliases").
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusOK, response.Code())

		var result []model.AddressAlias
		err = response.UnmarshalBodyToObject(&result)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, expected, result)
	})

	t.Run("Sucessfully get all addresses aliases filtered by domain", func(t *testing.T) {

		filter := "domain.com"
		mockReturn := []*domain.AddressAlias{{Alias: "email@alias.com", Addresses: "email@domain.com"}}
		expected := []model.AddressAlias{{Alias: "email@alias.com", Addresses: "email@domain.com"}}

		mockMailSrvUseCase.EXPECT().
			GetAddressesAliasesByDomain(gomock.Any(), filter).
			Times(1).
			Return(mockReturn, nil)
		response := testutil.NewRequest().Get(fmt.Sprintf("/addresses-aliases/filterByDomain/%s", filter)).
			WithJWSAuth(string(jwt)).
			WithAcceptJson().
			Go(t, e)

		require.Equal(t, http.StatusOK, response.Code())

		var result []model.AddressAlias
		err = response.UnmarshalBodyToObject(&result)
		assert.NoError(t, err, "error unmarshaling response")

		require.Equal(t, expected, result)
	})
}

// func domainAccountToModel(account domain.Account) model.Account {
// 	return model.Account{Name: account.Domain, Domain: account.Domain}
// }

func modelAccountToDomain(account model.Account) domain.Account {
	return domain.Account{Name: account.Domain, Domain: account.Domain}
}
