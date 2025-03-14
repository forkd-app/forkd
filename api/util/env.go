package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type environment = string

const DEV_ENV = environment("DEV")
const TEST_ENV = environment("TEST")
const PROD_ENV = environment("PROD")

const DEFAULT_PORT = "8000"
const DEFAULT_BASE_URL = "http://localhost:3000"
const DEFAULT_DB_CONN = "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable"
const DEFAULT_OBJECT_STORAGE_HOST = "minio:9000"
const DEFAULT_OBJECT_STORAGE_ACCESS_KEY = "forkd"
const DEFAULT_OBJECT_STORAGE_SECRET_KEY = "forkd-secret-key"

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	emailApiKey := os.Getenv("EMAIL_SERVICE_API_KEY")
	emailBaseUrl := os.Getenv("EMAIL_SERVICE_BASE_URL")

	// Optional, defaults are constants in this file. DO NOT PUT ANY SECRETS HERE
	baseUrl := envGetOrDefault("BASE_URL", func(s string) string {
		if s == "" {
			return DEFAULT_BASE_URL
		}
		return s
	})
	dbConnStr := envGetOrDefault("DB_CONN_STR", func(s string) string {
		if s == "" {
			return DEFAULT_DB_CONN
		}
		return s
	})
	port := envGetOrDefault("PORT", func(s string) string {
		if s == "" {
			return DEFAULT_PORT
		}
		return s
	})
	sendMagicEmail := envGetOrDefault("SEND_MAGIC_LINK_EMAIL", func(s string) bool {
		return strings.ToLower(s) != "false"
	})

	objectStorageHost := envGetOrDefault("OBJECT_STORAGE_HOST", func(s string) string {
		if s == "" {
			return DEFAULT_OBJECT_STORAGE_HOST
		}
		return s
	})

	objectStorageAccessKey := envGetOrDefault("OBJECT_STORAGE_ACCESS_KEY", func(s string) string {
		if s == "" {
			return DEFAULT_OBJECT_STORAGE_ACCESS_KEY
		}
		return s
	})

	objectStorageSecretKey := envGetOrDefault("OBJECT_STORAGE_SECRET_KEY", func(s string) string {
		if s == "" {
			return DEFAULT_OBJECT_STORAGE_SECRET_KEY
		}
		return s
	})

	environ := envGetOrDefault("ENVIRONMENT", func(s string) environment {
		switch strings.ToLower(s) {
		case "prod":
			return PROD_ENV
		case "test":
			return TEST_ENV
		case "dev":
			fallthrough
		case "":
			return DEV_ENV
		}

		panic(fmt.Sprintf(
			"invalid environment. Got %s, expected one of: %s",
			s,
			strings.Join(
				[]environment{
					DEV_ENV,
					TEST_ENV,
					PROD_ENV,
				},
				", ",
			),
		))
	})

	e = env{
		dbConnStr:              dbConnStr,
		emailApiKey:            emailApiKey,
		emailBaseUrl:           emailBaseUrl,
		baseUrl:                baseUrl,
		port:                   port,
		sendMagicLinkEmail:     sendMagicEmail,
		objectStorageHost:      objectStorageHost,
		objectStorageAccessKey: objectStorageAccessKey,
		objectStorageSecretKey: objectStorageSecretKey,
		environment:            environ,
	}
	if err := validateEnv(); err != nil {
		panic(err)
	}
}

func envGetOrDefault[T any](key string, coerceVal func(string) T) T {
	return coerceVal(os.Getenv(key))
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
	GetSendMagicLinkEmail() bool
	GetObjectStorageHost() string
	GetObjectStorageAccessKey() string
	GetObjectStorageSecretKey() string
	GetEnvironment() environment
}

type env struct {
	dbConnStr              string
	emailApiKey            string
	emailBaseUrl           string
	baseUrl                string
	port                   string
	sendMagicLinkEmail     bool
	objectStorageHost      string
	objectStorageAccessKey string
	objectStorageSecretKey string
	environment            string
}

// GetEnvironment implements Env.
func (e env) GetEnvironment() string {
	return e.environment
}

// GetObjectStorageAccessKey implements Env.
func (e env) GetObjectStorageAccessKey() string {
	return e.objectStorageAccessKey
}

// GetObjectStorageHost implements Env.
func (e env) GetObjectStorageHost() string {
	return e.objectStorageHost
}

// GetObjectStorageSecretKey implements Env.
func (e env) GetObjectStorageSecretKey() string {
	return e.objectStorageSecretKey
}

func (e env) GetSendMagicLinkEmail() bool {
	return e.sendMagicLinkEmail
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
