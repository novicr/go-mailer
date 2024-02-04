package config

import (
	"encoding/json"
	"os"
)

// Config represents the application's configurations
type Config struct {
	SMTPHost     string `json:"smtp_host"`
	SMTPPort     string `json:"smtp_port"`
	SMTPUsername string `json:"smtp_username"`
	SMTPPassword string `json:"smtp_password_env_var"` // Name of the environment variable
}

// LoadConfig loads configuration from a JSON file
func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	// Read the SMTP password from the environment variable
	config.SMTPPassword = os.Getenv(config.SMTPPassword)

	return config, nil
}
