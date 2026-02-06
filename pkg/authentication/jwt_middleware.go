package authentication

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/deb-ict/cloudbm-community/pkg/logging"
	"github.com/golang-jwt/jwt/v5"
)

type UserParserFunc func(token *jwt.Token) (Identity, error)

type JwtAuthenticationMiddleware struct {
	userParser UserParserFunc
}

func NewMiddleware(userParser UserParserFunc) *JwtAuthenticationMiddleware {
	middleware := &JwtAuthenticationMiddleware{
		userParser: userParser,
	}
	middleware.EnsureDefaults()
	return middleware
}

func (m *JwtAuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logging.GetLoggerFromContext(ctx)

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			logger.WarnContext(ctx, "Failed to authorized: Authorization header is missing")
			UnauthorizedResponse(w)
			return
		}

		if !strings.HasPrefix(tokenString, "Bearer ") {
			logger.WarnContext(ctx, "Failed to authorized: Invalid token format")
			UnauthorizedResponse(w)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		if tokenString == "" {
			logger.WarnContext(ctx, "Failed to authorized: Token is missing")
			UnauthorizedResponse(w)
			return
		}

		keyFunc := func(token *jwt.Token) (interface{}, error) {
			hmacMethod, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			if hmacMethod.Hash != jwt.SigningMethodHS256.Hash {
				return nil, fmt.Errorf("unexpected signing method")
			}

			// In a real application, use a secure way to manage your secret key
			secretKey := []byte("your-256-bit-secret")
			return secretKey, nil
		}
		token, err := jwt.Parse(tokenString, keyFunc,
			jwt.WithIssuedAt(),
			jwt.WithNotBeforeRequired(),
			jwt.WithExpirationRequired(),
			jwt.WithIssuer("https://localhost:8000"), //TODO: This should come from configuration
			jwt.WithAudience("cloudbm-users"),        //TODO: This should come from configuration
		)
		if err != nil {
			logger.ErrorContext(ctx, "Failed to authorized: Failed to parse token",
				slog.Any("error", err),
			)
			UnauthorizedResponse(w)
			return
		}
		if !token.Valid {
			logger.WarnContext(ctx, "Failed to authorized: Invalid token")
			UnauthorizedResponse(w)
			return
		}

		identity, err := m.userParser(token)
		if err != nil {
			logger.ErrorContext(ctx, "Failed to authorized: Failed to parse token claims",
				slog.Any("error", err),
			)
			UnauthorizedResponse(w)
			return
		}

		logger = logger.With("user.id", identity.GetId())
		logger.InfoContext(ctx, "User authenticated")

		// Call the next handler in the chain
		ctx = logging.WithLoggerInContext(ctx, logger)
		ctx = WithIdentityInContext(ctx, identity)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *JwtAuthenticationMiddleware) EnsureDefaults() {
	if m.userParser == nil {
		m.userParser = func(token *jwt.Token) (Identity, error) {
			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok {
				return nil, fmt.Errorf("invalid token claims")
			}
			return newIdentityFromClaims(claims)
		}
	}
}

func UnauthorizedResponse(w http.ResponseWriter) {
	http.Error(w, "Access denied", http.StatusUnauthorized)
}
