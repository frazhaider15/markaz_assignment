package models

import "time"

type Product struct {
	Id          int64      `json:"id" gorm:"column:id;primary_key"`
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description;default:null"`
	Category    string     `json:"category" gorm:"column:category;default:null"`
	Image       string     `json:"image" gorm:"column:image;default:null"`
	IsActive    bool       `json:"is_active" gorm:"column:is_active"`
	Price       float64    `json:"price,omitempty" gorm:"column:price;default:null"`
	SKU         string     `json:"sku" gorm:"column:sku;default:null"`
	Currency    string     `json:"currency" gorm:"currency"`
	MerchantID  int64      `json:"merchant_user_id" gorm:"column:user_id"`
	CreatedAt   time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at;default:null"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at" gorm:"column:deleted_at;default:null"`
}
