/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	app "http-server/internal/app/api"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// apiLaunchCmd represents the apiLaunch command
var apiLaunchCmd = &cobra.Command{
	Use:   "api:launch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fx.New(
			app.ApiServerModule,
			fx.Invoke(func(api *app.ApiServer) {
				api.StartServer()
			}),
		).Run()
	},
}

func init() {
	rootCmd.AddCommand(apiLaunchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiLaunchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiLaunchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
