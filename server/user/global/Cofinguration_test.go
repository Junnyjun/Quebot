package global

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	env := "test"

	config := GetInstance(env)
	assert.Equal(t, config.jwtKey, "secret", "jwtKey does not match")

	singleTone := GetInstance(env)
	assert.Equal(t, config, singleTone, "GetInstance should return the same instance")
}

func TestLoadConfigNoEnv(t *testing.T) {
	env := ""

	assert.Panic(t, func() {
		GetInstance(env)
	}, "ENV is not set")
}
