package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
)

// TODO 生成secret
// var hmacSampleSecret = []byte("I'm a secret")

func GenerateJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": 123,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	var hmacSampleSecret = []byte("I'm a secret")
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		panic(err)
	}
	return tokenString, nil
}

func JWTAuth(secret string) middleware.Middleware {
	hmacSampleSecret := []byte(secret)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {

				authString := tr.RequestHeader().Get("Authorization")

				spew.Dump(authString)

				auths := strings.SplitN(authString, " ", 2)
				spew.Dump(auths)
				if len(auths) != 2 || auths[0] != "Token" {
					return nil, errors.New("jwt token missing !")
				}
				tokenString := auths[1]
				// spew.Dump(tokenString)

				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
					return hmacSampleSecret, nil
				})

				if err != nil {
					return nil, err
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					spew.Dump(claims["userId"])
				} else {
					fmt.Println(err)
				}

			}
			return handler(ctx, req)
		}
	}
}
