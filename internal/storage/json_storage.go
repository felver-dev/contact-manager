package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/felver-dev/contact-manager/internal/models"
)

type Storage interface {
	Sauvegarder(contacts []models.Contact) error
	Charger() ([]models.Contact, error)
}

type JSONStorage struct {
	filename string
}

func NewJSONStorage(filename string) *JSONStorage {
	return &JSONStorage{filename: filename}
}

func (js *JSONStorage) Sauvegarder(contacts []models.Contact) error {
	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur lors de la sérialisation : %v", err)
	}

	err = os.WriteFile(js.filename, data, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture du fichier : %v", err)
	}

	return nil
}

func (js *JSONStorage) Charger() ([]models.Contact, error) {
	if _, err := os.Stat(js.filename); os.IsNotExist(err) {
		return []models.Contact{}, nil
	}

	data, err := os.ReadFile(js.filename)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier : %v", err)
	}

	var contacts []models.Contact
	err = json.Unmarshal(data, &contacts)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la désérialisation : %v", err)
	}

	return contacts, nil
}
