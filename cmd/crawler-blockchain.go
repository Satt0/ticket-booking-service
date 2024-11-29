/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	crawlerblockchain "http-server/internal/app/crawler-blockchain"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// crawlerBlockchainCmd represents the crawlerBlockchain command
var crawlerBlockchainCmd = &cobra.Command{
	Use:   "crawler",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fx.New(
			crawlerblockchain.CrawlerBlockchainModule,
			fx.Invoke(func(service *crawlerblockchain.CrawlerBlockchain) {
				service.Run()
			}),
		).Run()
	},
}

func init() {
	rootCmd.AddCommand(crawlerBlockchainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crawlerBlockchainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crawlerBlockchainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
