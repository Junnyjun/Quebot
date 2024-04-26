package usecase

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"myproject/global"
	"myproject/token/domain"
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
		log.Fatal("Failed to generate token")
	}
	return signedToken
}

func (tp *TokenProviderUsecase) ValidateToken(token string) bool {
	// 토큰 검증 로직 구현
	return true // 임시 반환값
}
