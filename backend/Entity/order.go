package entity

type Order struct {
	OrderId  int           `json:"orderid"`
	UserId   int           `json:"userid"`
	Articles []ArticleLine `json:"articles"`
}
