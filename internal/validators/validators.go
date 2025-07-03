package validators

import "regexp"

// ValiderEmail vérifie si l'email a un format valide
func ValiderEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// ValiderTelephone vérifie si le téléphone a un format valide
func ValiderTelephone(telephone string) bool {
	// Nettoyer le numéro (enlever espaces, tirets, points)
	nettoye := regexp.MustCompile(`[\s\-\.]`).ReplaceAllString(telephone, "")

	// Vérifier qu'il ne contient que des chiffres et + optionnel au début
	re := regexp.MustCompile(`^\+?[0-9]{8,15}$`)
	return re.MatchString(nettoye)
}
