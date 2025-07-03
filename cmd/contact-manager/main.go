package main

import (
	"fmt"
	"log"

	"github.com/felver-dev/contact-manager/internal/cli"
	"github.com/felver-dev/contact-manager/internal/services"
	"github.com/felver-dev/contact-manager/internal/storage"
)

func main() {

	storage := storage.JSONStorage("contacts.json")

	contactService := services.NewContactService(storage)

	cliApp := cli.NewCLI(contactService)

	fmt.Println("Bienvenu dans les gestionnaire de contacts!")

	if err := cliApp.Run(); err != nil {
		log.Fatal(err)
	}
}
