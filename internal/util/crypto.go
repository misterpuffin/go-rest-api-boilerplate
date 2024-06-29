package util

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
)

// Generate 16 bytes randomly and securely using the
// Cryptographically secure pseudorandom number generator (CSPRNG)
// in the crypto.rand package
func GenerateRandomSalt(saltSize int) string {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		// TODO: Remove Panic
		panic(err)
	}

	return hex.EncodeToString(salt)
}

func HashPassword(password string, salt string) string {
	var passwordBytes = []byte(password)
	var saltBytes = []byte(salt)
	var sha512Hasher = sha512.New()

	passwordBytes = append(passwordBytes, saltBytes...)
	sha512Hasher.Write(passwordBytes)

	var hashedPasswordsBytes = sha512Hasher.Sum(nil)
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordsBytes)

	return hashedPasswordHex
}
