package hash

import (
	"crypto/sha1"
	"fmt"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type SHA1Hasher struct {
	salt string
}

func NewSHA1Hasher(salt string) *SHA1Hasher {
	return &SHA1Hasher{salt: salt}
}

func (hash *SHA1Hasher) Hash(password string) (string, error) {
	newHash := sha1.New()

	if _, err := newHash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", newHash.Sum([]byte(hash.salt))), nil
}
