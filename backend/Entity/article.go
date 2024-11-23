package entity

type Article struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float32
	ImgUrl      string
}
