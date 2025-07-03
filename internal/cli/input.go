package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LireEntree lit une ligne d'entrée utilisateur
func LireEntree() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// LireEntreeObligatoire lit une entrée qui ne peut pas être vide
func LireEntreeObligatoire(message string) string {
	for {
		fmt.Print(message)
		entree := LireEntree()
		if entree != "" {
			return entree
		}
		fmt.Println("❌ Cette information est obligatoire.")
	}
}
