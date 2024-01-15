package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvAsStringFn(t *testing.T) {
	// Test when environment variable exists
	os.Setenv("TEST_ENV", "test_value")
	value, err := getEnvAsStringFunc("TEST_ENV")
	assert.NoError(t, err)
	assert.Equal(t, "test_value", value)
	os.Unsetenv("TEST_ENV")

	// Test when environment variable does not exist and default value is provided
	value, err = getEnvAsStringFunc("NON_EXISTENT_ENV", "default_value")
	assert.NoError(t, err)
	assert.Equal(t, "default_value", value)

	// Test when environment variable does not exist and no default value is provided
	// This should return an error
	value, err = getEnvAsStringFunc("NON_EXISTENT_ENV")
	assert.Error(t, err)
	assert.Equal(t, "", value)

	// Test when environment variable exists and default value is provided
	os.Setenv("TEST_ENV", "test_value")
	value, err = getEnvAsStringFunc("TEST_ENV", "default_value")
	assert.NoError(t, err)
	assert.Equal(t, "test_value", value)
	os.Unsetenv("TEST_ENV")

	// Test when environment variable is empty and default value is provided
	value, err = getEnvAsStringFunc("TEST_ENV", "default_value")
	assert.NoError(t, err)
	assert.Equal(t, "default_value", value)
}

func TestGetEnvAsIntFunc(t *testing.T) {
	// Test when environment variable exists and is an integer
	os.Setenv("TEST_ENV", "123")
	value, err := getEnvAsIntFunc("TEST_ENV")
	assert.NoError(t, err)
	assert.Equal(t, 123, value)
	os.Unsetenv("TEST_ENV")

	// Test when environment variable does not exist and default value is provided
	value, err = getEnvAsIntFunc("NON_EXISTENT_ENV", 456)
	assert.NoError(t, err)
	assert.Equal(t, 456, value)

	// Test when environment variable does not exist and no default value is provided
	// This should return an error
	value, err = getEnvAsIntFunc("NON_EXISTENT_ENV")
	assert.Error(t, err)
	assert.Equal(t, 0, value)

	// Test when environment variable exists, is not an integer and default value is provided
	os.Setenv("TEST_ENV", "not_an_integer")
	value, err = getEnvAsIntFunc("TEST_ENV", 789)
	assert.NoError(t, err)
	assert.Equal(t, 789, value)
	os.Unsetenv("TEST_ENV")

	// Test when environment variable exists, is not an integer and no default value is provided
	// This should return an error
	os.Setenv("TEST_ENV", "not_an_integer")
	value, err = getEnvAsIntFunc("TEST_ENV")
	assert.Error(t, err)
	assert.Equal(t, 0, value)
	os.Unsetenv("TEST_ENV")
}
