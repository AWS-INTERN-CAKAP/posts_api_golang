package token

import (
	"errors"
	"sync"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUseCase interface {
	GenerateAccessToken(claims JwtCustomClaims) (string, error)
	IsTokenBlacklisted(tokenString string) bool
	InvalidateToken(tokenString string) error
}

type tokenUseCase struct {
	secretKey string
	blacklist          map[string]bool
	mu                 sync.Mutex
}

type JwtCustomClaims struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Nama   string `json:"nama"`
	jwt.RegisteredClaims
}

func NewTokenUseCase(secretKey string) *tokenUseCase {
	return &tokenUseCase{
		secretKey: secretKey,
		blacklist:          make(map[string]bool),
	}
}

func (t *tokenUseCase) GenerateAccessToken(claims JwtCustomClaims) (string, error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := plainToken.SignedString([]byte(t.secretKey))

	if err != nil {
		return "", err
	}

	return encodedToken, nil
}

func (t *tokenUseCase) GetClaimsFromToken(tokenString string) (JwtCustomClaims, error) {
	claims := JwtCustomClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secretKey), nil
	})
	if err != nil {
		return JwtCustomClaims{}, err
	}
	return claims, nil
}

func (t *tokenUseCase) InvalidateToken(tokenString string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	_, err := t.GetClaimsFromToken(tokenString)
	if err != nil {
		return errors.New("failed to parse token, cannot invalidate")
	}

	t.blacklist[tokenString] = true
	return nil
}

func (t *tokenUseCase) IsTokenBlacklisted(tokenString string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.blacklist[tokenString]
}