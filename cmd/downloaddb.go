/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/goodylabs/awxhelper/internal/app"
	"github.com/goodylabs/awxhelper/pkg/di"
	"github.com/spf13/cobra"
)

var downloaddbCmd = &cobra.Command{
	Use:   "downloaddb",
	Short: "Download database dump",
	Run: func(cmd *cobra.Command, args []string) {
		err := di.CreateContainer().Invoke(func(us *app.RunDownloadDB) error {
			return us.Execute("download_db__")
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downloaddbCmd)
}
