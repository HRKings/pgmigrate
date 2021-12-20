package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var (
	config        *Config
	executableDir string
)

type Config struct {
	ConnectionString string `yaml:"connection_string"`
	ExecutableDir    string
}

func GetConfig() *Config {
	if config == nil {
		executablePath, err := os.Executable()
		if err != nil {
			panic(err)
		}
		executableDir = filepath.Dir(executablePath)

		connectionStringEnvVariable := os.Getenv("PG_MIGRATE_CONNECTION_STRING")
		if connectionStringEnvVariable == "" {
			yamlFile, err := ioutil.ReadFile(filepath.Join(executableDir, "pgMigrate.yaml"))

			if err != nil {
				panic(err)
			}

			err = yaml.Unmarshal(yamlFile, &config)
			if err != nil {
				panic(err)
			}
		} else {
			config = &Config{
				ConnectionString: connectionStringEnvVariable,
			}
		}

		config.ExecutableDir = executableDir
	}

	return config
}
