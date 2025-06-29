package cmd

import (
	"fmt"
	"github.com/KevinKalt0/loganizer/internal/config"
	"github.com/spf13/cobra"
	"os"
	"github.com/KevinKalt0/loganizer/internal/analyzer"
	"github.com/KevinKalt0/loganizer/internal/reporter"
)

var (
	configPath string
	outputPath string
)

func init() {
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Chemin vers le fichier de configuration JSON")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Chemin de sortie du fichier rapport JSON")
	analyzeCmd.MarkFlagRequired("config")

	rootCmd.AddCommand(analyzeCmd)
}

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse les fichiers de logs définis dans un fichier JSON",
	Run: func(cmd *cobra.Command, args []string) {
		logs, err := config.LoadConfig(configPath)
		if err != nil {
			fmt.Printf("Erreur de chargement du fichier de conf : %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Configuration chargée (%d logs) :\n", len(logs))
		for _, log := range logs {
			fmt.Printf(" - ID: %s | Path: %s | Type: %s\n", log.ID, log.Path, log.Type)
		}

		fmt.Println("\nAnalyse en cours...")
		results := analyzer.AnalyzeLogs(logs)

		fmt.Println("\nRésultats :")
		for _, res := range results {
			fmt.Printf("ID: %s | Statut: %s | Message: %s\n", res.LogID, res.Status, res.Message)
			if res.ErrorDetails != "" {
				fmt.Printf("  -> Détail : %s\n", res.ErrorDetails)
			}
		}

		if outputPath != "" {
			if err := reporter.Export(results, outputPath); err != nil {
				fmt.Printf("Erreur export JSON : %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Rapport exporté vers : %s\n", outputPath)
		}
	},
}

