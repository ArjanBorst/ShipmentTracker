package main

//https://www.ipify.org/
//https://yourbasic.org/golang/convert-string-to-byte-slice/
//https://golangr.com/json-decode/
//https://pipedream.com/sources/dc_lVuW3Vp/configuration

// https://5118-195-240-36-217.eu.ngrok.io/picklist/
// ngrok http 8087
// curl https://succubus.picqer.com/api/v1/hooks --header "Content-Type:application/json" --header "Accept:application/json" --basic --user "1R2gFhpvpSZpY0sU6y4dROWDHDsvcgykF5oHfYCp7oF3B3ID" --user-agent "Picqer PHP API Client 0.9.10 (www.picqer.com)" --insecure -d "{\"name\":\"products free stock changed\",\"event\":\"products.free_stock_changed\",\"address\":\"https://6ecf-195-240-36-217.eu.ngrok.io/picklist/\"}
//https://www.sqlbi.com/articles/implementing-real-time-updates-in-power-bi-using-push-datasets-instead-of-directquery/
//https://radacad.com/power-bi-real-time-streaming-dataset
//https://www.youtube.com/watch?v=Te9bF01iqWM
import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type abstractPicklist struct {
	Id              string `json:"Id"`
	Deliverycountry string `json:"Deliverycountry"`
	Deliveryzipcode string `json:"Deliveryzipcode"`
	Updated         string `json:"Updated"`
	Reference       string `json:"Reference"`
	Deliveryname    string `json:"Deliveryname"`
}

var mux = &sync.RWMutex{}

// https://pkg.go.dev/container/list
func main() {

	start := time.Now()
	fmt.Println("Starting time is: ", start.String())

	mux.Lock()
	Load()
	mux.Unlock()

	AddTrackAndTraceToShip24(5)
	GetTrackingStats(10, mux)

	// go DaemonShip24(1)
	//go DaemonTrackingStats(1,mux)

	//AddTrackAndTraceToShip24(3)
	//GetTrackingStats(10, mux)

	mux.RLock()
	Save()
	mux.RUnlock()

	t := time.Now()
	fmt.Println("Ending date and time is: ", t.String())
	elapsed := t.Sub(start)
	println(elapsed)

	println("Starting REST API endpoints")

	http.HandleFunc("/proces/picklist", HTTPSDelivered)

	http.HandleFunc("/stats/delivered", HTTPSDelivered)
	http.HandleFunc("/stats/pending", HTTPPending)
	http.HandleFunc("/stats/notshipped", HTTPSNotShipped)
	http.HandleFunc("/resource/couriers", HTTPGetCouriers)
	http.HandleFunc("/resource/countries", HTTPGetCountries)

	http.ListenAndServe(":8080", nil)
}
