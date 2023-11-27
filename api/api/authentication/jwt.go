package authentication

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	Email string
	jwt.StandardClaims
}

var JWTKey = []byte(os.Getenv("MY_VARIABLE"))

func GenerarToken(email string) (string, error) {
	envVarStr := os.Getenv("TOKEN_TIMEOUT")
	tokenTimeout, _ := strconv.Atoi(envVarStr)
	timeoutDuration := time.Duration(tokenTimeout) * time.Second

	claims := &JWTClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeoutDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKey)
}

func ParsearToken(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)

	if ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
