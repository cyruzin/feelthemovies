package controllers

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/cyruzin/feelthemovies/internal/app/config"
	"github.com/cyruzin/feelthemovies/internal/pkg/errhandler"
	"github.com/cyruzin/feelthemovies/internal/pkg/ratelimit"
	jwt "github.com/dgrijalva/jwt-go"
)

// LoggerMiddleware logs the details of all requests.
func (s *Setup) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		s.logger.Infow(
			"request_logging",
			"method", r.Method,
			"url", r.URL.String(),
			"agent", r.UserAgent(),
			"referer", r.Referer(),
			"proto", r.Proto,
			"remote_address", r.RemoteAddr,
			"latency", time.Since(start),
		)

		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware checks if the request contains Bearer Token
// on the headers and if it is valid.
func (s *Setup) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Capturing Authorizathion header.
		tokenHeader := r.Header.Get("Authorization")

		// Checking if the value is empty.
		if tokenHeader == "" {
			errhandler.DecodeError(w, r, s.logger, errEmptyToken, http.StatusBadRequest)
			return
		}

		// Checking if the header contains Bearer string and if the token exists.
		if !strings.Contains(tokenHeader, "Bearer") || len(strings.Split(tokenHeader, "Bearer ")) == 1 {
			errhandler.DecodeError(w, r, s.logger, errMalformedToken, http.StatusUnauthorized)
			return
		}

		// Capturing the token.
		jwtString := strings.Split(tokenHeader, "Bearer ")[1]

		// Parsing the token to verify its authenticity.
		token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			cfg, err := config.Load()
			if err != nil {
				return nil, err
			}
			return []byte(cfg.JWTSecret), nil
		})

		// Returning parsing errors.
		if err != nil {
			errhandler.DecodeError(w, r, s.logger, err.Error(), http.StatusUnauthorized)
			return
		}

		// If the token is valid.
		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			errhandler.DecodeError(w, r, s.logger, errInvalidJWTToken, http.StatusUnauthorized)
			return
		}
	})
}

// RateLimit middleware handles the rate limiting.
func (s *Setup) RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			errhandler.DecodeError(w, r, s.logger, errInternal, http.StatusInternalServerError)
			return
		}

		limiter := ratelimit.GetVisitor(ip)
		if !limiter.Allow() {
			errhandler.DecodeError(
				w,
				r,
				s.logger,
				http.StatusText(http.StatusTooManyRequests),
				http.StatusTooManyRequests,
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}
