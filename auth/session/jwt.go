package session

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignSessionFunc = func(userId string) (*string, error)

func CreateWithSecret(secret string) SignSessionFunc {
	return func(userId string) (*string, error) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss":    "bbnb-booking-api",                                    // issuer
			"sub":    "auth",                                                // subject
			"aud":    "webstart students",                                   // audience
			"userId": userId,                                                // extra info
			"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(), // not before
		})
		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			return nil, err
		}
		return &tokenString, nil
	}
}
