package api

type courier struct {
	Name    string
	Keyword string
}

var Courier = map[string]courier{
	"dhl":         {"DHL", "dhl"},
	"be-post":     {"BPost", "bpost"},
	"nl-post":     {"PostNL", "postnl"},
	"colis-prive": {"Colis Prive", "colis"},
	"fedex":       {"Fedex", "fedex"},
	"at-post":     {"AT Post", "atpost"},
	"ie-post":     {"IE Post", ""},
	"dpd":         {"DPD", "dpd"},
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
