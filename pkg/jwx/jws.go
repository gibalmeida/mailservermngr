package jwx

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/ecdsafile"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
)

const KeyID = `KID`
const JWSIssuer = "MailServer-Server"
const JWSAudience = "MailServer-Users"
const PermissionsClaim = "perm"

type JWSAuthenticator struct {
	PrivateKey *ecdsa.PrivateKey
	KeySet     jwk.Set
}

var _ JWSValidator = (*JWSAuthenticator)(nil)

// NewJWSAuthenticator creates an authenticator example which uses a hard coded
// ECDSA key to validate JWT's that it has signed itself.
func NewJWSAuthenticator(privateKey []byte) (*JWSAuthenticator, error) {

	privKey, err := ecdsafile.LoadEcdsaPrivateKey([]byte(privateKey))
	if err != nil {
		return nil, fmt.Errorf("loading PEM private key: %w", err)
	}

	set := jwk.NewSet()
	pubKey := jwk.NewECDSAPublicKey()

	err = pubKey.FromRaw(&privKey.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("parsing jwk key: %w", err)
	}

	err = pubKey.Set(jwk.AlgorithmKey, jwa.ES256)
	if err != nil {
		return nil, fmt.Errorf("setting key algorithm: %w", err)
	}

	err = pubKey.Set(jwk.KeyIDKey, KeyID)
	if err != nil {
		return nil, fmt.Errorf("setting key ID: %w", err)
	}

	set.Add(pubKey)

	return &JWSAuthenticator{PrivateKey: privKey, KeySet: set}, nil
}

// ValidateJWS ensures that the critical JWT claims needed to ensure that we
// trust the JWT are present and with the correct values.
func (f *JWSAuthenticator) ValidateJWS(jwsString string) (jwt.Token, error) {
	return jwt.Parse([]byte(jwsString), jwt.WithKeySet(f.KeySet),
		jwt.WithAudience(JWSAudience), jwt.WithIssuer(JWSIssuer), jwt.WithValidate(true))
}

// SignToken takes a JWT and signs it with our private key, returning a JWS.
func (f *JWSAuthenticator) SignToken(t jwt.Token) ([]byte, error) {
	hdr := jws.NewHeaders()
	if err := hdr.Set(jws.AlgorithmKey, jwa.ES256); err != nil {
		return nil, fmt.Errorf("setting algorithm: %w", err)
	}
	if err := hdr.Set(jws.TypeKey, "JWT"); err != nil {
		return nil, fmt.Errorf("setting type: %w", err)
	}
	if err := hdr.Set(jws.KeyIDKey, KeyID); err != nil {
		return nil, fmt.Errorf("setting Key ID: %w", err)
	}
	return jwt.Sign(t, jwa.ES256, f.PrivateKey, jwt.WithHeaders(hdr))
}

// CreateJWSWithClaims is a helper function to create JWT's with the specified
// claims.
func (f *JWSAuthenticator) CreateJWSWithClaims(claims []string) ([]byte, error) {
	t := jwt.New()
	err := t.Set(jwt.IssuerKey, JWSIssuer)
	if err != nil {
		return nil, fmt.Errorf("setting issuer: %w", err)
	}
	err = t.Set(jwt.AudienceKey, JWSAudience)
	if err != nil {
		return nil, fmt.Errorf("setting audience: %w", err)
	}
	now := time.Unix(time.Now().Unix(), 0)
	err = t.Set(jwt.IssuedAtKey, now.Unix())
	if err != nil {
		return nil, fmt.Errorf("setting issued at: %w", err)
	}
	err = t.Set(jwt.NotBeforeKey, now.Unix())
	if err != nil {
		return nil, fmt.Errorf("setting not before at: %w", err)
	}
	err = t.Set(jwt.ExpirationKey, now.Add(8*time.Hour).Unix())
	if err != nil {
		return nil, fmt.Errorf("setting expiration at: %w", err)
	}
	err = t.Set(PermissionsClaim, claims)
	if err != nil {
		return nil, fmt.Errorf("setting permissions: %w", err)
	}
	return f.SignToken(t)
}
