package services

import "github.com/felver-dev/contact-manager/models"

func NouveauGestionnaire(fichier string) *models.GestionnaireContacts {

	gc := &models.GestionnaireContacts{
		Contacts:   make([]models.Contact, 0),
		ProchainID: 1,
		Fichier:    fichier,
	}

	gc.ChargerFichier()

	return gc
}
