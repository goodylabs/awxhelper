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

var runbackupCmd = &cobra.Command{
	Use:   "runbackup",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		container := di.CreateContainer()
		err := container.Invoke(func(us *app.RunTemplateUseCase) error {
			return us.Execute("backup_")
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runbackupCmd)
}
