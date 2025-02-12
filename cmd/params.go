/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/garciawell/pos-go-stress-test/internal"
	"github.com/spf13/cobra"
)

// paramsCmd represents the params command
var paramsCmd = &cobra.Command{
	Use:   "params",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetFloat64("requests")
		concurrency, _ := cmd.Flags().GetFloat64("concurrency")

		internal.VegetaRun(internal.VegetaType{
			Url:         url,
			Requests:    requests,
			Concurrency: concurrency,
		})
	},
}

func init() {
	rootCmd.AddCommand(paramsCmd)
	paramsCmd.Flags().String("url", "u", "URL do serviço")
	paramsCmd.Flags().Float64P("requests", "r", 1000, "Número total de requests")
	paramsCmd.Flags().Float64P("concurrency", "c", 10, "Total de requests concorrentes")
	paramsCmd.MarkFlagRequired("url")
}
