package cli

import (
	"fmt"
	"strings"

	"github.com/felver-dev/contact-manager/internal/services"
)

type CLI struct {
	contactService *services.ContactService
}

func NewCLI(contactService *services.ContactService) *CLI {
	return &CLI{contactService: contactService}
}

func (cli *CLI) Run() error {
	for {
		cli.afficherMenu()
		choix := LireEntree()

		var error error

		switch choix {

		case "1":
			err = cli.ajouterContact()

		case "2":
			cli.listerContact()
		case "3":
			cli.rechercherContacts()
		case "4":
			cli.modifierContact()
		case "5":
			cli.supprimerContact()
		case "6":
			cli.afficherStatistiques()
		case "0":
			fmt.Println("\n Au revoir ! Vos contacts été sauvegardés.")
			return nil
		case "":
			continue
		default:
			fmt.Printf("Choix '%s' invalide. Choisissez entre 0 et 6.`n", choix)
		}

		if err != nil {
			fmt.Errorf("erreur : %s\n", err.Error())
		}

		fmt.Println("\nAppuyer sur Entrée pour continuer ...")
		LireEntree()
	}
}

func (cli *CLI) afficherMenu() {

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("\n GESTIONNAIRE DE CONTACTS")
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("1. Ajouter un contact")
	fmt.Println("2. Lister les contacts")
	fmt.Println("3. Rechercher des contacts")
	fmt.Println("4. Modifier un contact")
	fmt.Println("5. Supprimer un contact")
	fmt.Println("6. Afficher les statistiques")
	fmt.Println("0. Quitter")
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Votre choix")

}
