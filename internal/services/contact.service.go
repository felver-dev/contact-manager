package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/felver-dev/contact-manager/internal/models"
	"github.com/felver-dev/contact-manager/internal/storage"
	"github.com/felver-dev/contact-manager/internal/validators"
)

type GestionnaireContacts struct {
	contacts   []models.Contact
	prochainID int
	stockage   storage.Storage
}

func NouveauGestionnaireContacts(stockage storage.Storage) *GestionnaireContacts {
	gc := &GestionnaireContacts{
		contacts:   make([]models.Contact, 0),
		prochainID: 1,
		stockage:   stockage,
	}

	gc.ChargerContacts()
	return gc
}

// AjouterContact ajoute un nouveau contact
func (gc *GestionnaireContacts) AjouterContact(nom, telephone, email string) error {
	// Validation
	if !validators.ValiderTelephone(telephone) {
		return fmt.Errorf("format de téléphone invalide")
	}

	if !validators.ValiderEmail(email) {
		return fmt.Errorf("format d'email invalide")
	}

	// Vérifier si le contact existe déjà
	for _, contact := range gc.contacts {
		if strings.EqualFold(contact.Email, email) {
			return fmt.Errorf("un contact avec cet email existe déjà (ID: %d)", contact.ID)
		}
	}

	// Créer le nouveau contact
	maintenant := time.Now()
	nouveauContact := models.Contact{
		ID:        gc.prochainID,
		Nom:       nom,
		Telephone: telephone,
		Email:     email,
		Cree:      maintenant,
		Modifie:   maintenant,
	}

	gc.contacts = append(gc.contacts, nouveauContact)
	gc.prochainID++

	return gc.stockage.Sauvegarder(gc.contacts)
}

// ListerContacts retourne tous les contacts
func (gc *GestionnaireContacts) ListerContacts() []models.Contact {
	return gc.contacts
}

// RechercherContacts recherche des contacts par terme
func (gc *GestionnaireContacts) RechercherContacts(terme string) []models.Contact {
	var resultats []models.Contact
	terme = strings.ToLower(terme)

	for _, contact := range gc.contacts {
		if strings.Contains(strings.ToLower(contact.Nom), terme) ||
			strings.Contains(strings.ToLower(contact.Email), terme) {
			resultats = append(resultats, contact)
		}
	}

	return resultats
}

// TrouverContactParID trouve un contact par son ID
func (gc *GestionnaireContacts) TrouverContactParID(id int) (*models.Contact, int) {
	for i, contact := range gc.contacts {
		if contact.ID == id {
			return &gc.contacts[i], i
		}
	}
	return nil, -1
}

// ModifierContact modifie un contact existant
func (gc *GestionnaireContacts) ModifierContact(id int, nouveauNom, nouveauTelephone, nouvelEmail string) error {
	contact, index := gc.TrouverContactParID(id)
	if contact == nil {
		return fmt.Errorf("aucun contact trouvé avec l'ID %d", id)
	}

	if nouveauNom != "" {
		contact.Nom = nouveauNom
	}

	if nouveauTelephone != "" {
		if !validators.ValiderTelephone(nouveauTelephone) {
			return fmt.Errorf("format de téléphone invalide")
		}
		contact.Telephone = nouveauTelephone
	}

	if nouvelEmail != "" {
		if !validators.ValiderEmail(nouvelEmail) {
			return fmt.Errorf("format d'email invalide")
		}

		for _, c := range gc.contacts {
			if c.ID != contact.ID && strings.EqualFold(c.Email, nouvelEmail) {
				return fmt.Errorf("cet email est déjà utilisé par le contact ID %d", c.ID)
			}
		}
		contact.Email = nouvelEmail
	}

	contact.Modifie = time.Now()
	gc.contacts[index] = *contact

	return gc.stockage.Sauvegarder(gc.contacts)
}

// SupprimerContact supprime un contact
func (gc *GestionnaireContacts) SupprimerContact(id int) error {
	_, index := gc.TrouverContactParID(id)
	if index == -1 {
		return fmt.Errorf("aucun contact trouvé avec l'ID %d", id)
	}

	gc.contacts = append(gc.contacts[:index], gc.contacts[index+1:]...)
	return gc.stockage.Sauvegarder(gc.contacts)
}

// AfficherStatistiques retourne les statistiques des contacts
func (gc *GestionnaireContacts) AfficherStatistiques() map[string]interface{} {
	stats := make(map[string]interface{})

	total := len(gc.contacts)
	stats["total"] = total

	if total == 0 {
		return stats
	}

	// Compter les domaines email
	domaines := make(map[string]int)
	for _, contact := range gc.contacts {
		parts := strings.Split(contact.Email, "@")
		if len(parts) == 2 {
			domaine := strings.ToLower(parts[1])
			domaines[domaine]++
		}
	}
	stats["domaines"] = domaines

	// Contact le plus récent et le plus ancien
	plusRecent := gc.contacts[0]
	plusAncien := gc.contacts[0]

	for _, contact := range gc.contacts {
		if contact.Cree.After(plusRecent.Cree) {
			plusRecent = contact
		}
		if contact.Cree.Before(plusAncien.Cree) {
			plusAncien = contact
		}
	}

	stats["plus_recent"] = plusRecent
	stats["plus_ancien"] = plusAncien

	return stats
}

// ChargerContacts charge les contacts depuis le stockage
func (gc *GestionnaireContacts) ChargerContacts() error {
	contacts, err := gc.stockage.Charger()
	if err != nil {
		return err
	}

	gc.contacts = contacts

	for _, contact := range gc.contacts {
		if contact.ID >= gc.prochainID {
			gc.prochainID = contact.ID + 1
		}
	}

	return nil
}
