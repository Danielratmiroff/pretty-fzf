package main

import (
	"fmt"
	"log"

	"github.com/Danielratmiroff/pretty-fzf/config"
	"github.com/Danielratmiroff/pretty-fzf/fzf"
	"github.com/Danielratmiroff/pretty-fzf/themes"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Params struct {
	Theme string
}

var params Params

var rootCmd = &cobra.Command{
	Use:   "Pretty fzf",
	Short: "Pretty fzf wrapper",
	Long:  "Pretty fzf wrapper long",
	Run: func(cmd *cobra.Command, args []string) {
		if params.Theme != "" {
			fmt.Printf("Set new theme: %s\n", params.Theme)
			newTheme := themes.SelectTheme(params.Theme)
			viper.Set("theme", newTheme)
		}

		fzfConfig := fzf.NewDefaultConfig()
		// Apply params to fzfConfig here

		err := fzf.RunCommand(fzfConfig)
		if err != nil {
			log.Fatalf("Error running fzf: %v", err)
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	// Register flags
	rootCmd.Flags().StringVarP(&params.Theme, "theme", "t", "", "Set the theme")
	// Add more flags here

	// Bind flags to viper
	viper.BindPFlags(rootCmd.Flags())
}

func initConfig() {
	var err error
	// Load the configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Defaults
	theme := viper.GetString("theme")
	if theme == "" {
		viper.Set("theme", themes.Catpuccino())
	}

	viper.Set("Var1", config.Var1)
	viper.Set("Var2", config.Var2)
	viper.Set("Var3", config.Var3)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
