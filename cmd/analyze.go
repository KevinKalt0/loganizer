package cmd

import (
	"fmt"
	"log"
	"loganizer/internal/analyzer"
	"loganizer/internal/config"
	"loganizer/internal/reporter"
	"sync"

	"github.com/spf13/cobra"
)

var configPath, outputPath string

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze logs based on provided configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		logConfigs, err := config.Load(configPath)
		if err != nil {
			return err
		}

		results := analyzer.AnalyzeLogs(logConfigs)

		if outputPath != "" {
			err := reporter.ExportResults(outputPath, results)
			if err != nil {
				return err
			}
		}

		analyzer.PrintSummary(results)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to config JSON file")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output path for report JSON")
	analyzeCmd.MarkFlagRequired("config")
}
