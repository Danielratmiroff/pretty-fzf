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
	theme string
}

var params Params

var rootCmd = &cobra.Command{
	Use:   "Pretty fzf",
	Short: "Pretty fzf wrapper",
	Long:  "Pretty fzf wrapper long",
	Run: func(cmd *cobra.Command, args []string) {
		if err := saveParamsAsConfig(); err != nil {
			log.Fatalf("Error configuring with params: %v", err)
		}

		fzfConfig := fzf.NewDefaultConfig()

		err := fzf.RunCommand(fzfConfig)
		if err != nil {
			log.Fatalf("Error running fzf: %v", err)
		}
	},
}

func saveParamsAsConfig() error {
	if params.theme != "" {
		fmt.Printf("Set new theme: %s\n", params.theme)
		newTheme := themes.SelectTheme(params.theme)
		viper.Set("theme", newTheme)
	}

	return config.SaveConfig()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Register flags
	rootCmd.Flags().StringVarP(&params.theme, "theme", "t", "", "Set the theme")

	// Bind flags to viper
	viper.BindPFlags(rootCmd.Flags())
}

func initConfig() {
	err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
