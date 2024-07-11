package util

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/misterpuffin/go-rest-api-boilerplate/internal/config"
)

type JWTTokenPayload struct {
	UserId string
}
type JWTToken struct {
	UserId string
	jwt.StandardClaims
}

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

func CreateJWTToken(payload JWTTokenPayload, config config.Config) (token string, err error) {
	claims := JWTToken{
		UserId: payload.UserId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(config.JWT.HoursToExpire)).Unix(),
		},
	}
	tokenGenerator := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signingKey, err := base64.StdEncoding.DecodeString(config.SecretKey)
	if err != nil {
		return "", err
	}
	token, err = tokenGenerator.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return token, err
}

func ParseJWTToken(token string, config config.Config) (*JWTToken, error) {
	signingKey, err := base64.StdEncoding.DecodeString(config.SecretKey)
	if err != nil {
		return nil, err
	}
	parsedAccessToken, _ := jwt.ParseWithClaims(token, &JWTToken{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	return parsedAccessToken.Claims.(*JWTToken), nil
}
