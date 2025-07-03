package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/felver-dev/contact-manager/internal/models"
	"github.com/felver-dev/contact-manager/internal/services"
)

type CLI struct {
	contactService *services.GestionnaireContacts
}

func NewCLI(contactService *services.GestionnaireContacts) *CLI {
	return &CLI{contactService: contactService}
}

func (cli *CLI) Run() error {
	fmt.Println("🎉 Bienvenue dans le Gestionnaire de Contacts !")
	fmt.Println("📁 Fichier de sauvegarde : contacts.json")

	for {
		cli.afficherMenu()
		choix := LireEntree()

		var err error
		switch choix {
		case "1":
			err = cli.ajouterContact()
		case "2":
			cli.listerContacts()
		case "3":
			cli.rechercherContacts()
		case "4":
			err = cli.modifierContact()
		case "5":
			err = cli.supprimerContact()
		case "6":
			cli.afficherStatistiques()
		case "0":
			fmt.Println("\n👋 Au revoir ! Vos contacts ont été sauvegardés.")
			return nil
		case "":
			continue
		default:
			fmt.Printf("❌ Choix '%s' invalide. Choisissez entre 0 et 6.\n", choix)
		}

		if err != nil {
			fmt.Printf("\n❌ Erreur : %s\n", err.Error())
		}

		fmt.Println("\nAppuyez sur Entrée pour continuer...")
		LireEntree()
	}
}

func (cli *CLI) afficherMenu() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("👥  GESTIONNAIRE DE CONTACTS")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("1. Ajouter un contact")
	fmt.Println("2. Lister tous les contacts")
	fmt.Println("3. Rechercher des contacts")
	fmt.Println("4. Modifier un contact")
	fmt.Println("5. Supprimer un contact")
	fmt.Println("6. Afficher les statistiques")
	fmt.Println("0. Quitter")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Print("Votre choix : ")
}

func (cli *CLI) ajouterContact() error {
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("➕  AJOUTER UN NOUVEAU CONTACT")
	fmt.Println(strings.Repeat("=", 40))

	nom := LireEntreeObligatoire("Nom complet : ")

	var telephone string
	for {
		fmt.Print("Numéro de téléphone : ")
		telephone = LireEntree()
		if telephone == "" {
			fmt.Println("❌ Le numéro de téléphone est obligatoire.")
			continue
		}
		break
	}

	var email string
	for {
		fmt.Print("Adresse email : ")
		email = LireEntree()
		if email == "" {
			fmt.Println("❌ L'adresse email est obligatoire.")
			continue
		}
		break
	}

	if err := cli.contactService.AjouterContact(nom, telephone, email); err != nil {
		return err
	}

	fmt.Printf("\n✅ Contact '%s' ajouté avec succès !\n", nom)
	return nil
}

func (cli *CLI) listerContacts() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("📞  LISTE DES CONTACTS\n")
	fmt.Println(strings.Repeat("=", 60))

	contacts := cli.contactService.ListerContacts()

	if len(contacts) == 0 {
		fmt.Println("Aucun contact enregistré.")
		return
	}

	fmt.Printf("│ %-3s │ %-20s │ %-15s │ %-20s │\n", "ID", "Nom", "Téléphone", "Email")
	fmt.Printf("├%s┼%s┼%s┼%s┤\n",
		strings.Repeat("─", 5),
		strings.Repeat("─", 22),
		strings.Repeat("─", 17),
		strings.Repeat("─", 22))

	for _, contact := range contacts {
		nom := contact.Nom
		if len(nom) > 20 {
			nom = nom[:17] + "..."
		}
		email := contact.Email
		if len(email) > 20 {
			email = email[:17] + "..."
		}

		fmt.Printf("│ %-3d │ %-20s │ %-15s │ %-20s │\n",
			contact.ID, nom, contact.Telephone, email)
	}
	fmt.Printf("└%s┴%s┴%s┴%s┘\n",
		strings.Repeat("─", 5),
		strings.Repeat("─", 22),
		strings.Repeat("─", 17),
		strings.Repeat("─", 22))
}

