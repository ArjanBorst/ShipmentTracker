package api

import "time"

// https://mholt.github.io/json-to-go/

type CreateTracker struct {
	Data struct {
		Trackings []struct {
			Tracker struct {
				TrackerID         string    `json:"trackerId"`
				TrackingNumber    string    `json:"trackingNumber"`
				CreatedAt         time.Time `json:"createdAt"`
				IsSubscribed      bool      `json:"isSubscribed"`
				ShipmentReference string    `json:"shipmentReference"`
			} `json:"tracker"`
			Shipment struct {
				ShipmentID     string      `json:"shipmentId"`
				StatusCode     interface{} `json:"statusCode"`
				StatusCategory interface{} `json:"statusCategory"`
				Delivery       struct {
					EstimatedDeliveryDate interface{} `json:"estimatedDeliveryDate"`
					Service               interface{} `json:"service"`
					SignedBy              interface{} `json:"signedBy"`
				} `json:"delivery"`
				TrackingNumbers []struct {
					Tn string `json:"tn"`
				} `json:"trackingNumbers"`
				Recipient struct {
					Name        interface{} `json:"name"`
					Address     interface{} `json:"address"`
					PostCode    interface{} `json:"postCode"`
					City        interface{} `json:"city"`
					Subdivision interface{} `json:"subdivision"`
				} `json:"recipient"`
			} `json:"shipment"`
			Events []interface{} `json:"events"`
		} `json:"trackings"`
	} `json:"data"`
}