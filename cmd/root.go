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
	Run: func(cmd *cobra.Command, args []string) {
		if version, _ := cmd.Flags().GetBool("version"); version {
			fmt.Println(version)
			return
		}

		if verboseMode, _ := cmd.Flags().GetBool("verbose"); verboseMode {
			config.SetVerboseMode(true)
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
	rootCmd.Flags().BoolP("verbose", "verbose", false, "Run in verbose mode")
}
