package models

import (
	"fmt"
	"strings"
	"time"
)

type Contact struct {
	ID        int       `json:"id"`
	Nom       string    `json:"Nom"`
	Telephone string    `json:"telephone"`
	Email     string    `json:"email"`
	Cree      time.Time `json:"cree"`
	Modifie   time.Time `json:"modifie"`
}

func (c Contact) String() string {
	return fmt.Sprintf("ID: %d | %s | %s | %s", c.ID, c.Nom, c.Telephone, c.Email)
}

func (c Contact) AfficherDetails() {
	fmt.Printf("┌%s┐\n", strings.Repeat("─", 50))
	fmt.Printf("│ Contact #%d%s│\n", c.ID, strings.Repeat(" ", 50-len(fmt.Sprintf(" Contact #%d", c.ID))))
	fmt.Printf("├%s┤\n", strings.Repeat("─", 50))
	fmt.Printf("│ Nom       : %-35s │\n", c.Nom)
	fmt.Printf("│ Téléphone : %-35s │\n", c.Telephone)
	fmt.Printf("│ Email     : %-35s │\n", c.Email)
	fmt.Printf("│ Créé      : %-35s │\n", c.Cree.Format("02/01/2006 15:04:05"))
	fmt.Printf("│ Modifié   : %-35s │\n", c.Modifie.Format("02/01/2006 15:04:05"))
	fmt.Printf("└%s┘\n", strings.Repeat("─", 50))
}
