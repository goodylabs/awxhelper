/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/goodylabs/awxhelper/pkg/config"
	"github.com/spf13/cobra"
)

var forceupdateCmd = &cobra.Command{
	Use:   "forceupdate",
	Short: "Force check for new updates and install if available",
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.GetReleaser().ForceUpdate(); err != nil {
			fmt.Println("Error checking for updates:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(forceupdateCmd)
}
