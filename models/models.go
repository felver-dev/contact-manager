package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Contact struct {
	ID        int       `json:"id"`
	Nom       string    `json:"nom"`
	Telephone string    `json:"telephone"`
	Email     string    `json:"email"`
	Cree      time.Time `json:"cree"`
	Modifie   time.Time `json:"modifie"`
}

type GestionnaireContacts struct {
	Contacts   []Contact
	ProchainID int
	Fichier    string
}

func (gc *GestionnaireContacts) ChargerFichier() error {

	if _, err := os.Stat(gc.Fichier); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(gc.Fichier)
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier : %v", err)
	}

	err = json.Unmarshal(data, &gc.Contacts)
	if err != nil {
		return fmt.Errorf("erreur lors de la désérialisation : %v", err)
	}

	for _, contact := range gc.Contacts {
		if contact.ID >= gc.ProchainID {
			gc.ProchainID = contact.ID + 1
		}
	}
	return nil
}
