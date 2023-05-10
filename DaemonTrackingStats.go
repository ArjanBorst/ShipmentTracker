package main

import (
	"fmt"
	"log"
	b2c "main/b2c"
	pApi "main/picqer/api"
	sApi "main/ship24/api"
	"math"
	"sync"
	"time"
)

type Shipment struct {
	ShipmentCreated string  `json:"shipmentcreated"`
	Country         string  `json:"country"`
	DurationInDays  float64 `json:"durationindays"`
	Courier         string  `json:"courier"`
	TrackingStatus  string  `json:"currenttrackingstatus"`
	TrackAndTrace   string  `json:"trackandtrace"`
}

type ShipmentNotShipped struct {
	ShipmentCreated string  `json:"shipmentcreated"`
	Country         string  `json:"country"`
	DurationInDays  float64 `json:"durationindays"`
}

/*

type Courier struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

type Country struct {
	Name string `json:"name"`
}

type TrackingStatus struct {
	Name string `json:"name"`
}*/

var mDelivered map[int]Shipment
var mPending map[int]Shipment
var mNotShipped map[int]ShipmentNotShipped

//var mCountry map[string]Country
//var mCourier map[string]Courier
//var mStatus map[string]TrackingStatus

func DaemonTrackingStats(pages int, mux *sync.RWMutex) {

	//PrepareData()

	//mDelivered = make(map[int]Shipment)
	//mPending = make(map[int]Shipment)
	//mNotShipped = make(map[int]ShipmentNotShipped)

	ticker := time.NewTicker(120 * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			fmt.Println("Getting latest statisticks for Dashboard")
			GetTrackingStats(pages, mux)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func GetTrackingStats(pages int, mux *sync.RWMutex) {

	for i := 1; i < pages; i++ {

		offset := i * 100
		picklists, err := pApi.GetPicklistsByOffset(offset)
		if err != nil {
			log.Panic()
		}

		for _, picklist := range picklists {

			shipments, _ := pApi.GetShipments(picklist.Idpicklist)

			for _, shipment := range shipments {

				if mDelivered[picklist.Idpicklist].TrackingStatus != "Afgeleverd" {

					trackAndTraceNumber, _ := b2c.GetTrackAndTrace(shipment.Tracktraceurl)
					created, _ := time.Parse("2006-01-02 15:04:05", shipment.Created)

					if trackAndTraceNumber == "" {

						mux.Lock()
						AddToNotShipped(picklist.Idpicklist, created, picklist.Deliverycountry)
						mux.Unlock()

					} else {
						sShipment := sApi.Shipment{}
						sShipment, _ = sApi.GetShipmentByTrackingNumber(trackAndTraceNumber)

						for _, trackings := range sShipment.Data.Trackings {

							if len(trackings.Events) > 0 {

								tracking := Shipment{
									created.Format("02-Jan-2006"),
									sApi.Country[picklist.Deliverycountry].Name,
									math.Round((trackings.Events[0].Datetime.Sub(created).Hours()/24)*100) / 100,
									sApi.Courier[trackings.Events[0].CourierCode].Name,
									sApi.Status[trackings.Events[0].StatusMilestone].Description,
									trackAndTraceNumber}

								mux.Lock()
								AddToShipped(picklist.Idpicklist, created, tracking)
								mux.Unlock()

							} else {
								mux.Lock()
								AddToNotShipped(picklist.Idpicklist, created, picklist.Deliverycountry)
								mux.Unlock()
							}

							//println(tracking.TrackingStatus)

							//powerbi.PushTracker(customerName,url,deliveryCountry,courierCode,statusMilestone,"",pShipmentDate,durationInDays)
						}
					}
				}
			}
		}
	}
}

func AddToShipped(idpicklist int, created time.Time, tracking Shipment) {

	if tracking.TrackingStatus == "Afgeleverd" {
		mDelivered[idpicklist] = tracking

		delete(mPending, idpicklist)
	} else {
		mPending[idpicklist] = tracking

		delete(mDelivered, idpicklist)
	}

	delete(mNotShipped, idpicklist)
}

func AddToNotShipped(idpicklist int, created time.Time, country string) {
	durationInDays := math.Round((time.Now().Sub(created).Hours()/24)*100) / 100

	notShipped := ShipmentNotShipped{
		created.Format("02-Jan-2006"),
		sApi.Country[country].Name,
		durationInDays}

	mNotShipped[idpicklist] = notShipped
}

/*
func PrepareData() {

	mCountry = make(map[string]Country)
	mCountry["NL"] = Country{"Nederland"}
	mCountry["BE"] = Country{"Belgie"}
	mCountry["DE"] = Country{"Duitsland"}
	mCountry["AT"] = Country{"Oostenrijk"}
	mCountry["US"] = Country{"Verenigde Staten"}
	mCountry["FR"] = Country{"Frankrijk"}
	mCountry["CH"] = Country{"Zwitserland"}
	mCountry["IE"] = Country{"Ierland"}
	mCountry["FI"] = Country{"Finland"}
	mCountry["NO"] = Country{"Noorwegen"}
	mCountry["SE"] = Country{"Sweden"}
	mCountry["PL"] = Country{"Polen"}
	mCountry["DK"] = Country{"Denemarken"}
	mCountry["PT"] = Country{"Portugal"}
	mCountry["AU"] = Country{"Australie"}
	mCountry["CA"] = Country{"Canada"}
	mCountry["EE"] = Country{"Estland"}
	mCountry["ES"] = Country{"Spanje"}
	mCountry["IT"] = Country{"Italie"}
	mCountry["NZ"] = Country{"Nieuw Zeeland"}
	mCountry["CZ"] = Country{"Tsjechië "}
	mCountry["GR"] = Country{"Griekenland"}
	mCountry["HR"] = Country{"Kroatië"}
	mCountry["HU"] = Country{"Hongarije"}
	mCountry["IS"] = Country{"IJsland"}
	mCountry["LU"] = Country{"Luxemburg"}
	mCountry["RO"] = Country{"Roemenië"}

	mCourier = make(map[string]Courier)
	mCourier["dhl"] = Courier{"DHL", ""}
	mCourier["be-post"] = Courier{"BPost", ""}
	mCourier["nl-post"] = Courier{"PostNL", ""}
	mCourier["colis-prive"] = Courier{"Colis Prive", ""}
	mCourier["fedex"] = Courier{"Fedex", ""}
	mCourier["at-post"] = Courier{"AT Post", ""}
	mCourier["ie-post"] = Courier{"IE Post", ""}
	mCourier["dpd"] = Courier{"DPD", ""}
	mCourier["dpd-pl"] = Courier{"DPD Poland", ""}
	mCourier["se-post"] = Courier{"SE Post", ""}
	mCourier["pt-post"] = Courier{"CTT", ""}
	mCourier["no-post"] = Courier{"Norway Post", ""}
	mCourier["fi-post"] = Courier{"Posti", ""}
	mCourier["us-post"] = Courier{"USPS", ""}
	mCourier["es-post"] = Courier{"Correos", ""}
	mCourier["ch-post"] = Courier{"Swiss Post", ""}
	mCourier["is-post"] = Courier{"Posturinn", ""}
	mCourier["it-post"] = Courier{"Post of Italy", ""}
	mCourier["fr-post"] = Courier{"La Poste", ""}
	mCourier["ca-post"] = Courier{"Canada Post", ""}
	mCourier["ht-post"] = Courier{"Hrvatska Posta", ""}
	mCourier["sda-it"] = Courier{"SDA", ""}
	mCourier["is-post"] = Courier{"Posturinn", ""}

	mStatus = make(map[string]TrackingStatus)
	mStatus["delivered"] = TrackingStatus{"Afgeleverd"}
	mStatus["in_transit"] = TrackingStatus{"Onderweg"}
	mStatus["out_for_delivery"] = TrackingStatus{"Onderweg naar Klant"}
	mStatus["available_for_pickup"] = TrackingStatus{"Ligt op Postkantoor"}
	mStatus["info_received"] = TrackingStatus{"Aangemeld bij Vervoeder"}
	mStatus["failed_attempt"] = TrackingStatus{"Aflevering mislukt"}
}
*/
