package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const DEFAULT_PORT = "8000"
const DEFAULT_BASE_URL = "http://localhost:3000"
const DEFAULT_DB_CONN = "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable"

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	emailApiKey := os.Getenv("EMAIL_SERVICE_API_KEY")
	emailBaseUrl := os.Getenv("EMAIL_SERVICE_BASE_URL")

	// Optional, defaults are constants in this file. DO NOT PUT ANY SECRETS HERE
	baseUrl := envGetOrDefault("BASE_URL", DEFAULT_BASE_URL)
	dbConnStr := envGetOrDefault("DB_CONN_STR", DEFAULT_DB_CONN)
	port := envGetOrDefault("PORT", DEFAULT_PORT)
	e = env{
		dbConnStr,
		emailApiKey,
		emailBaseUrl,
		baseUrl,
		port,
	}
	if err := validateEnv(); err != nil {
		panic(err)
	}
}

func envGetOrDefault(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	return val
}

func validateEnv() error {
	required := []string{
		"EMAIL_SERVICE_API_KEY",
		"EMAIL_SERVICE_BASE_URL",
	}
	issues := make([]string, 0)
	for _, key := range required {
		if os.Getenv(key) == "" {
			issues = append(issues, key)
		}
	}
	if len(issues) > 0 {
		fmt.Println(issues)
		return fmt.Errorf("missing required environment variables: %s", strings.Join(issues, ", "))
	}
	return nil
}

type Env interface {
	GetDbConnStr() string
	GetEmailApiKey() string
	GetEmailBaseUrl() string
	GetBaseUrl() string
	GetPort() string
}

type env struct {
	dbConnStr    string
	emailApiKey  string
	emailBaseUrl string
	baseUrl      string
	port         string
}

func (e env) GetPort() string {
	return e.port
}

func (e env) GetBaseUrl() string {
	return e.baseUrl
}

var e env

func (e env) GetDbConnStr() string {
	return e.dbConnStr
}

func (e env) GetEmailApiKey() string {
	return e.emailApiKey
}

func (e env) GetEmailBaseUrl() string {
	return e.emailBaseUrl
}

func GetEnv() Env {
	return e
}
