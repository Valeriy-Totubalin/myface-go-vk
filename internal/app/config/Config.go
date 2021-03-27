package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Valeriy-Totubalin/myface-go-vk/internal/app/interfaces/config_interfaces"
	"github.com/joho/godotenv"
)

const NoEnvFile = "No .env file found"

type Config struct {
	db     config_interfaces.DBConfig
	server config_interfaces.ServerConfig
}

func NewConfig(dbConfig config_interfaces.DBConfig, serverConfig config_interfaces.ServerConfig) *Config {
	return &Config{
		db:     dbConfig,
		server: serverConfig,
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsSeconds(name string, defaultVal time.Duration) time.Duration {
	valStr := getEnv(name, "")
	if value, err := strconv.Atoi(valStr); err == nil {
		return time.Duration(value) * time.Second
	}

	return defaultVal
}

func (c *Config) DB() config_interfaces.DBConfig {
	return c.db
}

func (c *Config) Srv() config_interfaces.ServerConfig {
	return c.server
}

func Init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print(NoEnvFile)
	}
}
