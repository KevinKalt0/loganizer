package analyzer

import (
	"github.com/KevinKalt0/loganizer/internal/config"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Result struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details,omitempty"`
}

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
			LogID:        log.ID,
			FilePath:     log.Path,
			Status:       "FAILED",
			Message:      "Fichier introuvable.",
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
