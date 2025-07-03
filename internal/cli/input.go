package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LireEntree() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func LireEntreeObligatoire(message string) string {
	for {
		fmt.Print(message)
		entree := LireEntree()

		if entree != "" {
			return entree
		}
		fmt.Println("Cette information est obligatoire")
	}
}
