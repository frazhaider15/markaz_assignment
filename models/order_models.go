package models

import "time"

type Order struct {
	Id          int64
	OrderNumber string
	UserId      int64
	Status      OrderStatus
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type OrderItem struct {
	Id          int64
	OrderNumber string
	UserId      int64
	ProductId   int64
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}
