/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/garciawell/pos-go-stress-test/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pos-go-stress-test",
	Short: "Comando para teste de carga em serviços",
	Long:  `Comando para teste de carga em serviços. Utiliza a ferramenta vegeta para realizar os testes`,
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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("url", "u", "URL do serviço")
	rootCmd.Flags().Float64P("requests", "r", 1000, "Número total de requests")
	rootCmd.Flags().Float64P("concurrency", "c", 10, "Total de requests concorrentes")
	rootCmd.MarkFlagRequired("url")
}
