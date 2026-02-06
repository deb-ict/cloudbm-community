package authentication

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Identity interface {
	GetID() string
	GetUsername() string
	GetEmail() string
	IsEmailVerified() bool

	GetClaim(key string) (any, bool)

	HasRole(role string) bool
	HasAnyRole(roles ...string) bool
	HasAllRoles(roles ...string) bool

	HasScope(scope string) bool
	HasAnyScope(scopes ...string) bool
	HasAllScopes(scopes ...string) bool
}

type identity struct {
	id            string
	username      string
	email         string
	emailVerified bool
	roles         []string
	scopes        []string
	IssuedAt      time.Time
	ExpiresAt     time.Time
	NotBefore     time.Time
	claims        jwt.MapClaims
}

func (i *identity) GetID() string {
	return i.id
}

func (i *identity) GetUsername() string {
	return i.username
}

func (i *identity) GetEmail() string {
	return i.email
}

func (i *identity) IsEmailVerified() bool {
	return i.emailVerified
}

func (i *identity) GetClaim(key string) (any, bool) {
	value, ok := i.claims[key]
	return value, ok
}

func (i *identity) HasRole(role string) bool {
	for _, r := range i.roles {
		if r == role {
			return true
		}
	}
	return false
}

func (i *identity) HasAnyRole(roles ...string) bool {
	for _, role := range roles {
		if i.HasRole(role) {
			return true
		}
	}
	return false
}

func (i *identity) HasAllRoles(roles ...string) bool {
	for _, role := range roles {
		if !i.HasRole(role) {
			return false
		}
	}
	return true
}

func (i *identity) HasScope(scope string) bool {
	for _, s := range i.scopes {
		if s == scope {
			return true
		}
	}
	return false
}

func (i *identity) HasAnyScope(scopes ...string) bool {
	for _, scope := range scopes {
		if i.HasScope(scope) {
			return true
		}
	}
	return false
}

func (i *identity) HasAllScopes(scopes ...string) bool {
	for _, scope := range scopes {
		if !i.HasScope(scope) {
			return false
		}
	}
	return true
}

func newIdentityFromClaims(claims jwt.MapClaims) (Identity, error) {
	id, _ := claims["sub"].(string)
	username, _ := claims["name"].(string)
	email, _ := claims["email"].(string)
	emailVerified, _ := claims["email_verified"].(bool)

	var roles []string
	if rolesClaim, ok := claims["role"]; ok {
		switch v := rolesClaim.(type) {
		case string:
			roles = strings.Fields(v)
		case []interface{}:
			for _, r := range v {
				if roleStr, ok := r.(string); ok {
					roles = append(roles, roleStr)
				}
			}
		}
	}

	var scopes []string
	if scopesClaim, ok := claims["scope"]; ok {
		switch v := scopesClaim.(type) {
		case string:
			scopes = strings.Fields(v)
		case []interface{}:
			for _, s := range v {
				if scopeStr, ok := s.(string); ok {
					scopes = append(scopes, scopeStr)
				}
			}
		}
	}

	return &identity{
		id:            id,
		username:      username,
		email:         email,
		emailVerified: emailVerified,
		roles:         roles,
		scopes:        scopes,
		claims:        claims,
	}, nil
}
