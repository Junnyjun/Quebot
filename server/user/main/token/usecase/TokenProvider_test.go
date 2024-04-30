package usecase

import (
	"myproject/main/global"
	"myproject/main/token/domain"
	"testing"
)

func TestTokenProviderUsecase_GenerateToken(t *testing.T) {
	tokenProvider := NewTokenProviderUsecase(*global.GetInstance("test"))
	token := tokenProvider.GenerateToken(*domain.NewToken("1", "admin"))

	if token == "" {
		t.Errorf("Token is empty")
	}

}

func TestTokenProviderUsecase_ValidateToken(t *testing.T) {
	tokenProvider := NewTokenProviderUsecase(*global.GetInstance("test"))
	token := tokenProvider.GenerateToken(*domain.NewToken("1", "admin"))

	if !tokenProvider.ValidateToken(token) {
		t.Errorf("Token is not valid")
	}
