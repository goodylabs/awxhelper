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

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure your connection to AWX (Ansible automation platform)",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		opts := &app.ConfigureOpts{
			URL:      url,
			Username: username,
			Password: password,
		}

		container := di.CreateContainer()
		err := container.Invoke(func(us *app.ConfigureUseCase) error {
			return us.Execute(opts)
		})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Successfuly configured!")
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.Flags().String("url", "", "AWX URL")
	configureCmd.Flags().String("username", "", "AWX username")
	configureCmd.Flags().String("password", "", "AWX password")
}
