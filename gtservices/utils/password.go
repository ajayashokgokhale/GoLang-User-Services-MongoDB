package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/crypto/argon2"
)

// Password hashing constants for Argon2.
const (
	saltLength  = 16
	memory      = 64 * 1024 // 64 MB
	iterations  = 4
	parallelism = 1
	keyLength   = 32
)

func CreateHashPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("password cannot be empty")
	}

	// Generate random salt.
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	hash := argon2.IDKey([]byte(password), salt, iterations, memory, uint8(parallelism), keyLength)

	// Encode parameters, salt, and hash into a single string.
	encoded := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		memory, iterations, parallelism,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash))

	return encoded, nil
}

// IsValidPassword checks if a password meets security requirements:
// - At least 6 characters.
// - Includes uppercase and lowercase letters, a number, and a special character (!, @, #, $, %, ^).
// Returns true if all criteria are met, false otherwise.
func IsValidPassword(password string) bool {
	if len(password) < 6 {
		return false
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasNumber = true
		case strings.ContainsRune("!@#$%^", ch):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}

// VerifyPassword checks if a plain password matches an Argon2id encoded hash.
// The hash must be in the format $argon2id$v=19$m=...,t=...,p=...$salt$hash.
// Returns true if the password matches, false otherwise, and an error if the hash is invalid.
func VerifyPassword(password, encodedHash string) (bool, error) {
	if password == "" {
		return false, fmt.Errorf("password cannot be empty")
	}
	if encodedHash == "" {
		return false, fmt.Errorf("encoded hash cannot be empty")
	}

	// Parse hash components.
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 || parts[1] != "argon2id" || parts[2] != "v=19" {
		return false, fmt.Errorf("invalid hash format")
	}

	// Extract parameters.
	var memory, time uint32
	var threads uint8
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return false, fmt.Errorf("failed to parse hash parameters: %w", err)
	}

	// Decode salt and hash.
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %w", err)
	}
	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, fmt.Errorf("failed to decode hash: %w", err)
	}

	// Compute hash for comparison.
	computedHash := argon2.IDKey([]byte(password), salt, time, memory, threads, uint32(len(hash)))

	// Use constant-time comparison to prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, computedHash) == 1 {
		return true, nil
	}
	return false, nil
}
