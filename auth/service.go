package auth

import (

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtservice struct {
}

var SECRET_KEY = []byte("BWASTARTUP_s3creT_k3Y")

func NewService() *jwtservice {
	return &jwtservice{}
}

func (s *jwtservice) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

