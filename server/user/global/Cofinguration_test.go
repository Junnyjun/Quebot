package global

import (
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestGetInstance(t *testing.T) {
	os.Setenv("ENV", "test")
	viper.AddConfigPath("path_to_test_config")

	config := GetInstance("test")
	assert.Equal(t, config.jwtKey, "secret", "jwtKey does not match")

	singleTone := GetInstance("test")
	assert.Equal(t, config, singleTone, "GetInstance should return the same instance")
}
