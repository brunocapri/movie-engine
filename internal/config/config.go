package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	port      string
	db        string
	openAIKey string
	openAIUri string
}

func InitConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	cfg := Config{}

	envVars := map[string]*string{
		"PORT":       &cfg.port,
		"DB_URL":     &cfg.db,
		"OPENAI_URI": &cfg.openAIUri,
		"OPENAI_KEY": &cfg.openAIKey,
	}

	for env, ptr := range envVars {
		*ptr = os.Getenv(env)
		if *ptr == "" {
			log.Fatalf("%s environment variable is not set", env)
		}
	}
	return cfg
}

func (c Config) Port() string {
	return c.port
}

func (c Config) OpenAIKey() string {
	return c.openAIKey
}

func (c Config) DB() string {
	return c.db
}

func (c Config) OpenAIUri() string {
	return c.openAIUri
}
