package entity

type OrderLine struct {
	OrderId   int `json:"orderid" gorm:"primaryKey"`
	ArticleId int `json:"articleid" gorm:"primaryKey"`
	Quantity  int `json:"quantity"`
}
