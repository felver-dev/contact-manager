package cli

import "github.com/felver-dev/contact-manager/internal/services"

type CLI struct {
	contactService *services.ContactService
}

func NewCLI(contactService *services.ContactService) *CLI {
	return &CLI{contactService: contactService}
}
