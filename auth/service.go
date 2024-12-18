package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	validateToken(token string) (*jwt.Token, error)
}

type jwtservice struct {
}

// validateToken implements Service.
func (s *jwtservice) validateToken(token string) (*jwt.Token, error) {
	panic("unimplemented")
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

func (s *jwtservice) ValidateToken(encodedtoken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedtoken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
