package jwthelper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const JWT_SECRET_KEY = "secret"
const JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT = 1440

// GenerateNewTokens func for generate a new Access & Refresh tokens.
func GenerateNewToken(data Login, credentials []string, additionalParam map[string]interface{}) (*Token, error) {
	// Generate JWT Access token.
	accessToken, err := generateNewAccessToken(data, credentials, additionalParam)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &Token{
		Access: accessToken,
	}, nil
}

func generateNewAccessToken(user Login, credentials []string, additionalParam map[string]interface{}) (string, error) {

	// Set secret key from .env file.
	secret := JWT_SECRET_KEY
	// Set expires minutes count for secret key from .env file.
	minutesCount := JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT

	// additionalParamPrint, _ := json.MarshalIndent(additionalParam, "", "\t")
	// log.Println("### additionalParamPrint ###")
	// log.Println(string(additionalParamPrint))
	// log.Println("### End Of additionalParamPrint ###")

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["id"] = user.ID
	claims["email"] = user.Email

	claims["expires"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	// claims["book:create"] = false
	// claims["book:update"] = false
	// claims["book:delete"] = false

	// Set private token credentials:
	for _, credential := range credentials {
		claims[credential] = true
	}

	for key, value := range additionalParam {
		claims[key] = value
	}

	// claimsPrint, _ := json.MarshalIndent(claims, "", "\t")
	// log.Println("### claimsPrint ###")
	// log.Println(string(claimsPrint))
	// log.Println("### End Of claimsPrint ###")

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// tokenPrint, _ := json.MarshalIndent(token, "", "\t")
	// log.Println("### tokenPrint ###")
	// log.Println(string(tokenPrint))
	// log.Println("### End Of tokenPrint ###")

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
