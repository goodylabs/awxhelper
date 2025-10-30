/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/goodylabs/awxhelper/internal/app"
	"github.com/goodylabs/awxhelper/internal/domain/entities"
	"github.com/goodylabs/awxhelper/pkg/di"
	"github.com/spf13/cobra"
)

var downloaddbCmd = &cobra.Command{
	Use:   "downloaddb",
	Short: "Download database dump",
	Run: func(cmd *cobra.Command, args []string) {

		var extraVars = make(entities.ExtraVars)

		// date, _ := cmd.Flags().GetString("date")
		// if date != "" {
		// 	extraVars["date"] = date
		// }

		err := di.CreateContainer().Invoke(func(us *app.DownloadDB) error {
			return us.Execute("download_db__", &extraVars)
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downloaddbCmd)
	// downloaddbCmd.Flags().StringP("date", "d", "", "Specify date in YYYY-MM-DD format to download db from that date")
}
