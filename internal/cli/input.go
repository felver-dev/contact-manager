package cli

import (
	"bufio"
	"os"
	"strings"
)

func lireEntree() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
