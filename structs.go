package shipstation

type DeleteResponse struct {
	Success bool   `json:"success" bson:"success,omitempty"`
	Message string `json:"message" bson:"message,omitempty"`
}

// BillTo and ShipTo
type Address struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Company     string `json:"company,omitempty" bson:"company,omitempty"`
	Street1     string `json:"street1,omitempty" bson:"street1,omitempty"`
	Street2     string `json:"street2,omitempty" bson:"street2,omitempty"`
	Street3     string `json:"street3,omitempty" bson:"street3,omitempty"`
	City        string `json:"city,omitempty" bson:"city,omitempty"`
	State       string `json:"state,omitempty" bson:"state,omitempty"`
	PostalCode  string `json:"postalCode,omitempty" bson:"postalCode,omitempty"`
	Country     string `json:"country,omitempty" bson:"country,omitempty"`
	Phone       string `json:"phone,omitempty" bson:"phone,omitempty"`
	Residential bool   `json:"residential,omitempty" bson:"residential,omitempty"`
}

type Weight struct {
	Value float64 `json:"value,omitempty" bson:"value,omitempty"`
	Units string  `json:"units,omitempty" bson:"units,omitempty"`
}

type Option struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Value string `json:"value,omitempty" bson:"value,omitempty"`
}

