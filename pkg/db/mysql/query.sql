-- name: GetAccount :one
SELECT name, domain, home_dir, quota FROM account 
WHERE name = ? AND domain = ?
LIMIT 1;

-- name: CreateAccount :execresult
INSERT INTO account 
(name,domain,password,home_dir,quota)
VALUES (?,?,?,?,?);

-- name: UpdateAccountPassword :exec
UPDATE account 
SET password = ?
WHERE name = ? AND domain = ?;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE name = ? AND domain = ?;

-- name: GetAccounts :many
SELECT name, domain FROM account
ORDER BY domain, name;

-- name: GetAccountsFilteredByDomain :many
SELECT name, domain FROM account
WHERE domain = ? ORDER BY name;

-- name: CreateDomain :execresult
INSERT INTO domain 
(domain)
VALUES (?);

-- name: GetDomain :one
SELECT * FROM domain
WHERE domain = ?
LIMIT 1;

-- name: DeleteDomain :exec
DELETE FROM domain
WHERE domain = ?;

-- name: GetDomains :many
SELECT * FROM domain
ORDER BY domain;

-- name: CreateAddressAlias :execresult
INSERT INTO address_alias 
(alias,addresses)
VALUES (?,?);

-- name: GetAddressAlias :one
SELECT * FROM address_alias
WHERE alias = ?
LIMIT 1;

-- name: GetAddressesAliases :many
SELECT * FROM address_alias
ORDER BY alias;

-- name: GetAddressesAliasesFilteredByDomain :many
SELECT * FROM address_alias
WHERE alias LIKE sqlc.arg('domain')
ORDER BY alias;

-- name: UpdateAddressAlias :exec
UPDATE address_alias
SET addresses = ?
WHERE alias = ?;

-- name: DeleteAddressAlias :exec
DELETE FROM address_alias
WHERE alias = ?;

-- name: CreateDomainAlias :execresult
INSERT INTO domain_alias 
(alias,domain)
VALUES (?,?);

-- name: GetDomainAlias :one
SELECT * FROM domain_alias 
WHERE alias = ? LIMIT 1;

-- name: GetDomainsAliases :many
SELECT * FROM domain_alias
ORDER BY alias;

-- name: GetDomainsAliasesFilteredByDomain :many
SELECT alias, domain FROM domain_alias
WHERE domain = ? ORDER BY alias;

-- name: UpdateDomainAlias :exec
UPDATE domain_alias
SET domain = ?
WHERE alias = ?;

-- name: DeleteDomainAlias :exec
DELETE FROM domain_alias 
WHERE alias = ?;


