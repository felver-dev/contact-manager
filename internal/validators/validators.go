package validators

import "regexp"

func ValiderEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func ValiderTelephone(telephone string) bool {
	nettoye := regexp.MustCompile(`[\s-\.]`).ReplaceAllString(telephone, "")
	re := regexp.MustCompile(`^\+?[0-9]{8, 15}`)

	return re.MatchString(nettoye)
}
