package analyzer

import (
	"loganalyzer/internal/config"
	"github.com/KevinKalt0/loganalyzer/internal/analyzer"
	"github.com/KevinKalt0/loganalyzer/internal/reporter"
	)

type Result struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"` 
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details,omitempty"`
}

package analyzer

import (
"errors"
"fmt"
"math/rand"
"os"
"sync"
"time"
"loganalyzer/internal/config"
)
func init() {
	rand.Seed(time.Now().UnixNano())
	}
	

func AnalyzeLogs(logs []config.LogConfig) []Result {
var wg sync.WaitGroup
resultsChan := make(chan Result, len(logs))
for _, log := range logs {
    wg.Add(1)
    go func(log config.LogConfig) {
        defer wg.Done()
        res := analyzeSingleLog(log)
        resultsChan <- res
    }(log)
}

wg.Wait()
close(resultsChan)

var results []Result
for r := range resultsChan {
    results = append(results, r)
}
return results
}

func analyzeSingleLog(log config.LogConfig) Result {
	if _, err := os.Stat(log.Path); err != nil {
	return Result{
	LogID: log.ID,
	FilePath: log.Path,
	Status: "FAILED",
	Message: "Fichier introuvable.",
	ErrorDetails: err.Error(),
	}
	}


time.Sleep(time.Duration(rand.Intn(150)+50) * time.Millisecond)


if rand.Float64() < 0.10 {
    parseErr := &ParsingError{LogID: log.ID}
    return Result{
        LogID:        log.ID,
        FilePath:     log.Path,
        Status:       "FAILED",
        Message:      "Erreur de parsing.",
        ErrorDetails: parseErr.Error(),
    }
}

return Result{
    LogID:        log.ID,
    FilePath:     log.Path,
    Status:       "OK",
    Message:      "Analyse terminée avec succès.",
    ErrorDetails: "",
}
}

Run: func(cmd *cobra.Command, args []string) {
	logs, err := config.LoadConfig(configPath)
	if err != nil {
	fmt.Printf("Erreur de chargement : %v\n", err)
	os.Exit(1)
	}

	fmt.Println("Analyse en cours...")
	results := analyzer.AnalyzeLogs(logs)
	
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
}

