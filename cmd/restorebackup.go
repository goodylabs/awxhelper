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

var restorebackupCmd = &cobra.Command{
	Use:   "restorebackup",
	Short: "Restore db from backup with specified job",
	Run: func(cmd *cobra.Command, args []string) {
		container := di.CreateContainer()
		err := container.Invoke(func(us *app.RunTemplateUseCase) error {
			return us.Execute("restore_")
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	restorebackupCmd.Hidden = true
	rootCmd.AddCommand(restorebackupCmd)
}
