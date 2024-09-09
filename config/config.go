package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}
	configDir := filepath.Join(homeDir, ".pretty-fzf")

	// Ensure the config directory exists
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("PRETTYFZF")

	// Set default values
	viper.SetDefault("theme", "catpuccino")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; create it
			fmt.Println("Config file not found. Creating default config file.")
			if err := viper.SafeWriteConfig(); err != nil {
				return fmt.Errorf("error creating default config file: %w", err)
			}
			fmt.Printf("Default config file created at: %s\n", viper.ConfigFileUsed())
		} else {
			// Config file was found but another error was produced
			return fmt.Errorf("error reading config file: %w", err)
		}
	}
	return nil
}

func SaveConfig() error {
	return viper.WriteConfig()
}
