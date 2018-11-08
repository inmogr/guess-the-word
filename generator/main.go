package generator

func Generate(req string) (data []string) {
	switch req {
	case "colors":
		return []string{"Blue", "Red", "Yellow", "Green", "Pink"}
	case "directions":
		return []string{"Right", "Left", "Front", "Back"}
	case "names":
		return []string{"Morhaf", "Mohanad", "Majed", "Sami"}
	default:
		return []string{}
	}
}
