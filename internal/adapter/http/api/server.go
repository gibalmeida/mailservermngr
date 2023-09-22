package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/oapi-codegen/echo-middleware"

	"github.com/gibalmeida/mailservermngr/internal/adapter/http/api/model"
	"github.com/gibalmeida/mailservermngr/internal/app/apperror"
	"github.com/gibalmeida/mailservermngr/internal/domain"
	"github.com/gibalmeida/mailservermngr/internal/port"

	"github.com/gibalmeida/mailservermngr/pkg/jwx"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=types.cfg.yaml ./opendomain.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml ./opendomain.yaml

type Server struct {
	authUseCase       port.AuthUseCase
	mailServerUseCase port.MailServerUseCase
}

func NewServer(authUseCase port.AuthUseCase, mailServerUseCase port.MailServerUseCase) *Server {
	return &Server{
		authUseCase:       authUseCase,
		mailServerUseCase: mailServerUseCase,
	}
}

func CreateMiddleware(v jwx.JWSValidator) ([]echo.MiddlewareFunc, error) {
	spec, err := GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("Error loading swagger spec: %w", err)
	}

	validator := echomiddleware.OapiRequestValidatorWithOptions(spec,
		&echomiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: jwx.NewJWTAuthenticator(v),
			},
		})

	return []echo.MiddlewareFunc{validator}, nil
}

// sendMailServerError wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendMailServerError(ctx echo.Context, code int, message string) error {
	mailServerErr := model.Error{
		Code:    int32(code),
		Message: message,
	}
	return ctx.JSON(code, mailServerErr)
}

func sendBadRequestOrInternalServerError(ctx echo.Context, err error) error {
	if _, ok := err.(apperror.AppError); ok {
		return sendMailServerError(ctx, http.StatusBadRequest, err.Error())
	}

	log.Println(err)
	return sendMailServerError(ctx, http.StatusInternalServerError, "Internal error")
}

func sendBadRequestError(ctx echo.Context, message string) error {
	return sendMailServerError(ctx, http.StatusBadRequest, message)
}

// (GET /accounts)
func (m *Server) GetAccounts(ctx echo.Context) error {

	var accounts []*domain.Account

	accounts, err := m.mailServerUseCase.GetAccounts(ctx.Request().Context())

	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, accounts)
}

// (POST /accounts)
func (m *Server) CreateAccount(ctx echo.Context) error {

	var newAccount model.NewAccount

	// here (and in all handlers) we ignore the error that can be returned by ctx.Bind because we are using
	// middleware.OapiRequestValidatorWithOptions to validate the requests
	ctx.Bind(&newAccount)

	err := m.mailServerUseCase.CreateAccount(ctx.Request().Context(),
		domain.NewAccount{
			Name:     newAccount.Name,
			Domain:   newAccount.Domain,
			Password: newAccount.Password,
		})

	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, newAccount)
}

// (GET /accounts/filterByDomain/{domain})
func (m *Server) GetAccountsByDomain(ctx echo.Context, addressDomain string) error {
	var result []*domain.Account

	result, err := m.mailServerUseCase.GetAccountsByDomain(ctx.Request().Context(), addressDomain)

	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, result)
}

// (DELETE /accounts/{emailAddress})
func (m *Server) DeleteAccount(ctx echo.Context, emailAddress model.EmailAddress) error {

	emailAddressSplitted := strings.Split(emailAddress, "@")

	if len(emailAddressSplitted) != 2 {
		return sendBadRequestError(ctx, "invalid email address")
	}

	err := m.mailServerUseCase.DeleteAccount(ctx.Request().Context(), emailAddressSplitted[0], emailAddressSplitted[1])

	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// (GET /accounts/{emailAddress})
func (m *Server) GetAccount(ctx echo.Context, emailAddress model.EmailAddress) error {
	emailAddressSplitted := strings.Split(emailAddress, "@")

	account, err := m.mailServerUseCase.GetAccount(ctx.Request().Context(), emailAddressSplitted[0], emailAddressSplitted[1])

	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, account)
}

