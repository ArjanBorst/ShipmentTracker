package api

type courier struct {
	Name     string
	Location string
}

var Courier = map[string]courier{
	"dhl":         {"DHL", ""},
	"be-post":     {"BPost", ""},
	"nl-post":     {"PostNL", ""},
	"colis-prive": {"Colis Prive", ""},
	"fedex":       {"Fedex", ""},
	"at-post":     {"AT Post", ""},
	"ie-post":     {"IE Post", ""},
	"dpd":         {"DPD", ""},
	"dpd-pl":      {"DPD Poland", ""},
	"se-post":     {"SE Post", ""},
	"pt-post":     {"CTT", ""},
	"no-post":     {"Norway Post", ""},
	"fi-post":     {"Posti", ""},
	"us-post":     {"USPS", ""},
	"es-post":     {"Correos", ""},
	"ch-post":     {"Swiss Post", ""},
	"is-post":     {"Posturinn", ""},
	"it-post":     {"Post of Italy", ""},
	"fr-post":     {"La Poste", ""},
	"ca-post":     {"Canada Post", ""},
	"ht-post":     {"Hrvatska Posta", ""},
	"sda-it":      {"SDA", ""},
}
