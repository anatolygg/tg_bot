package config

import (
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	LogPath   string `yaml:"log_path"`
	MLService `yaml:"ml_service"`
}

type MLService struct {
	URL string `yaml:"url"`
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func LoadEnv(envName, defaultName string) string {
	if err := godotenv.Load(".env"); err != nil {
		return defaultName
	}

	if env, exists := os.LookupEnv(envName); exists {
		return env
	}
	return defaultName
}
