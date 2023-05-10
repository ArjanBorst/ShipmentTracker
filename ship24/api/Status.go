package api

type trackingStatus struct {
	Description string
}

var Status = map[string]trackingStatus{
	"delivered":            {"Afgeleverd"},
	"in_transit":           {"Onderweg"},
	"out_for_delivery":     {"Onderweg naar Klant"},
	"available_for_pickup": {"Ligt op Postkantoor"},
	"info_received":        {"Aangemeld bij Vervoeder"},
	"failed_attempt":       {"Aflevering mislukt"},
}
