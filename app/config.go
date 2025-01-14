package app

type PatternConfig struct {
	LangKey         string
	TranslatedName  string
	LangFR          bool
	TranslatedLangs []string
}

type ColorWrapper struct {
	Colors map[string]string
}

func CreateColorWrapper(c PatternConfig) *ColorWrapper {
	w := ColorWrapper{make(map[string]string)}
	if c.LangFR {
		w.Colors["black"] = "Noir"
		w.Colors["blue"] = "Bleu"
		w.Colors["brown"] = "Brun"
		w.Colors["cyan"] = "Cyan"
		w.Colors["gray"] = "Gris"
		w.Colors["green"] = "Vert"
		w.Colors["light_blue"] = "Bleu Pâle"
		w.Colors["light_gray"] = "Gris Pâle"
		w.Colors["lime"] = "lime"
		w.Colors["magenta"] = "Magenta"
		w.Colors["orange"] = "Orange"
		w.Colors["pink"] = "Rose"
		w.Colors["purple"] = "Violet"
		w.Colors["red"] = "Rouge"
		w.Colors["white"] = "Blanc"
		w.Colors["yellow"] = "Jaune"
	} else {
		w.Colors["black"] = "Black"
		w.Colors["blue"] = "Blue"
		w.Colors["brown"] = "Brown"
		w.Colors["cyan"] = "Cyan"
		w.Colors["gray"] = "Gray"
		w.Colors["green"] = "Green"
		w.Colors["light_blue"] = "Light Blue"
		w.Colors["light_gray"] = "Light Gray"
		w.Colors["lime"] = "Lime"
		w.Colors["magenta"] = "Magenta"
		w.Colors["orange"] = "Orange"
		w.Colors["pink"] = "Pink"
		w.Colors["purple"] = "Purple"
		w.Colors["red"] = "Red"
		w.Colors["white"] = "White"
		w.Colors["yellow"] = "Yellow"
	}

	return &w
}
