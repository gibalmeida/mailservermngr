package domain

// Account defines model for Account.
type Account struct {
	Domain string `json:"domain"`
	Name   string `json:"name"`
}

// Account defines model for a New Account.
type NewAccount struct {
	Domain   string `json:"domain"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// AddressAlias defines model for Address Alias.
type AddressAlias struct {
	Addresses []EmailAddress `json:"addresses"`
	Alias     EmailAddress   `json:"alias"`
}

// Domain defines a model for an Internet Domain (ex: example.com)
type Domain struct {
	// Domain Domínio
	Domain string `json:"domain"`
}

// DomainAlias defines a model for an Internet Domain Alias (ex: @domain-alias.com)
type DomainAlias struct {
	Alias  string `json:"alias"`
	Domain string `json:"domain"`
}

// EmailAddress defines model for EmailAddress.
type EmailAddress = string
