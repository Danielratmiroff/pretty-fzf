// config/config.go
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Var1 string
	Var2 string
	Var3 string
}

func LoadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting home directory: %w", err)
	}
	configDir := filepath.Join(homeDir, ".pretty-fzf")

	// Ensure the config directory exists
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, fmt.Errorf("error creating config directory: %w", err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("PRETTYFZF")

	// Set default values
	viper.SetDefault("Var1", "default_value1")
	viper.SetDefault("Var2", "default_value2")
	viper.SetDefault("Var3", "default_value3")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; create it
			fmt.Println("Config file not found. Creating default config file.")
			if err := viper.SafeWriteConfig(); err != nil {
				return nil, fmt.Errorf("error creating default config file: %w", err)
			}
			fmt.Printf("Default config file created at: %s\n", viper.ConfigFileUsed())
		} else {
			// Config file was found but another error was produced
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into config struct: %w", err)
	}

	return &config, nil
}

// main.go remains largely the same
