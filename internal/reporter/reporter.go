package reporter

import (
	"encoding/json"
	"fmt"
	"github.com/KevinKalt0/loganizer/internal/analyzer"
	"os"
)

func Export(results []analyzer.Result, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("impossible de créer le fichier : %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "erreur lors de la fermeture du fichier : %v\n", cerr)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(results); err != nil {
		return fmt.Errorf("erreur d'encodage JSON : %w", err)
	}

	return nil
}
