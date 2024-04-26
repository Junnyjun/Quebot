package domain

import "testing"

func TestToken(t *testing.T) {
	token := NewToken("1", "admin")

	if token.UserId() != "1" {
		t.Errorf("UserId does not match")
	}

	if token.Role() != "admin" {
		t.Errorf("Role does not match")
	}
}
