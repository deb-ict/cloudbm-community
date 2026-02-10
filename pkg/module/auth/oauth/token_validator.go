package oauth

import (
	"strconv"
	"strings"

	"github.com/deb-ict/go-router/authentication"
	"github.com/golang-jwt/jwt/v5"
)

type jwtValidator struct {
}

func NewTokenValidator() authentication.BearerAuthenticationValidator {
	return &jwtValidator{}
}

func (v *jwtValidator) GetBearerAuthenticationData(token string) (authentication.ClaimMap, error) {
	jwtClaims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, jwtClaims, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return tokenSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	claims := make(authentication.ClaimMap)
	claims["jti"] = &authentication.Claim{
		Name:   "jti",
		Values: []string{jwtClaims["jti"].(string)},
	}
	claims["iss"] = &authentication.Claim{
		Name:   "iss",
		Values: []string{jwtClaims["iss"].(string)},
	}
	claims["aud"] = &authentication.Claim{
		Name:   "aud",
		Values: []string{jwtClaims["aud"].(string)},
	}
	claims["sub"] = &authentication.Claim{
		Name:   "sub",
		Values: []string{jwtClaims["sub"].(string)},
	}
	claims["name"] = &authentication.Claim{
		Name:   "name",
		Values: []string{jwtClaims["name"].(string)},
	}
	claims["email"] = &authentication.Claim{
		Name:   "email",
		Values: []string{jwtClaims["email"].(string)},
	}
	claims["email_verified"] = &authentication.Claim{
		Name:   "email_verified",
		Values: []string{strconv.FormatBool(jwtClaims["email_verified"].(bool))},
	}
	claims["phone"] = &authentication.Claim{
		Name:   "phone",
		Values: []string{jwtClaims["phone"].(string)},
	}
	claims["phone_verified"] = &authentication.Claim{
		Name:   "phone_verified",
		Values: []string{strconv.FormatBool(jwtClaims["phone_verified"].(bool))},
	}
	claims["role"] = &authentication.Claim{
		Name:   "role",
		Values: strings.Split(jwtClaims["role"].(string), " "),
	}
	claims["scope"] = &authentication.Claim{
		Name:   "scope",
		Values: strings.Split(jwtClaims["scope"].(string), " "),
	}

	return claims, nil
}
