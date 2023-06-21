package shipstation

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/SmartPrintsInk/prettyfie"
	"github.com/SwiftPODEns/request"
)

func CreateOrder(order ShipStationOrder) (response ShipStationOrder, err error) {
	jsonData, err := json.Marshal(order)
	message := fmt.Sprintf("at Marshaling Orders\n%+v\n", order)
	check(err, message)
	requestPayload := request.RequestPayload{
		URL:    createEndpoint,
		Method: http.MethodPost,
		Headers: http.Header{
			"Host":          {"ssapi.shipstation.com"},
			"Authorization": {authorization},
			"Content-Type":  {"application/json"},
			"Accept":        {"application/json"},
		},
		QueryParams: nil,
		Payload:     jsonData,
	}
	responseBody, err := requestPayload.MakeHttpRequest()
	err = json.Unmarshal(responseBody, &response)
	fmt.Printf("Create order reponse %+v\n", prettyfie.Pretty(response))
	message = fmt.Sprintf("at Creating Orders\n%+v\n", order)
	check(err, message)
	return
}

func DeleteOrder(orderId int64) (response DeleteResponse, err error) {
	id := strconv.FormatInt(orderId, 10)
	uri := fmt.Sprintf("%s/%s", ordersEndpoint, id)
	requestPayload := request.RequestPayload{
		URL:    uri,
		Method: http.MethodDelete,
		Headers: http.Header{
			"Host":          {"ssapi.shipstation.com"},
			"Authorization": {authorization},
			"Content-Type":  {"application/json"},
			"Accept":        {"application/json"},
		},
		QueryParams: nil,
		Payload:     nil,
	}
	responseBody, err := requestPayload.MakeHttpRequest()
	check(err, "at Delete orderId "+id+"\n")
	json.Unmarshal(responseBody, &response)
	fmt.Printf("Delete order reponse %+v\n", prettyfie.Pretty(response))
	return
}

func MasrAs(shipstationMarkedItem ShipStationMarkAsShipped) (response ShipstationMarkShippedResponse, err error) {
	jsonData, err := json.Marshal(shipstationMarkedItem)
	id := strconv.FormatInt(shipstationMarkedItem.OrderID, 10)
	message := fmt.Sprintf("at Marshaling for marked as shipped orderId %s\n", id)
	check(err, message)
	requestPayload := request.RequestPayload{
		URL:    markAsShipped,
		Method: http.MethodPost,
		Headers: http.Header{
			"Host":          {"ssapi.shipstation.com"},
			"Authorization": {authorization},
			"Content-Type":  {"application/json"},
			"Accept":        {"application/json"},
		},
		QueryParams: nil,
		Payload:     jsonData,
	}
	responseBody, err := requestPayload.MakeHttpRequest()
	message = fmt.Sprintf("at Marked as shipped for orderId %s\n", id)
	check(err, message)
	err = json.Unmarshal(responseBody, &response)
	fmt.Printf("Mark As Shipped Reponse %+v\n", prettyfie.Pretty(response))
	return
}

func HoldDis(orderId int64, until string) (response ShipstationHoldResponse, err error) {
	shipstationHoldItem := ShipstationHoldItem{
		OrderID:       orderId,
		HoldUntilDate: until,
	}
	id := strconv.FormatInt(orderId, 10)
	jsonData, err := json.Marshal(shipstationHoldItem)
	message := fmt.Sprintf("at marshaling holding orderId %s\n", id)
	check(err, message)
	requestPayload := request.RequestPayload{
		URL:    holdEndpoint,
		Method: http.MethodPost,
		Headers: http.Header{
			"Host":          {"ssapi.shipstation.com"},
			"Authorization": {authorization},
			"Content-Type":  {"application/json"},
			"Accept":        {"application/json"},
		},
		QueryParams: nil,
		Payload:     jsonData,
	}
	responseBody, err := requestPayload.MakeHttpRequest()
	check(err, "at Holding orderId "+id+"\n")
	err = json.Unmarshal(responseBody, &response)
	fmt.Printf("Holding Reponse %+v\n", prettyfie.Pretty(response))
	return
}

func GetOrders(params url.Values) (orders []ShipStationOrder, err error) {
	var orderList ShipStationOrderList
	reqPayload := request.RequestPayload{
		URL:    ordersEndpoint,
		Method: http.MethodGet,
		Headers: http.Header{
			"Host":          {"ssapi.shipstation.com"},
			"Authorization": {authorization},
			"Content-Type":  {"application/json"},
			"Accept":        {"application/json"},
		},
		QueryParams: params,
		Payload:     nil,
	}
	resp, err := reqPayload.MakeHttpRequest()
	check(err, "Get ShipStation Order list")
	err = json.Unmarshal(resp, &orderList)
	log.Printf("ShipStation Response:\n%s\n", prettyfie.Pretty(orderList))
	orders = orderList.Orders
	if orderList.Page < orderList.Pages {
		params.Set("page", fmt.Sprintf("%d", orderList.Page+1))
		newOrders, _ := GetOrders(params)
		orders = append(orders, newOrders...)
	}
	return
}

func GetShipments(params url.Values) (shipments []Shipment, err error) {
	var list ShipmentList
	reqPayload := request.RequestPayload{
		URL:    shipmentEndpoint,
		Method: http.MethodGet,
		Headers: http.Header{
			"Host":          {"ssapi.shipstation.com"},
			"Authorization": {authorization},
			"Content-Type":  {"application/json"},
			"Accept":        {"application/json"},
		},
		QueryParams: params,
		Payload:     nil,
	}
	resp, err := reqPayload.MakeHttpRequest()
	check(err, "Get ShipStation Order list")
	err = json.Unmarshal(resp, &list)
	log.Printf("ShipStation Response:\n%s\n", prettyfie.Pretty(list))
	shipments = list.Shipments
	if list.Page < list.Pages {
		params.Set("page", fmt.Sprintf("%d", list.Page+1))
		newShipments, _ := GetShipments(params)
		shipments = append(shipments, newShipments...)
	}
	return
}
