package entity

type Basket struct {
	UserId     int           `json:"userid"`
	Articles   []ArticleLine `json:"articles"`
	TotalPrice float32       `json:"totalprice"`
}
