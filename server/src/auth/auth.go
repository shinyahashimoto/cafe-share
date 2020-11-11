package auth

import (
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

// GetTokenHandler get token
func GetTokenHandler() string {

	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "54546557354"
	claims["name"] = "taro"
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 電子署名
	// tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	secret := "secretkey"
	tokenString, _ := token.SignedString([]byte(secret))

	// JWTを返却
	// w.Write([]byte(tokenString))
	return tokenString
}

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		// return []byte(os.Getenv("SIGNINGKEY")), nil
		secret := "secretkey"
		return []byte(secret), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
