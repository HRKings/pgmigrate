package utils

import (
	"os"
	"path/filepath"
)

var (
	config *Config
)

type Config struct {
	ConnectionString string
	ExecutableDir    string
}

func EnsureConfig() {
	if config == nil {
		executablePath, err := os.Executable()
		if err != nil {
			panic(err)
		}
		executableDir := filepath.Dir(executablePath)

		config = &Config{
			ExecutableDir: executableDir,
		}
	}
}

func GetConfig() *Config {
	EnsureConfig()

	return config
}

func SetConnectionString(connectionString string) {
	EnsureConfig()

	if connectionString == "" {
		panic("The connection string cannot be empty")
	}

	config.ConnectionString = connectionString
}
