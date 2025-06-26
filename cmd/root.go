package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "Un outil CLI pour analyser des fichiers de logs",
	Long:  "loganalyzer permet d'analyser plusieurs logs de manière concurrente et d'en extraire un rapport.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Erreur : ", err)
		os.Exit(1)
	}
}
