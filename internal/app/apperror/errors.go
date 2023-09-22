package apperror

type AppError interface {
	Error() string
	IsServiceError() bool
}

type appError struct {
	s string
}

func NewAppError(text string) *appError {
	return &appError{text}
}

func (e *appError) Error() string {
	return e.s
}

func (e *appError) IsServiceError() bool {
	return true
}

var ErrAccountAlreadyExist = NewAppError("Account already exist")
var ErrAccountNotExist = NewAppError("Account doesn't exist")
var ErrAddressAliasAlreadyExist = NewAppError("Address alias already exist")
var ErrAddressAliasNotExist = NewAppError("Address Alias doesn't exist")
var ErrAddressAliasAlreadyExistAsRegular = NewAppError("Email address already registred as a regular address")
var ErrDomainAlreadyExist = NewAppError("Domain already exist")
var ErrDomainNotExist = NewAppError("Domain doesn't exist")
var ErrDomainAliasAlreadyExist = NewAppError("Domain alias already exist")
var ErrDomainAliasNotExist = NewAppError("Domain Alias doesn't exist")
var ErrDomainOrDomainAliasAlreadyExist = NewAppError("Domain or Domain alias already registered")
var ErrDomainOrDomainAliasNotExist = NewAppError("Domain or Domain alias doesn't exist")
