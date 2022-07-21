package auth

import "github.com/golang-jwt/jwt/v4"

type Service interface {
	GenerateToken(UserID int) (string, error)
}

type jwtService struct {
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
