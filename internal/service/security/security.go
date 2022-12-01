package security

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/argon2"
)

// Hash derives argon2 key (hash) from the password.
// Encodes to base64.
func Hash(password string, salt string) string {
	passwordByte := []byte(password)
	saltByte := []byte(salt)
	keyByte := argon2.IDKey(passwordByte, saltByte, 1, 64*1024, 4, 32)

	return base64.URLEncoding.EncodeToString(keyByte)
}

// IsSameHash safely checks whether the hash was derived from the password.
func IsSameHash(password, hash, salt string) bool {
	newHash := Hash(password, salt)

	return subtle.ConstantTimeCompare([]byte(newHash), []byte(hash)) == 1
}

// GenerateSalt generates random 32 byte salt.
func GenerateSalt() (string, error) {
	salt := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", fmt.Errorf("salt generation: %w", err)
	}

	return base64.URLEncoding.EncodeToString(salt), nil
}

// GenerateSessionID generates base64 random 256-byte session identifier.
func GenerateSessionID() (string, error) {
	b := make([]byte, 256)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", fmt.Errorf("session id generation: %w", err)
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

// IsSameSessionID safely checks whether expected and actual sessionIDs are the same.
func IsSameSessionID(expectedID, actualID string) bool {
	return subtle.ConstantTimeCompare([]byte(expectedID), []byte(actualID)) == 1
}
