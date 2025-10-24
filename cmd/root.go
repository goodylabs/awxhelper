package cmd

import (
	"fmt"
	"os"

	"github.com/goodylabs/awxhelper/pkg/config"
	"github.com/spf13/cobra"
)

var version = "dev"

var rootCmd = &cobra.Command{
	Use:   "awxhelper",
	Short: "A brief description of your application",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if debugMode, _ := cmd.Flags().GetBool("debug"); debugMode {
			config.SetDebugMode(true)
			fmt.Println("--> Debug mode enabled.")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if v, _ := cmd.Flags().GetBool("version"); v {
			fmt.Println(version)
			return
		}

		if debugMode, _ := cmd.Flags().GetBool("debug"); debugMode {
			config.SetDebugMode(true)
		}

		cmd.Help()
	},
}

func Execute() {
	updated, err := config.GetReleaser().Run()
	if err != nil {
		fmt.Println("Error checking for updates:", err)
	} else if updated {
		fmt.Println("Application has been updated.")
		os.Exit(0)
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("version", "v", false, "Print version and exit")
	rootCmd.PersistentFlags().BoolP("debug", "x", false, "Run in debug mode")
}
