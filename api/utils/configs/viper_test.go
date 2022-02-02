package configs

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var sampleConfig = []byte(`
DATABASE_NAME: "cms"
`)

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
