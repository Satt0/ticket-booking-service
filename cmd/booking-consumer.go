/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	bookingjob "http-server/internal/app/booking-consumer"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// checkOutdatedPaymentCmd represents the checkOutdatedPayment command
var bookingConsumerCmd = &cobra.Command{
	Use:   "consume-booking-job",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fx.New(
			bookingjob.BookingJobModule,
			fx.Invoke(func(service *bookingjob.BookingJob) {
				service.ConsumeBookingJob()
			}),
		).Run()
	},
}

func init() {
	rootCmd.AddCommand(bookingConsumerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkOutdatedPaymentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkOutdatedPaymentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