// Change account password
// (POST /accounts/changePassword)
func (m *Server) UpdateAccountPassword(ctx echo.Context) error {
	var changeAccountPassword model.ChangeAccountPassword

	ctx.Bind(&changeAccountPassword)

	emailAddressSplitted := strings.Split(changeAccountPassword.EmailAddress, "@")

	if err := m.mailServerUseCase.UpdateAccountPassword(ctx.Request().Context(), emailAddressSplitted[0], emailAddressSplitted[1], changeAccountPassword.NewPassword); err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// (GET /addresses-aliases)
func (m *Server) GetAddressAliases(ctx echo.Context) error {
	var addressesAliases []*domain.AddressAlias

	addressesAliases, err := m.mailServerUseCase.GetAddressesAliases(ctx.Request().Context())

	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, addressesAliases)
}

// (POST /addresses-aliases)
func (m *Server) CreateAddressAlias(ctx echo.Context) error {
	var addressAlias domain.AddressAlias

	ctx.Bind(&addressAlias)

	if err := m.mailServerUseCase.CreateAddressAlias(ctx.Request().Context(), addressAlias.Alias, addressAlias.Addresses); err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, addressAlias)
}

// (GET /addresses-aliases/filterByDomain/{domain})
func (m *Server) GetAddressAliasesByDomain(ctx echo.Context, domain string) error {
	result, err := m.mailServerUseCase.GetAddressesAliasesByDomain(ctx.Request().Context(), domain)

	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, result)
}

// (DELETE /addresses-aliases/{alias})
func (m *Server) DeleteAddressAlias(ctx echo.Context, alias string) error {

	err := m.mailServerUseCase.DeleteAddressAlias(ctx.Request().Context(), alias)
	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// (PUT /addresses-aliases/{alias})
func (m *Server) UpdateAddressAlias(ctx echo.Context, alias string) error {
	var addressAlias domain.AddressAlias

	ctx.Bind(&addressAlias)

	result, err := m.mailServerUseCase.UpdateAddressAlias(ctx.Request().Context(), addressAlias)

	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, *result)
}

// (GET /domains-aliases)
func (m *Server) GetDomainsAliases(ctx echo.Context) error {
	var result []*domain.DomainAlias

	result, err := m.mailServerUseCase.GetDomainsAliases(ctx.Request().Context())

	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, result)
}

// (POST /domains-aliases)
func (m *Server) CreateDomainAlias(ctx echo.Context) error {
	var domainAlias domain.DomainAlias

	ctx.Bind(&domainAlias)

	err := m.mailServerUseCase.CreateDomainAlias(ctx.Request().Context(), domainAlias.Alias, domainAlias.Domain)

	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, domainAlias)
}

// (DELETE /domains-aliases/{domain_alias})
func (m *Server) DeleteDomainAlias(ctx echo.Context, domainAlias string) error {

	err := m.mailServerUseCase.DeleteDomainAlias(ctx.Request().Context(), domainAlias)
	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// (GET /domains)
func (m *Server) GetDomains(ctx echo.Context) error {

	var result []*domain.Domain

	result, err := m.mailServerUseCase.GetDomains(ctx.Request().Context())
	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, result)
}

// (POST /domains)
func (m *Server) CreateDomain(ctx echo.Context) error {

	var domain domain.Domain

	ctx.Bind(&domain)

	err := m.mailServerUseCase.CreateDomain(ctx.Request().Context(), domain.Domain)

	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, domain)
}

// (DELETE /domains/{domain})
func (m *Server) DeleteDomain(ctx echo.Context, domain string) error {

	err := m.mailServerUseCase.DeleteDomain(ctx.Request().Context(), domain)
	if err != nil {
		return sendBadRequestOrInternalServerError(ctx, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// (POST /getToken)
func (m *Server) GetToken(ctx echo.Context) error {
	var auth model.Auth
	var token model.GetTokenResponse
	var err error

	ctx.Bind(&auth)

	token.AccessToken, err = m.authUseCase.GetTokenUsingUsernameAndPassword(ctx.Request().Context(), auth.Username, auth.Password)

	if err != nil {
		return sendMailServerError(ctx, http.StatusForbidden, "Invalid credentials!")
	}

	return ctx.JSON(http.StatusOK, token)
}
