/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/moha-abdi/telxchange/internal/merchant"
	"github.com/spf13/cobra"
)

var (
	maxMerchants      int
	maxGoroutines     int
	startWithMerchant int
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate and check merchant accounts",
	Long: `This command generates and checks with the API endpoint to find
a valid merchant IDs starting from 300000 until provided limit + 300000.`,
	Run: func(cmd *cobra.Command, args []string) {
		merchant.GenerateMerchants(maxMerchants, maxGoroutines, startWithMerchant)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().IntVarP(&maxMerchants, "max-merchants", "m", 10, "Maximum number of merchants to generate")
	generateCmd.Flags().IntVarP(&maxGoroutines, "max-goroutines", "g", 1, "Maximum number of concurrent goroutine channels")
	generateCmd.Flags().IntVarP(&startWithMerchant, "start", "s", 3000000, "The merchant number to start generation from")
}
