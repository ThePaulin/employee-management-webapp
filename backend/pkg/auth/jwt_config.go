package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenConfig interface {
	NewJWT(userId string, ttl time.Duration) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type Config struct {
	signingKey string
}

func NewConfig(signingKey string) (*Config, error) {
	if signingKey == "" {
		return nil, errors.New("no signing key provided")
	}

	return &Config{signingKey: signingKey}, nil
}

func (config *Config) NewJWT(userId string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   userId,
		ExpiresAt: time.Now().Add(ttl).Unix(),
	})

	return token.SignedString([]byte(config.signingKey))
}

func (config *Config) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.signingKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error geting user claims from token")
	}

	return claims["sub"].(string), nil
}

func (config *Config) NewRefreshToken() (string, error) {
	b32 := make([]byte, 32)

	source := rand.NewSource(time.Now().Unix())
	random := rand.New(source)

	if _, err := random.Read(b32); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b32), nil
}
