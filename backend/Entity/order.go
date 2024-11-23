package entity

type Order struct {
	ID     int `gorm:"primaryKey"`
	UserId int
}
