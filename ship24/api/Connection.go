package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func createNewRequest(requestType, url string, data []byte) (*http.Response, error) {

	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+bearer)

	if err != nil {
		println("Error creating connection to Ship24")
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		println("Error creating connection to Ship24")
		return nil, err
	}

	return resp, err
}

func processRequest(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("Error while retrieving data from Ship24")
		return nil, err
	}

	return body, err
}

func GetShipmentByTrackingNumber(trackingNumber string) (Shipment, error) {

	resp, err := createNewRequest("GET", baseUrl+"/public/v1/trackers/search/"+trackingNumber+"/results", nil)
	if err != nil {
		return Shipment{}, err
	}

	body, err := processRequest(resp)
	if err != nil || resp.StatusCode == 404 {
		ship24Error := Errors{}
		json.Unmarshal(body, &ship24Error)
		return Shipment{}, err
	}

	shipment := Shipment{}
	json.Unmarshal(body, &shipment)

	return shipment, err
}

func createJSONRequest(trackingNumber, shipmentReference, originCountryCode, destinationCountryCode, destinationPostCode, dateISO8601, orderNumber string) []byte {

	var jsonStr = []byte(`{
		"trackingNumber": "` + trackingNumber + `",
			"shipmentReference": "` + shipmentReference + `",
			"originCountryCode": "` + originCountryCode + `",
			"destinationCountryCode": "` + destinationCountryCode + `",
			"destinationPostCode": "` + destinationPostCode + `",
			"courierCode": ["be-post"],
			"shippingDate": "` + dateISO8601 + `",
			"orderNumber": "` + orderNumber + `"
		}`)

	//"courierCode": ["be-post","dhl","nl-post","colis-prive","fedex","at-post","ie-post","dpd","dpd-pl","se-post"],

	return jsonStr
}

func createPostRequest(json []byte) {

	req, err := http.NewRequest("POST", baseUrl+"/public/v1/trackers", bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+bearer)

	if err != nil {
		println("Error creating connection to ship24")
		log.Fatal()
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}

func AddTracker(trackingNumber, shipmentReference, originCountryCode, destinationCountryCode, destinationPostCode, shippingDate, orderNumber, recipientName string) {

	date, _ := time.Parse("2006-01-02 15:04:05", shippingDate)
	dateISO8601 := date.Format("2006-01-02") + "T" + date.Format("15:04:05") + ".000Z"

	// courierCode can only be a list of max 11 items
	//courierCode := `"dhl","be-post","nl-post","colis-prive","fedex","at-post","ie-post","dpd","dpd-pl","se-post"`

	if destinationCountryCode == "FR" {
		trackingNumber = trackingNumber + destinationPostCode
	}

	if destinationCountryCode == "BE" {
		jsonStr := createJSONRequest(trackingNumber, shipmentReference, originCountryCode, destinationCountryCode, destinationPostCode, dateISO8601, orderNumber)
		createPostRequest(jsonStr)
	}

	//jsonStr := createJSONRequest(trackingNumber, shipmentReference, originCountryCode, destinationCountryCode, destinationPostCode, dateISO8601, orderNumber)

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}

func GetTrackers(page, limit int) (Trackers, error) {

	resp, err := createNewRequest("GET", "https://api.ship24.com/public/v1/trackers?page="+strconv.Itoa(page)+"&limit="+strconv.Itoa(limit), nil)
	if err != nil {
		return Trackers{}, err
	}

	body, err := processRequest(resp)

	if err != nil {
		return Trackers{}, err
	}

	trackers := Trackers{}
	json.Unmarshal(body, &trackers)

	return trackers, nil
}

func GetTracker(trackerId string) (Tracker, error) {

	resp, err := createNewRequest("GET", "https://api.ship24.com/public/v1/trackers/"+trackerId, nil)
	if err != nil {
		return Tracker{}, err
	}

	body, err := processRequest(resp)

	if err != nil {
		return Tracker{}, err
	}

	tracker := Tracker{}
	json.Unmarshal(body, &tracker)

	return tracker, nil
}

func GetTrackerResults(trackerId string) (TrackerResults, error) {

	resp, err := createNewRequest("GET", baseUrl+"/public/v1/trackers/"+trackerId+"/results", nil)
	if err != nil {
		return TrackerResults{}, err
	}

	body, err := processRequest(resp)

	if err != nil {
		return TrackerResults{}, err
	}

	trackerResults := TrackerResults{}
	json.Unmarshal(body, &trackerResults)

	return trackerResults, nil
}
