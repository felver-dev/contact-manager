package models

import "time"

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
