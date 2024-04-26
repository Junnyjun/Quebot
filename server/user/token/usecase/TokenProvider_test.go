package usecase

import (
	"myproject/token/domain"
	"testing"
)

func TestTokenProviderUsecase_GenerateToken(t *testing.T) {
	tokenProvider := NewTokenProviderUsecase()

	token := tokenProvider.GenerateToken(domain.NewToken("1", "admin"))

	if token == "" {
		t.Errorf("Token is empty")
	}

}
