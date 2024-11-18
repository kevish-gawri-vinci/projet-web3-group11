package entity

type Basket struct {
	UserId     int          `json:"userid"`
	Articles   []BasketItem `json:"articles"`
	TotalPrice float32      `json:"totalprice"`
}
