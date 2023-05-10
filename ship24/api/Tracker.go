package api

import "time"

type Tracker struct {
	Tracker []struct {
		TrackerID         string    `json:"trackerId"`
		TrackingNumber    string    `json:"trackingNumber"`
		ShipmentReference string    `json:"shipmentReference"`
		IsSubscribed      bool      `json:"isSubscribed"`
		CreatedAt         time.Time `json:"createdAt"`
	} `json:"trackers"`
}
