package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCEP_GetDigits(t *testing.T) {
	c := CEP("89120-000")
	assert.Equal(t, "89120000", c.GetDigits())
}

func TestCEP_IsValid(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		assert.True(t, CEP("89120-000").IsValid())
		assert.True(t, CEP("89120000").IsValid())
	})
	t.Run("Invalid", func(t *testing.T) {
		assert.False(t, CEP("89120.000").IsValid())
		assert.False(t, CEP("89120-0000").IsValid())
		assert.False(t, CEP("891200000").IsValid())
	})
}
