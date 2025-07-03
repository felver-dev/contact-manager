package main

import (
	"log"

	"github.com/felver-dev/contact-manager/internal/cli"
	"github.com/felver-dev/contact-manager/internal/services"
	"github.com/felver-dev/contact-manager/internal/storage"
)

func main() {
	// Initialiser le stockage
	storage := storage.NewJSONStorage("contacts.json")

	// Initialiser le service
	contactService := services.NouveauGestionnaireContacts(storage)

	// Initialiser l'interface CLI
	cliApp := cli.NewCLI(contactService)

	// Lancer l'application
	if err := cliApp.Run(); err != nil {
		log.Fatal(err)
	}
}
