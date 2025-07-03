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
	stockage   storage.Storage
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

func (cs *ContactService) SearchContacts(terme string) []models.Contact {
	var results []models.Contact
	terme = strings.ToLower(terme)

	for _, contact := range cs.contacts {
		if strings.Contains(strings.ToLower(contact.Nom), terme) ||
			strings.Contains(strings.ToLower(contact.Email), terme) {
			results = append(results, contact)
		}
	}

	return results
}

func (cs *ContactService) GetContactByID(id int) (*models.Contact, int) {
	for i, contact := range cs.contacts {
		if contact.ID == id {
			return &cs.contacts[i], i
		}
	}

	return nil, -1
}

func (cs *ContactService) UpdateContact(id int, nom, telephone, email string) error {
	contact, index := cs.GetContactByID(id)

	if contact == nil {
		return fmt.Errorf("aucun contact trouvé avec l'ID %d", id)
	}

	if nom != "" {
		contact.Nom = nom
	}

	if telephone != "" {
		if !validators.ValiderTelephone(telephone) {
			return fmt.Errorf("format de téléphone invalide")
		}
		contact.Telephone = telephone
	}

	if email != "" {
		if !validators.ValiderEmail(email) {
			return fmt.Errorf("format d'email invalide")
		}

		for _, c := range cs.contacts {
			if c.ID != contact.ID && strings.EqualFold(c.Email, email) {
				return fmt.Errorf("cet email est déjà utilisé par le contact ID %d", c.ID)
			}
		}

		contact.Email = email
	}

	contact.Modifie = time.Now()
	cs.contacts[index] = *contact

	return cs.SaveContacts()
}

func (cs *ContactService) DeleteContact(id int) error {
	_, index := cs.GetContactByID(id)

	if index == -1 {
		return fmt.Errorf("aucun contact trouvé avec l'ID %d", id)
	}

	cs.contacts = append(cs.contacts[:index], cs.contacts[index+1:]...)

	return cs.SaveContacts()
}

func (cs *ContactService) getStatistics() map[string]interface{} {

	stats := make(map[string]interface{})

	total := len(cs.contacts)
	stats["total"] = total

	if total == 0 {
		return stats
	}

	domaines := make(map[string]int)
	for _, contact := range cs.contacts {
		parts := strings.Split(contact.Email, "@")

		if len(parts) == 2 {
			domaine := strings.ToLower(parts[1])
			domaines[domaine]++
		}
	}

	stats["domaines"] = domaines

	plusRecent := cs.contacts[0]
	plusAncien := cs.contacts[0]

	for _, contact := range cs.contacts {

		if contact.Cree.After(plusRecent.Cree) {
			plusRecent = contact
		}
		if contact.Cree.Before(plusRecent.Cree) {
			plusAncien = contact
		}
	}

	stats["plus_recent"] = plusRecent
	stats["plus_ancien"] = plusAncien

	return stats
}
