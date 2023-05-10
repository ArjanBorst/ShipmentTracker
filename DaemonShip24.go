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

func AddTrackAndTraceToShip24(pages int) {

	for i := 0; i < pages; i++ {

		offset := i * 100
		picklists, err := pApi.GetPicklistsByOffset(offset)
		if err != nil {
			log.Panic()
		}

		for _, picklist := range picklists {

			shipments := pApi.Shipments{}
			shipments, err = pApi.GetShipments(picklist.Idpicklist)

			for _, shipment := range shipments {

				trackAndTrace, err := b2c.GetTrackAndTrace(shipment.Tracktraceurl)

				if err != nil || trackAndTrace == "" {

					break
				}

				if trackAndTrace != "" {
					//println("Verify if tracking is already in Ship24 database before continue")
					res, err := sApi.GetShipmentByTrackingNumber(trackAndTrace)
					if err != nil {
						fmt.Println("Error while checking for tracking with id " + trackAndTrace)
					}

					if (reflect.DeepEqual(res, sApi.Shipment{})) {
						fmt.Println("Add tracking with track and trace number: " + trackAndTrace)
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
