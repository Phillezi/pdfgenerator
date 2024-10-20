package cmd

import (
	"os"

	"github.com/Phillezi/pdfgenerator/internal/generator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "pdfgenerator [patterns]",
	Short: "Collect and append file contents based on regex patterns into a PDF",
	Args:  cobra.MinimumNArgs(0),
	Run:   generator.Run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var outputFile string

	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "output.pdf", "The PDF file where contents will be appended")
	viper.BindPFlag("outputFile", rootCmd.Flags().Lookup("output"))
	viper.SetDefault("outputFile", "output.pdf")
}
