package models

type Source struct {
	OrderNumber         string `csv:"Order Number"`
	OrderStatus         string `csv:"Order Status"`
	OrderDate           string `csv:"Order Date"`
	CustomerNote        string `csv:"Customer Note"`
	FirstNameBilling    string `csv:"First Name (Billing)"`
	LastNameBilling     string `csv:"Last Name (Billing)"`
	CompanyBilling      string `csv:"Company (Billing)"`
	AddressBilling      string `csv:"Address 1&2 (Billing)"`
	CityBilling         string `csv:"City (Billing)"`
	StateCodeBilling    string `csv:"State Code (Billing)"`
	PostcodeBilling     string `csv:"Postcode (Billing)"`
	CountryCodeBilling  string `csv:"Country Code (Billing)"`
	EmailBilling        string `csv:"Email (Billing)"`
	PhoneBilling        string `csv:"Phone (Billing)"`
	FirstNameShipping   string `csv:"First Name (Shipping)"`
	LastNameShipping    string `csv:"Last Name (Shipping)"`
	AddressShipping     string `csv:"Address 1&2 (Shipping)"`
	CityShipping        string `csv:"City (Shipping)"`
	StateCodeShipping   string `csv:"State Code (Shipping)"`
	PostcodeShipping    string `csv:"Postcode (Shipping)"`
	CountryCodeShipping string `csv:"Country Code (Shipping)"`
	PaymentMethodTitle  string `csv:"Payment Method Title"`
	CartDiscountAmount  string `csv:"Cart Discount Amount"`
	OrderSubtotalAmount string `csv:"Order Subtotal Amount"`
	ShippingMethodTitle string `csv:"Shipping Method Title"`
	OrderShippingAmount string `csv:"Order Shipping Amount"`
	OrderRefundAmount   string `csv:"Order Refund Amount"`
	OrderTotalAmount    string `csv:"Order Total Amount"`
	OrderTotalTaxAmount string `csv:"Order Total Tax Amount"`
	SKU                 string `csv:"SKU"`
	ItemNumber          string `csv:"Item #"`
	ItemName            string `csv:"Item Name"`
	Quantity            string `csv:"Quantity"`
	ItemCost            string `csv:"Item Cost"`
	CouponCode          string `csv:"Coupon Code"`
	DiscountAmount      string `csv:"Discount Amount"`
	DiscountAmountTax   string `csv:"Discount Amount Tax"`
}
