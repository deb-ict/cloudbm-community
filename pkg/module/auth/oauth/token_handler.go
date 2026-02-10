package oauth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/deb-ict/cloudbm-community/pkg/module/auth"
	"github.com/deb-ict/cloudbm-community/pkg/module/auth/model"
	"github.com/deb-ict/go-router"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const TokenLifetimeSeconds = 3600

var tokenSecret = []byte("your-256-bit-secret")

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
	//example parameter
}

type ErrorResponse struct {
	Error       string `json:"error"`
	Description string `json:"error_description,omitempty"`
	HelpUri     string `json:"error_uri,omitempty"`
}

func (t *TokenResponse) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(t)
}

func (e *ErrorResponse) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")
	w.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(w).Encode(e)
}

type tokenHandler struct {
	service   auth.Service
	issuerUri string
}

func (api *tokenHandler) RegisterRoutes(r *router.Router) {
	r.HandleFunc("/oauth/token", api.TokenEndpoint,
		router.AllowedMethod(http.MethodPost),
	)
}

func (h *tokenHandler) TokenEndpoint(w http.ResponseWriter, r *http.Request) {
	// Validate the method
	if r.Method != http.MethodPost {
		h.tokenHandlerError(w, "invalid_request")
		return
	}

	// Validate the content type
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		h.tokenHandlerError(w, "invalid_request")
		return
	}

	// Parse the form
	if r.Form == nil {
		r.ParseForm()
	}

	// Get the client
	clientId, clientSecret, useBasicAuth := r.BasicAuth()
	if !useBasicAuth {
		clientIdParam := r.Form["client_id"]
		if len(clientIdParam) != 1 {
			h.tokenHandlerError(w, "invalid_request")
			return
		}
		clientId = clientIdParam[0]

		clientSecretParam := r.Form["client_secret"]
		if len(clientSecretParam) > 1 {
			h.tokenHandlerError(w, "invalid_request")
			return
		}
		clientSecret = clientSecretParam[0]
	}
	//TODO: This should not be hardcoded!
	if clientId != "cloudbm" || clientSecret != "XX0rQ0zgD2MHZ2KdwzDi" {
		h.tokenHandlerError(w, "unauthorized_client")
		return
	}

	// Get the grant type
	grantType := r.Form["grant_type"]
	if len(grantType) != 1 {
		h.tokenHandlerError(w, "invalid_request")
		return
	}

	// Handle the grant type
	switch grantType[0] {
	case "password":
		h.passwordTokenHandler(w, r)
	default:
		h.tokenHandlerError(w, "unsupported_grant_type")
	}
}

func (h *tokenHandler) passwordTokenHandler(w http.ResponseWriter, r *http.Request) {
	usernameParam := r.Form["username"]
	if len(usernameParam) != 1 {
		h.tokenHandlerError(w, "invalid_request")
		return
	}
	username := usernameParam[0]

	passwordParam := r.Form["password"]
	if len(passwordParam) != 1 {
		h.tokenHandlerError(w, "invalid_request")
		return
	}
	password := passwordParam[0]

	user, err := h.service.GetUserByUsername(r.Context(), username)
	if err != nil {
		h.tokenHandlerError(w, "access_denied")
		return
	}
	if user == nil {
		h.tokenHandlerError(w, "access_denied")
		return
	}

	if !user.VerifyPassword(h.service.PasswordHasher(), password) {
		h.tokenHandlerError(w, "access_denied")
		return
	}

	tokenString, err := h.generateJwtToken(user)
	if err != nil {
		h.tokenHandlerError(w, "server_error")
		return
	}

	response := &TokenResponse{
		AccessToken: tokenString,
		TokenType:   "bearer",
		ExpiresIn:   TokenLifetimeSeconds,
	}
	response.Send(w)
}

func (h *tokenHandler) tokenHandlerError(w http.ResponseWriter, e string) {
	errorResponse := &ErrorResponse{
		Error: e,
	}
	errorResponse.Send(w)
}

func (h *tokenHandler) generateJwtToken(user *model.User) (string, error) {
	tokenId := uuid.New().String()

	claims := jwt.MapClaims{}
	claims["jti"] = tokenId
	claims["iss"] = h.issuerUri
	claims["aud"] = "cloudbm"
	claims["iat"] = jwt.NewNumericDate(time.Now().UTC())
	claims["nbf"] = jwt.NewNumericDate(time.Now().UTC())
	claims["exp"] = jwt.NewNumericDate(time.Now().UTC().Add(time.Duration(TokenLifetimeSeconds) * time.Second))
	claims["sub"] = user.Id
	claims["name"] = user.Username
	claims["email"] = user.Email
	claims["email_verified"] = user.EmailVerified
	claims["phone"] = user.Phone
	claims["phone_verified"] = user.PhoneVerified
	claims["role"] = "user admin"
	claims["scope"] = "product:read product:write"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(tokenSecret)
	if err != nil {
		return "", err
	}

	//TODO: We should store the token in the database, so we can revoke it if needed

	//	/token/introspect
	//	/token
	//	/token/revoke
	//	/userinfo
	//	/logout

	return tokenString, nil
}