type ShipStationItem struct {
	LineItemKey       string   `json:"lineItemKey,omitempty" bson:"lineItemKey,omitempty"`
	Sku               string   `json:"sku,omitempty" bson:"sku,omitempty"`
	Name              string   `json:"name,omitempty" bson:"name,omitempty"`
	ImageURL          string   `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
	Weight            *Weight  `json:"weight,omitempty" bson:"weight,omitempty"`
	Quantity          int      `json:"quantity,omitempty" bson:"quantity,omitempty"`
	UnitPrice         float64  `json:"unitPrice,omitempty" bson:"unitPrice,omitempty"`
	TaxAmount         float64  `json:"taxAmount,omitempty" bson:"taxAmount,omitempty"`
	ShippingAmount    float64  `json:"shippingAmount,omitempty" bson:"shippingAmout,omitempty"`
	WarehouseLocation string   `json:"warehouseLocation,omitempty" bson:"warehouseLocation,omitempty"`
	Options           []Option `json:"options,omitempty" bson:"options,omitempty"`
	ProductID         int      `json:"productId,omitempty" bson:"productId,omitempty"`
	FulfillmentSku    string   `json:"fulfillmentSku,omitempty" bson:"fulfillmentSku,omitempty"`
	Adjustment        bool     `json:"adjustment,omitempty" bson:"adjustment,omitempty"`
	Upc               string   `json:"upc,omitempty" bson:"upc,omitempty"`
}

type AdvancedOptions struct {
	Source       string `json:"source" bson:"source"`
	StoreId      int    `json:"storeId,omitempty" bson:"storeId,omitempty"`
	Warehouseid  int    `json:"warehouseId,omitempty" bson:"warehouseId,omitempty"`
	CustomField1 string `json:"customField1,omitempty" bson:"customField1,omitempty"`
	CustomField2 string `json:"customField2,omitempty" bson:"customField2,omitempty"`
	CustomField3 string `json:"customField3,omitempty" bson:"customField3,omitempty"`
}

type ShipStationOrder struct {
	OrderId          int64             `json:"orderId,omitempty" bson:"orderId,omitempty"`
	OrderNumber      string            `json:"orderNumber,omitempty" bson:"orderNumber,omitempty"`
	OrderKey         string            `json:"orderKey,omitempty" bson:"orderKey,omitempty"`
	OrderDate        string            `json:"orderDate,omitempty" bson:"orderDate,omitempty"`
	CreateDate       string            `json:"createDate,omitempty" bson:"createDate,omitempty"`
	ModifyDate       string            `json:"modifyDate,omitempty" bson:"modifyDate,omitempty"`
	PaymentDate      string            `json:"paymentDate,omitempty" bson:"paymentDate,omitempty"`
	ShipByDate       string            `json:"shipByDate,omitempty" bson:"shipByDate,omitempty"`
	OrderStatus      string            `json:"orderStatus,omitempty" bson:"orderStatus,omitempty"`
	CustomerId       string            `json:"customerId,omitempty" bson:"customerId,omitempty"`
	CustomerUsername string            `json:"customerUsername,omitempty" bson:"customerUsername,omitempty"`
	CustomerEmail    string            `json:"customerEmail,omitempty" bson:"customerEmail,omitempty"`
	BillTo           *Address          `json:"billTo,omitempty" bson:"billTo,omitempty"`
	ShipTo           *Address          `json:"shipTo,omitempty" bson:"shipTo,omitempty"`
	Items            []ShipStationItem `json:"items,omitempty" bson:"items,omitempty"`
	OrderTotal       float64           `json:"orderTotal,omitempty" bson:"orderTotal,omitempty"`
	AmountPaid       float64           `json:"amountPaid,omitempty" bson:"amountPaid,omitempty"`
	TaxAmount        float64           `json:"taxAmount,omitempty" bson:"taxAmount,omitempty"`
	ShippingAmount   float64           `json:"shippingAmount,omitempty" bson:"shippingAmount,omitempty"`
	CustomerNotes    string            `json:"cutomerNotes,omitempty" bson:"customerNotes,omitempty"`
	InternalNotes    string            `json:"internalNotes,omitempty" bson:"internalNotes,omitempty"`
	Gift             bool              `json:"gift,omitempty" bson:"gift,omitempty"`
	GiftMessage      string            `json:"giftMessage,omitempty" bson:"giftMessage,omitempty"`
	PaymentMethod    string            `json:"paymentMethod,omitempty" bson:"paymentMethod,omitempty"`
	CarrierCode      string            `json:"carrierCode,omitempty" bson:"carrierCode,omitempty"`
	ServiceCode      string            `json:"ServiceCode,omitempty" bson:"ServiceCode,omitempty"`
	PackageCode      string            `json:"packageCode,omitempty" bson:"packageCode,omitempty"`
	ShipDate         string            `json:"shipDate,omitempty" bson:"shipDate,omitempty"`
	AdvancedOptions  AdvancedOptions   `json:"advancedOptions,omitempty" bson:"advanceOptions,omitempty"`
	TagIds           []int             `json:"tagIds,omitempty" bson:"tagsIds,omitempty"`
}

type ShipStationOrderList struct {
	Orders []ShipStationOrder `json:"orders" bson:"orders,omitempty"`
	Total  int                `json:"total" bson:"total,omitempty"`
	Page   int                `json:"page" bson:"page,omitempty"`
	Pages  int                `json:"pages" bson:"pages,omitempty"`
}

type ShipStationResponse struct {
	HasErrors bool `json:"hasErrors,omitempty" bson:"hasErrors,omitempty"`
	Results   []struct {
		OrderID      int    `json:"orderId,omitempty" bson:"orderId,omitempty"`
		OrderNumber  string `json:"orderNumber,omitempty" bson:"orderNumber,omitempty"`
		OrderKey     string `json:"orderKey,omitempty" bson:"orderKey,omitempty"`
		Success      bool   `json:"success,omitempty" bson:"success,omitempty"`
		ErrorMessage string `json:"errorMessage,omitempty" bson:"errorMessage,omitempty"`
	} `json:"results,omitempty" bson:"results,omitempty"`
}

type ShipStationMarkAsShipped struct {
	OrderID            int64  `json:"orderId,omitempty" bson:"orderId,omitempty"`
	CarrierCode        string `json:"carrierCode,omitempty" bson:"carrierCode,omitempty"`
	ShipDate           string `json:"shipDate,omitempty" bson:"shipDate,omitempty"`
	TrackingNumber     string `json:"trackingNumber,omitempty" bson:"trackingNumber,omitempty"`
	NotifyCustomer     bool   `json:"notifyCustomer,omitempty" bson:"notifyCustomer,omitempty"`
	NotifySalesChannel bool   `json:"notifySalesChannel,omitempty" bson:"notifySalesChannel,omitempty"`
}

// Shipments

type ShipmentList struct {
	Shipments []Shipment `json:"shipments"  bson:"shipments,omitempty"`
	Total     int        `json:"total"  bson:"total,omitempty"`
	Page      int        `json:"page"  bson:"page,omitempty"`
	Pages     int        `json:"pages"  bson:"pages,omitempty"`
}

type Shipment struct {
	ShipmentID          int              `json:"shipmentId,omitempty" bson:"shipmentId,omitempty"`
	OrderID             int              `json:"orderId,omitempty" bson:"orderId,omitempty"`
	OrderKey            string           `json:"orderKey,omitempty"  bson:"orderKey,omitempty"`
	UserID              string           `json:"userId,omitempty"  bson:"userId,omitempty"`
	OrderNumber         string           `json:"orderNumber,omitempty" bson:"orderNumber,omitempty"`
	CreateDate          string           `json:"createDate,omitempty" bson:"createDate,omitempty"`
	ShipDate            string           `json:"shipDate,omitempty" bson:"shipDate,omitempty"`
	ModifyDate          string           `json:"modifyDate,omitempty" bson:"modifyDate,omitempty"`
	PaymentDate         string           `json:"paymentDate,omitempty" bson:"paymentDate,omitempty"`
	ShipmentCost        float64          `json:"shipmentCost,omitempty"  bson:"shipmentCost,omitempty"`
	InsuranceCost       float64          `json:"insuranceCost,omitempty" bson:"insuranceCost,omitempty"`
	TrackingNumber      string           `json:"trackingNumber,omitempty" bson:"trackingNumber,omitempty"`
	IsReturnLabel       bool             `json:"isReturnLabel,omitempty" bson:"isReturnLabel,omitempty"`
	BatchNumber         string           `json:"batchNumber,omitempty" bson:"batchNumber,omitempty"`
	CarrierCode         string           `json:"carrierCode,omitempty" bson:"carrierCode,omitempty"`
	ServiceCode         string           `json:"serviceCode,omitempty" bson:"serviceCode,omitempty"`
	PackageCode         string           `json:"packageCode,omitempty" bson:"packageCode,omitempty"`
	Confirmation        string           `json:"confirmation,omitempty" bson:"confirmation,omitempty"`
	WarehouseID         int              `json:"warehouseId,omitempty" bson:"warehouseId,omitempty"`
	Voided              bool             `json:"voided,omitempty" bson:"voided,omitempty"`
	VoidDate            string           `json:"voidDate,omitempty" bson:"voidDate,omitempty"`
	MarketplaceNotified bool             `json:"marketplaceNotified,omitempty" bson:"marketplaceNotified,omitempty"`
	NotifyErrorMessage  interface{}      `json:"notifyErrorMessage,omitempty"  bson:"notifyErrorMessage,omitempty"`
	ShipTo              ShipTo           `json:"shipTo,omitempty" bson:"shipTo,omitempty"`
	Weight              Weight           `json:"weight,omitempty" bson:"weight,omitempty"`
	Dimensions          interface{}      `json:"dimensions,omitempty" bson:"dimensions,omitempty"`
	InsuranceOptions    InsuranceOptions `json:"insuranceOptions,omitempty" bson:"insuranceOptions,omitempty"`
	AdvancedOptions     AdvancedOptions  `json:"advancedOptions,omitempty" bson:"advanceOptions,omitempty"`
	ShipmentItems       []ShipmentItems  `json:"shipmentItems,omitempty" bson:"shipmentItems,omitempty"`
	LabelData           interface{}      `json:"labelData,omitempty" bson:"labelData,omitempty"`
	FormData            interface{}      `json:"formData,omitempty" bson:"formData,omitempty"`
}
type ShipTo struct {
	Name        string `json:"name" bson:"name,omitempty"`
	Company     string `json:"company" bson:"company,omitempty"`
	Street1     string `json:"street1" bson:"street1,omitempty"`
	Street2     string `json:"street2" bson:"street2,omitempty"`
	Street3     string `json:"street3" bson:"street3,omitempty"`
	City        string `json:"city" bson:"city,omitempty"`
	State       string `json:"state" bson:"state,omitempty"`
	PostalCode  string `json:"postalCode" bson:"postalCode,omitempty"`
	Country     string `json:"country" bson:"country,omitempty"`
	Phone       string `json:"phone" bson:"phone,omitempty"`
	Residential bool   `json:"residential" bson:"residential,omitempty"`
}
type InsuranceOptions struct {
	Provider       string  `json:"provider" bson:"provider,omitempty"`
	InsureShipment bool    `json:"insureShipment" bson:"insureShipment,omitempty"`
	InsuredValue   float64 `json:"insuredValue" bson:"insureValue,omitempty"`
}
type ShipmentItems struct {
	OrderItemID       int    `json:"orderItemId" bson:"orderItemId,omitempty"`
	LineItemKey       string `json:"lineItemKey" bson:"lineItemKey,omitempty"`
	Sku               string `json:"sku" bson:"sku,omitempty"`
	Name              string `json:"name" bson:"name,omitempty"`
	ImageURL          string `json:"imageUrl" bson:"imageUrl,omitempty"`
	Weight            Weight `json:"weight" bson:"weight,omitempty"`
	Quantity          int    `json:"quantity" bson:"quantity,omitempty"`
	UnitPrice         int    `json:"unitPrice" bson:"unitPrice,omitempty"`
	WarehouseLocation string `json:"warehouseLocation" bson:"warehouseLocation,omitempty"`
	Options           Option `json:"options" bson:"options,omitempty"`
	ProductID         int    `json:"productId" bson:"productId,omitempty"`
	FulfillmentSku    string `json:"fulfillmentSku" bson:"fulfillmentSku,omitempty"`
}

type ShipstationHoldItem struct {
	OrderID       int64  `json:"orderId" bson:"orderId,omitempty"`
	HoldUntilDate string `json:"holdUntilDate" bson:"holdUntilDate,omitempty"`
}

type ShipstationMarkShippedResponse struct {
	OrderID     int64  `json:"orderId" bson:"orderId,omitempty"`
	OrderNumber string `json:"orderNumber" bson:"orderNumber,omitempty"`
}

type ShipstationHoldResponse struct {
	Success  bool   `json:"success" bson:"success,omitempty"`
	Messages string `json:"message" bson:"message,omitempty"`
}
