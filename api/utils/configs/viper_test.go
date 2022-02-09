package configs

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	viper.Set("DATABASE_NAME", "cms")
	value := GetConfig("DATABASE_NAME")
	assert.Equal(t, "cms", value)
}

func TestSetConfig(t *testing.T) {
	SetConfig("DATABASE_NAME_2", "cms")
	value := GetConfig("DATABASE_NAME_2")
	assert.Equal(t, "cms", value)
}
