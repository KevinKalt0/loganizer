package analyzer

import "loganalyzer/internal/config"

type Result struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"` 
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details,omitempty"`
}

