package api

import "time"

//https://mholt.github.io/json-to-go/
type TrackerResults struct {
	Data struct {
		Trackings []struct {
			Tracker struct {
				TrackerID         string    `json:"trackerId"`
				TrackingNumber    string    `json:"trackingNumber"`
				IsSubscribed      bool      `json:"isSubscribed"`
				ShipmentReference string    `json:"shipmentReference"`
				CreatedAt         time.Time `json:"createdAt"`
			} `json:"tracker"`
			Shipment struct {
				ShipmentID             string `json:"shipmentId"`
				StatusCode             string `json:"statusCode"`
				StatusCategory         string `json:"statusCategory"`
				StatusMilestone        string `json:"statusMilestone"`
				OriginCountryCode      string `json:"originCountryCode"`
				DestinationCountryCode string `json:"destinationCountryCode"`
				Delivery               struct {
					EstimatedDeliveryDate interface{} `json:"estimatedDeliveryDate"`
					Service               interface{} `json:"service"`
					SignedBy              interface{} `json:"signedBy"`
				} `json:"delivery"`
				TrackingNumbers []struct {
					Tn string `json:"tn"`
				} `json:"trackingNumbers"`
				Recipient struct {
					Name        interface{} `json:"name"`
					Address     string      `json:"address"`
					PostCode    string      `json:"postCode"`
					City        interface{} `json:"city"`
					Subdivision interface{} `json:"subdivision"`
				} `json:"recipient"`
			} `json:"shipment"`
			Events []struct {
				EventID             string      `json:"eventId"`
				TrackingNumber      string      `json:"trackingNumber"`
				EventTrackingNumber string      `json:"eventTrackingNumber"`
				Status              string      `json:"status"`
				OccurrenceDatetime  string	    `json:"occurrenceDatetime"`
				Order               interface{} `json:"order"`
				Datetime            time.Time   `json:"datetime"`
				HasNoTime           bool        `json:"hasNoTime"`
				UtcOffset           string      `json:"utcOffset"`
				Location            interface{} `json:"location"`
				SourceCode          string      `json:"sourceCode"`
				CourierCode         string      `json:"courierCode"`
				StatusCode          string      `json:"statusCode"`
				StatusCategory      string      `json:"statusCategory"`
				StatusMilestone     string      `json:"statusMilestone"`
			} `json:"events"`
			Statistics struct {
				Timestamps struct {
					InfoReceivedDatetime       time.Time   `json:"infoReceivedDatetime"`
					InTransitDatetime          time.Time   `json:"inTransitDatetime"`
					OutForDeliveryDatetime     time.Time   `json:"outForDeliveryDatetime"`
					FailedAttemptDatetime      interface{} `json:"failedAttemptDatetime"`
					AvailableForPickupDatetime interface{} `json:"availableForPickupDatetime"`
					ExceptionDatetime          interface{} `json:"exceptionDatetime"`
					DeliveredDatetime          time.Time   `json:"deliveredDatetime"`
				} `json:"timestamps"`
			} `json:"statistics"`
		} `json:"trackings"`
	} `json:"data"`
}