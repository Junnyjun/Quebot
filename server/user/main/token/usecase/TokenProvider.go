package usecase

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"myproject/main/global"
	"myproject/main/token/domain"
	"time"
)

type TokenProvider interface {
	GenerateToken(userId string, role string) string
	ValidateToken(token string) bool
}

type TokenProviderUsecase struct {
	config global.Config
}

func NewTokenProviderUsecase(config global.Config) *TokenProviderUsecase {
	return &TokenProviderUsecase{config: config}
}

func (tp *TokenProviderUsecase) GenerateToken(token domain.Token) string {
	claims := jwt.MapClaims{
		"user_id": token.UserId(),
		"role":    token.Role(),
		"exp":     time.Now().Add(time.Second * 24).UTC(),
	}

	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(tp.config.JwtKey()))
	if err != nil {
		log.Fatalln("Failed to generate token")
	}
	return signedToken
}

func (tp *TokenProviderUsecase) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(tp.config.JwtKey()), nil
	})
	if err != nil {
		log.Println("Token is not valid")
		return false
	}
	return true
}
