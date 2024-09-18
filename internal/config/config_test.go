package config

import (
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func setup(t *testing.T, keyValueItems ...string) {
	fs := afero.NewMemMapFs()
	viper.SetFs(fs)

	f, err := fs.Create(".env")
	if assert.NoError(t, err) {
		defer f.Close()
		for _, keyValue := range keyValueItems {
			_, _ = f.Write([]byte(keyValue + "\n"))
		}
	}
}

func Test_Load_FromEnv(t *testing.T) {
	// Assemble
	setup(t)
	os.Setenv("PORT", "3000")
	os.Setenv("WEATHER_API_KEY", "key123")
	t.Cleanup(func() {
		os.Unsetenv("PORT")
		os.Unsetenv("WEATHER_API_KEY")
	})

	// Act
	cfg, err := Load()

	// Verify
	assert.NoError(t, err)
	if assert.NotNil(t, cfg) {
		assert.Equal(t, 3000, cfg.Port)
		assert.Equal(t, "key123", cfg.WeatherApiKey)
	}
}

func Test_Load_FromFile(t *testing.T) {
	// Assemble
	setup(t, "PORT=4000", "WEATHER_API_KEY=key456")

	// Act
	cfg, err := Load()

	// Verify
	assert.NoError(t, err)
	if assert.NotNil(t, cfg) {
		assert.Equal(t, 4000, cfg.Port)
		assert.Equal(t, "key456", cfg.WeatherApiKey)
	}
}

func Test_Load_EnvOverrideFile(t *testing.T) {
	// Assemble
	setup(t, "PORT=4000", "WEATHER_API_KEY=key456")
	os.Setenv("PORT", "3000")
	os.Setenv("WEATHER_API_KEY", "key123")
	t.Cleanup(func() {
		os.Unsetenv("PORT")
		os.Unsetenv("WEATHER_API_KEY")
	})

	// Act
	cfg, err := Load()

	// Verify
	assert.NoError(t, err)
	if assert.NotNil(t, cfg) {
		assert.Equal(t, 3000, cfg.Port)
		assert.Equal(t, "key123", cfg.WeatherApiKey)
	}
}
