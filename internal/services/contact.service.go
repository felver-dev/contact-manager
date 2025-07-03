package services

import (
	"fmt"

	"github.com/felver-dev/contact-manager/internal/models"
	"github.com/felver-dev/contact-manager/internal/storage"
)

type ContactService struct {
	contacts   []models.Contact
	prochainID int
	storage    storage.Storage
}

func NewContactService(storage storage.Storage) *ContactService {
	cs := &ContactService{
		contacts:   make([]models.Contact, 0),
		prochainID: 1,
		storage:    storage,
	}

	if err := cs.LoadContacts(); err != nil {
		fmt.Printf("Erreur lors du chargement : %v\n", err)
	}

	return cs
}