func (cli *CLI) rechercherContacts() {
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("🔍  RECHERCHER DES CONTACTS")
	fmt.Println(strings.Repeat("=", 40))

	fmt.Print("Entrez le terme de recherche (nom ou email) : ")
	terme := LireEntree()

	if terme == "" {
		fmt.Println("❌ Veuillez saisir un terme de recherche.")
		return
	}

	resultats := cli.contactService.RechercherContacts(terme)

	fmt.Printf("\n🎯 %d résultat(s) trouvé(s) pour '%s':\n", len(resultats), terme)
	if len(resultats) == 0 {
		fmt.Println("Aucun contact correspondant.")
		return
	}

	for _, contact := range resultats {
		fmt.Printf("\n%s\n", contact.String())
	}
}

func (cli *CLI) modifierContact() error {
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("✏️   MODIFIER UN CONTACT")
	fmt.Println(strings.Repeat("=", 40))

	fmt.Print("ID du contact à modifier : ")
	idStr := LireEntree()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("ID invalide : %s", idStr)
	}

	contact, _ := cli.contactService.TrouverContactParID(id)
	if contact == nil {
		return fmt.Errorf("aucun contact trouvé avec l'ID %d", id)
	}

	fmt.Println("\nContact actuel :")
	contact.AfficherDetails()

	fmt.Println("\nLaissez vide pour conserver la valeur actuelle.")

	fmt.Printf("Nouveau nom (%s) : ", contact.Nom)
	nouveauNom := LireEntree()

	fmt.Printf("Nouveau téléphone (%s) : ", contact.Telephone)
	nouveauTel := LireEntree()

	fmt.Printf("Nouvel email (%s) : ", contact.Email)
	nouvelEmail := LireEntree()

	if err := cli.contactService.ModifierContact(id, nouveauNom, nouveauTel, nouvelEmail); err != nil {
		return err
	}

	fmt.Printf("\n✅ Contact ID %d modifié avec succès !\n", id)
	return nil
}

func (cli *CLI) supprimerContact() error {
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("🗑️   SUPPRIMER UN CONTACT")
	fmt.Println(strings.Repeat("=", 40))

	fmt.Print("ID du contact à supprimer : ")
	idStr := LireEntree()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("ID invalide : %s", idStr)
	}

	contact, _ := cli.contactService.TrouverContactParID(id)
	if contact == nil {
		return fmt.Errorf("aucun contact trouvé avec l'ID %d", id)
	}

	fmt.Println("\nContact à supprimer :")
	contact.AfficherDetails()

	fmt.Print("\n⚠️  Êtes-vous sûr de vouloir supprimer ce contact ? (oui/non) : ")
	confirmation := strings.ToLower(LireEntree())

	if confirmation != "oui" && confirmation != "o" && confirmation != "yes" && confirmation != "y" {
		fmt.Println("❌ Suppression annulée.")
		return nil
	}

	nom := contact.Nom

	if err := cli.contactService.SupprimerContact(id); err != nil {
		return err
	}

	fmt.Printf("✅ Contact '%s' (ID: %d) supprimé avec succès !\n", nom, id)
	return nil
}

func (cli *CLI) afficherStatistiques() {
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("📊  STATISTIQUES")
	fmt.Println(strings.Repeat("=", 40))

	stats := cli.contactService.AfficherStatistiques()

	total := stats["total"].(int)
	fmt.Printf("Nombre total de contacts : %d\n", total)

	if total == 0 {
		return
	}

	if domaines, ok := stats["domaines"].(map[string]int); ok {
		fmt.Println("\nDomaines email les plus utilisés :")
		for domaine, count := range domaines {
			fmt.Printf("  %s : %d contact(s)\n", domaine, count)
		}
	}

	if plusRecent, ok := stats["plus_recent"].(models.Contact); ok {
		fmt.Printf("\nContact le plus récent : %s (%s)\n",
			plusRecent.Nom, plusRecent.Cree.Format("02/01/2006"))
	}

	if plusAncien, ok := stats["plus_ancien"].(models.Contact); ok {
		fmt.Printf("Contact le plus ancien : %s (%s)\n",
			plusAncien.Nom, plusAncien.Cree.Format("02/01/2006"))
	}
}
