package main

import (
	"fmt"
	"log"

	"github.com/Danielratmiroff/pretty-fzf/config"
	"github.com/Danielratmiroff/pretty-fzf/fzf"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "Pretty fzf",
	Short: "Pretty fzf wrapper",
	Long:  "Pretty fzf wrapper long",
	Run: func(cmd *cobra.Command, args []string) {
		if runFlag, _ := cmd.Flags().GetBool("run"); runFlag {
			fmt.Println("Running the prompt...")
		}

		params := fzf.NewDefaultConfig()

		err := fzf.RunCommand(params)

		if err != nil {
			log.Fatalf("Error running fzf: %v", err)
		}

	},
}

func init() {
	cobra.OnInitialize(initConfig)

	// Register flags
	rootCmd.Flags().BoolP("run", "r", false, "Execute the prompt")

	// Bind flags to viper
	viper.BindPFlag("run", rootCmd.Flags().Lookup("run"))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Var1: %s\n", config.Var1)
	fmt.Printf("Var2: %s\n", config.Var2)
	fmt.Printf("Var3: %s\n", config.Var3)
}
