package cfg

import (
	"log"
	"os"
	"sync"

	env "github.com/joho/godotenv"
)

type Config struct {
	HttpPort string
}

var (
	once     sync.Once
	instance *Config
)

func Inst() *Config {
	once.Do(func() {
		err := env.Load()
		if err != nil {
			// Warn if the .env file could not be loaded, falling back to system environment variables
			log.Printf("Warning: Could not load .env file, falling back to system environment variables")
		}

		instance = &Config{
			HttpPort: getEnv("HTTP_PORT", "8080"),
		}
	})
	return instance
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
