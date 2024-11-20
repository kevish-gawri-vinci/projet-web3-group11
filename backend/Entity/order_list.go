package entity

//For /order/getall ==> Returning a list of orders
type OrderListLine struct {
	OrderId       int     `json:"orderid"`
	TotalPrice    float32 `json:"totalprice"`
	TotalQuantity int     `json:"totalquantity"`
}
