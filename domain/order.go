package domain

import "time"

type Order struct {
	ID         string `gorm:"primaryKey" `
	UserID     string `gorm:"index" `
	User       User
	ProductID  string
	Product    Product
	Quantity   int       `gorm:"not null"`
	TotalPrice int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
}
