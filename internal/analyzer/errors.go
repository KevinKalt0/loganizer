package analyzer

import "fmt"

type FileNotFoundError struct {
	Path string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("Fichier introuvable : %s", e.Path)
}

type ParsingError struct {
	LogID string
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("Erreur de parsing pour le log : %s", e.LogID)
}
