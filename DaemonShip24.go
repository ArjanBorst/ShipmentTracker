package main

import (
	"fmt"
	"log"
	b2c "main/b2c"
	pApi "main/picqer/api"
	sApi "main/ship24/api"
	"reflect"
	"time"
)

const (
	firstRunPageSize int = 100
	pageSize         int = 100
)

type tracker struct {
	trackingNumber string
	courierCode    string
}

func DaemonShip24(pages int) {
	ticker := time.NewTicker(120 * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			fmt.Println("Adding latest track and trace numbers to Ship24 API")
			AddTrackAndTraceToShip24(pages)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func getTrackAndTrace(url string, idpicklist int) string {
	trackAndTrace, err := b2c.GetTrackAndTrace(url)
	switch {
	case err != nil:
		log.Printf("Failed to get shipments for picklist %d: %v", idpicklist, err)
		return ""
	case trackAndTrace == "":
		log.Printf("Track and Trace not yet available for picklist %d", idpicklist)
		return ""
	}

	return trackAndTrace
}

func GetPicklists(page int) (pApi.Picklists, error) {
	offset := page * pageSize
	retryAttempts := 3
	retryDelay := time.Second * 20

	var picklists pApi.Picklists
	var err error

	for attempt := 1; attempt <= retryAttempts; attempt++ {

		picklists, err = pApi.GetPicklistsByOffset(offset)
		if err == nil {
			return picklists, nil
		}

		message := err.Error()
		switch message {
		case "429 Too Many Requests":
			log.Printf("429 Too Many Requests, retrying after delay...")
			time.Sleep(retryDelay)
		default:
			log.Panic(err)
		}
	}

	log.Printf("Retries exhausted, unable to retrieve picklists: %v", err)
	return nil, err
}

func AddTrackAndTraceToShip24(pages int) {

	for pageNr := 0; pageNr < pages; pageNr++ {

		picklists, err := pApi.GetPicklistsByOffset(pageNr)
		if err != nil {
			log.Panic()
		}

		for _, picklist := range picklists {

			//shipments := pApi.Shipments{}
			shipments, err := pApi.GetShipments(picklist.Idpicklist)
			if err != nil {
				log.Printf("Failed to get picklists: %v", err)
				break
			}

			for _, shipment := range shipments {

				//courierInfo := sApi.Courier{}
				//courierInfo.trackingNumber = getTrackAndTrace(shipment.Tracktraceurl, picklist.Idpicklist)
				trackAndTrace := getTrackAndTrace(shipment.Tracktraceurl, picklist.Idpicklist)

				if trackAndTrace == "" {
					break
				}

				//println("Verify if tracking is already in Ship24 database before continue")
				shipmentRes, err := sApi.GetShipmentByTrackingNumber(trackAndTrace)
				if err != nil {
					log.Printf("Error while checking for tracking with id " + trackAndTrace)
					break
				}

				if (reflect.DeepEqual(shipmentRes, sApi.Shipment{})) {
					log.Printf("Add tracking with track and trace number: " + trackAndTrace)
					sApi.AddTracker(trackAndTrace,
						shipment.Trackingcode,
						picklist.Deliverycountry,
						picklist.Deliverycountry,
						picklist.Deliveryzipcode,
						picklist.Updated,
						picklist.Reference,
						picklist.Deliveryname)
				}

			}
		}
	}
}

/*
func AddToShip24(pages int) {


	shipments := pApi.Shipments{}
	shipments, _ = pApi.GetShipments(picklist.Idpicklist)

	for _, shipment := range shipments {

		trackAndTraceNumber := b2c.GetTrackAndTrace(shipment.Tracktraceurl)

		if trackAndTraceNumber == "" {
			//println("Track and Trace not found in URL: " + shipment.Tracktraceurl)
		} else {
			//println("Verify if tracking is already in Ship24 database before continue")
			res, err := sApi.GetShipmentByTrackingNumber(trackAndTraceNumber)
			if err != nil {
				fmt.Println("Error while checking for tracking with id " + trackAndTraceNumber)
			}

			if (reflect.DeepEqual(res, sApi.Shipment{})) {
				fmt.Println("Add tracking with track and trace number: " + trackAndTraceNumber)
				sApi.AddTracker(trackAndTraceNumber,
					shipment.Trackingcode,
					picklist.Deliverycountry,
					picklist.Deliverycountry,
					picklist.Deliveryzipcode,
					picklist.Updated,
					picklist.Reference,
					picklist.Deliveryname)
			}
		}
	}
}
*/
