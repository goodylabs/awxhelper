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
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		container := di.CreateContainer()
		err := container.Invoke(func(us *app.RestoreBackupUseCase) error {
			return us.Execute()
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(restorebackupCmd)
}
