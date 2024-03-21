package models

type Destination struct {
	OrderNumber                string `csv:"order_number"`
	SellerSKUCode              string `csv:"seller_sku_code"`
	Quantity                   string `csv:"quantity"`
	SalesChannel               string `csv:"sales_channel"`
	OrderCreatedAt             string `csv:"order_created_at"`
	ShippingMethod             string `csv:"shipping_method"`
	ShippingAddress            string `csv:"shipping_address"`
	ShippingCity               string `csv:"shipping_city"`
	ShippingCountry            string `csv:"shipping_country"`
	ShippingRegion             string `csv:"shipping_region"`
	ShippingPhone              string `csv:"shipping_phone"`
	ShippingCountryCallingCode string `csv:"shipping_country_calling_code"`
	ShippingCountryCode        string `csv:"shipping_country_code"`
	CustomerName               string `csv:"customer_name"`
	CustomerPhone              string `csv:"customer_phone"`
	CustomerEmail              string `csv:"customer_email"`
	CustomerCountryCallingCode string `csv:"customer_country_calling_code"`
	CustomerCountryCode        string `csv:"customer_country_code"`
	Currency                   string `csv:"currency"`
	Subtotal                   string `csv:"subtotal"`
	ShippingTax                string `csv:"shipping_tax"`
	SubTotalTaxInclusive       string `csv:"sub_total_tax_inclusive"`
	SubTotalDiscountInclusive  string `csv:"sub_total_discount_inclusive"`
	ShippingTaxInclusive       string `csv:"shipping_tax_inclusive"`
	ShippingDiscountInclusive  string `csv:"shipping_discount_inclusive"`
	ShippingPrice              string `csv:"shipping_price"`
	ShippingRefund             string `csv:"shipping_refund"`
	Tax                        string `csv:"tax"`
	TaxPercent                 string `csv:"tax_percent"`
	Discount                   string `csv:"discount"`
	Total                      string `csv:"total"`
	TotalPaid                  string `csv:"total_paid"`
	TotalDue                   string `csv:"total_due"`
	TotalRefunded              string `csv:"total_refunded"`
	PaymentMethod              string `csv:"payment_method"`
	LoyaltyDiscount            string `csv:"loyalty_discount"`
	WalletAmount               string `csv:"wallet_amount"`
	DisplayPrice               string `csv:"display_price"`
	SellingPrice               string `csv:"selling_price"`
	ItemTaxPercent             string `csv:"item_tax_percent"`
	ItemTax                    string `csv:"item_tax"`
	UnitPrice                  string `csv:"unit_price"`
	ItemTotal                  string `csv:"item_total"`
	TaxInclusive               string `csv:"tax_inclusive"`
	ItemDiscount               string `csv:"item_discount"`
	DeliveryDate               string `csv:"delivery_date"`
	SlotStartTime              string `csv:"slot_start_time"`
	SlotEndTime                string `csv:"slot_end_time"`
	OrderAlias                 string `csv:"order_alias"`
	PartialPicking             string `csv:"partial_picking"`
	ItemWeight                 string `csv:"item_weight"`
	TotalWeight                string `csv:"total_weight"`
}
