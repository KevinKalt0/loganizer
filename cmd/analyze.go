package cmd

import (
	"github.com/spf13/cobra"
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
	},
}
