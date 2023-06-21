package shipstation

import (
	"os"
)

var authorization string

const (
	ordersEndpoint   string = "https://ssapi.shipstation.com/orders"
	holdEndpoint     string = "https://ssapi.shipstation.com/orders/holduntil"
	createEndpoint   string = "https://ssapi.shipstation.com/orders/createorder"
	shipmentEndpoint string = "https://ssapi.shipstation.com/shipments"
	markAsShipped    string = "https://ssapi.shipstation.com/orders/markasshipped"
)

func init() {
	apiKey := os.Getenv("ShipstationAPIKey")
	authorization = "Basic " + apiKey

}
