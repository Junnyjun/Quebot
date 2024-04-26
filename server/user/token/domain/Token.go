package domain

import "github.com/pelletier/go-toml/v2"

type Token struct {
	userId    string
	role      string
	expiresAt toml.LocalDateTime
}

func NewToken(userId string, role string) *Token {
	return &Token{
		userId: userId,
		role:   role,
	}
}

func (t *Token) UserId() string {
	return t.userId
}

func (t *Token) Role() string {
	return t.role
}
