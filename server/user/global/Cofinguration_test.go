package global

import (
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

func TestLoadConfigNoEnv(t *testing.T) {
	os.Clearenv()
	// ENV 환경 변수 삭제
	_ = os.Unsetenv("ENV")

	assert.Panic(t, func() {
		GetInstance()
	}, "ENV is not set")
}

func TestGetInstance(t *testing.T) {
	os.Clearenv()
	_ = os.Setenv("ENV", "test")

	config := GetInstance()
	assert.Equal(t, config.jwtKey, "secret", "jwtKey does not match")

	singleTone := GetInstance()
	assert.Equal(t, config, singleTone, "GetInstance should return the same instance")
}

func TestLoadConfigInvalidEnv(t *testing.T) {
	os.Clearenv()
	_ = os.Setenv("ENV", "invalid")

	assert.Panic(t, func() {
		GetInstance()
	}, "Error reading config file, Config File \"config-invalid\" Not Found in ")
}
