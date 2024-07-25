package config

import (
	"os"
	"strconv"
)

type Config struct {
	Address string
	DB      DBConfig
}

type DBConfig struct {
	Address      string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

func LoadConfig() Config {
	return Config{
		Address: GetString("ADDRESS", ":8080"),
		DB: DBConfig{
			Address:      GetString("DB_ADDRESS", "postgres://user:adminpassword@localhost/social?sslmode=disable"),
			MaxOpenConns: GetInt("DB_MAX_DB_CONNS", 30),
			MaxIdleConns: GetInt("DB_MAX_DB_CONNS", 30),
			MaxIdleTime:  GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}
}

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return valAsInt
}
