package service

import (
	"os"
	"time"

	"github.com/ShoreLab/shorelab-backend/lib/dto"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	ID string
	jwt.RegisteredClaims
}

func generateToken(userId string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		ID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	return tokenString, err
}

func ValidateToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("SECRET"), nil
	})

	return token, claims, err
}

func (s *Service) LoginService(authData *dto.AuthRequest) (*dto.AuthResponse, error) {
	u, err := s.repository.GetUser(authData.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(authData.Password))
	if err != nil {
		return nil, err
	}

	jwtToken, err := generateToken(u.ID)
	if err != nil {
		return nil, err
	}

	res := &dto.AuthResponse{
		Token: jwtToken,
		Name:  u.Name,
	}

	return res, nil
}
