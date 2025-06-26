package reporter

import (
	"encoding/json"
	"os"
	"loganalyzer/internal/analyzer"
)

func Export(results []analyzer.Result, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(results)
}
