package api

import "time"

// https://mholt.github.io/json-to-go/
	
type Trackers struct {
	Data struct {
		Trackers []struct {
			TrackerID         string    `json:"trackerId"`
			TrackingNumber    string    `json:"trackingNumber"`
			ShipmentReference string    `json:"shipmentReference"`
			IsSubscribed      bool      `json:"isSubscribed"`
			CreatedAt         time.Time `json:"createdAt"`
		} `json:"trackers"`
	} `json:"data"`
}