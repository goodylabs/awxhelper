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

var customtemplateCmd = &cobra.Command{
	Use:   "customtemplate",
	Short: "Run custom template command",
	Run: func(cmd *cobra.Command, args []string) {
		err := di.CreateContainer().Invoke(func(us *app.CustomTemplateUseCase) error {
			return us.Execute()
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	customtemplateCmd.Hidden = true
	rootCmd.AddCommand(customtemplateCmd)
}
