package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToken(UserID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("123456")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signToken, err
	}

	return signToken, nil

}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil

}
