package gocrud

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// This is for generating tokens
// not much to explain
func GenerateToken(userID int32) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 168).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(GoCRUDConfig.authSecret))
	return t, err
}

// This is for getting the jwt from header
// not much to explain
func GetJWTFromHeader(c *fiber.Ctx, authScheme string) (string, error) {

	auth := c.Get("Authorization")
	l := len(authScheme)
	if len(auth) > l+1 && strings.EqualFold(auth[:l], authScheme) {
		return strings.TrimSpace(auth[l:]), nil
	}
	return "", fmt.Errorf("missing or malformed JWT")
}

// This is for decoding the jwt
// not much to explain
func DecodeJWT(tokenString string) (jwt.MapClaims, error) {

	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Ensure the signing method is HMAC (HS256) for security
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GoCRUDConfig.authSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract and return claims if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
