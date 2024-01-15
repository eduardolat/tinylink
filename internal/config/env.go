package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/eduardolat/tinylink/internal/logger"
	"github.com/joho/godotenv"
)

type DBType string

const (
	InMemoryDBType DBType = "in-memory"
)

const (
	DefaultDBType DBType = InMemoryDBType
	DefaultPort   int    = 3000
)

type Env struct {
	DBType DBType
	PORT   int
}

func GetEnv() *Env {
	err := godotenv.Load()
	if err == nil {
		logger.Info("👉 Using .env file")
	}

	env := &Env{
		DBType: DBType(getEnvAsString("DB_TYPE", string(DefaultDBType))),
		PORT:   getEnvAsInt("PORT", DefaultPort),
	}

	return env
}

func getEnvAsString(name string, defaultValue ...string) string {
	value, err := getEnvAsStringFunc(name, defaultValue...)

	if err != nil {
		logger.FatalError(
			"error getting env variable",
			"name", name,
			"error", err,
		)
	}

	return value
}

func getEnvAsStringFunc(name string, defaultValue ...string) (string, error) {
	defaultValueItem := ""
	if len(defaultValue) > 0 {
		defaultValueItem = defaultValue[0]
	}

	value, exists := os.LookupEnv(name)

	if !exists && defaultValueItem == "" {
		return "", errors.New("env variable does not exist")
	}

	if !exists && defaultValueItem != "" {
		return defaultValueItem, nil
	}

	return value, nil
}

func getEnvAsInt(name string, defaultValue ...int) int {
	value, err := getEnvAsIntFunc(name, defaultValue...)

	if err != nil {
		logger.FatalError(
			"error getting env variable",
			"name", name,
			"error", err,
		)
	}

	return value
}

func getEnvAsIntFunc(name string, defaultValue ...int) (int, error) {
	var defaultValueItem *int
	if len(defaultValue) > 0 {
		defaultValueItem = &defaultValue[0]
	}

	valueStr, err := getEnvAsStringFunc(name)
	if err != nil && defaultValueItem == nil {
		return 0, errors.New("env variable does not exist")
	}

	if err != nil && defaultValueItem != nil {
		return *defaultValueItem, nil
	}

	value, err := strconv.Atoi(valueStr)

	if err != nil && defaultValueItem == nil {
		return 0, errors.New("env variable is not an integer")
	}

	if err != nil && defaultValueItem != nil {
		return *defaultValueItem, nil
	}

	return value, nil
}
