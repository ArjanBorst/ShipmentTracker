package api

// https://mholt.github.io/json-to-go/
//{"errors":[{"code":"tracker_not_found","message":"No Tracker found with this tracking number."}],"data":null}
type Errors struct {
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Data interface{} `json:"data"`
}