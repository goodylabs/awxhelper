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

var runJobCmd = &cobra.Command{
	Use:   "runjob",
	Short: "Run any job you want!",
	Run: func(cmd *cobra.Command, args []string) {
		container := di.CreateContainer()
		err := container.Invoke(func(us *app.RunTemplateUseCase) error {
			return us.Execute("")
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runJobCmd)
}
