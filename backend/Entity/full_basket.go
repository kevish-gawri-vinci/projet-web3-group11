/*
*Entity for response to /basket/get
 */

package entity

type FullBasket struct {
	UserId     int            `json:"userid"`
	Articles   []BasketDetail `json:"articles"`
	TotalPrice float32        `json:"totalprice"`
}

type BasketDetail struct {
	ArticleDetail Article `json:"article"`
	Quantity      int     `json:"quantity"`
	LinePrice     float32 `json:"lineprice"`
}
