package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func createNewRequest(url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(username, password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return resp, err
}

func processRequest(resp *http.Response) ([]byte, error) {

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		println("Error while retrieving data from picqer")
	}

	return body, err
}

func GetPicklists() (Picklists, error) {
	return GetPicklistsByOffset(0)
}

func GetPicklistsByOffset(offset int) (Picklists, error) {

	_url := url + "/api/v1/picklists"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	resp, err := createNewRequest(_url)
	if err != nil {
		return nil, err
	}

	body, err := processRequest(resp)
	if err != nil {
		return nil, err
	}

	picklists := Picklists{}
	json.Unmarshal(body, &picklists)

	return picklists, err
}

func GetShipments(idpicklist int) (Shipments, error) {
	resp, err := createNewRequest(url + "/api/v1/picklists/" + strconv.Itoa(idpicklist) + "/shipments")
	if err != nil {
		return nil, err
	}

	body, err := processRequest(resp)

	shipments := Shipments{}
	json.Unmarshal(body, &shipments)

	return shipments, err
}

func GetPurchaseOrders() (PurchaseOrders, error) {
	return GetPurchaseOrdersByOffset(0)
}

func GetPurchaseOrdersByOffset(offset int) (PurchaseOrders, error) {

	_url := url + "/api/v1/purchaseorders"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	resp, err := createNewRequest(_url)
	if err != nil {
		return nil, err
	}

	body, err := processRequest(resp)
	if err != nil {
		return nil, err
	}

	purchaseOrders := PurchaseOrders{}
	json.Unmarshal(body, &purchaseOrders)

	return purchaseOrders, nil
}

func GetPurchaseOrder(idpurchaseorder int) (PurchaseOrder, error) {
	_url := url + "/api/v1/purchaseorders/" + strconv.Itoa(idpurchaseorder)

	resp, err := createNewRequest(_url)
	if err != nil {
		return PurchaseOrder{}, err
	}

	body, err := processRequest(resp)
	if err != nil {
		return PurchaseOrder{}, err
	}

	purchaseOrder := PurchaseOrder{}
	json.Unmarshal(body, &purchaseOrder)

	return purchaseOrder, nil
}

func GetSuppliers() (Suppliers, error) {
	return GetSuppliersByOffset(0)
}

func GetSuppliersByOffset(offset int) (Suppliers, error) {

	_url := url + "/api/v1/suppliers"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	resp, err := createNewRequest(_url)
	if err != nil {
		return nil, err
	}

	body, err := processRequest(resp)
	if err != nil {
		return nil, err
	}

	suppliers := Suppliers{}
	json.Unmarshal(body, &suppliers)

	return suppliers, nil
}

func GetProducts() (Products, error) {
	return GetProductsByOffset(0)
}

func GetProductsByOffset(offset int) (Products, error) {

	_url := url + "/api/v1/products"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	resp, err := createNewRequest(_url)
	if err != nil {
		return nil, err
	}

	body, err := processRequest(resp)
	if err != nil {
		return nil, err
	}

	products := Products{}
	json.Unmarshal(body, &products)

	return products, nil
}
