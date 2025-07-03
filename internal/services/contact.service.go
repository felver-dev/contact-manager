package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/felver-dev/contact-manager/internal/models"
	"github.com/felver-dev/contact-manager/internal/storage"
	"github.com/felver-dev/contact-manager/internal/validators"
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

func (cs *ContactService) LoadContacts() error {
	contacts, err := cs.storage.Load()

	if err != nil {
		return err
	}

	cs.contacts = contacts

	for _, contact := range cs.contacts {
		if contact.ID >= cs.prochainID {
			cs.prochainID = contact.ID + 1
		}
	}

	return nil
}

func (cs *ContactService) SaveContacts() error {

	return cs.storage.Save(cs.contacts)
}

func (cs *ContactService) AddContact(nom, telephone, email string) error {
	if !validators.ValiderTelephone(telephone) {
		return fmt.Errorf("format de téléphone invalid")
	}

	if !validators.ValiderEmail(email) {
		return fmt.Errorf("format d'email invalid")
	}

	for _, contact := range cs.contacts {
		if strings.EqualFold(contact.Email, email) {
			return fmt.Errorf("un contact avec cet email existe déjà (ID: %d)", contact.ID)
		}
	}

	maintenant := time.Now()
	nouveauContact := models.Contact{
		ID:        cs.prochainID,
		Nom:       nom,
		Telephone: telephone,
		Email:     email,
		Cree:      maintenant,
		Modifie:   maintenant,
	}

	cs.contacts = append(cs.contacts, nouveauContact)
	cs.prochainID++

	return cs.SaveContacts()
}

func (cs *ContactService) GetAllContacts() []models.Contact {
	return cs.contacts
}
