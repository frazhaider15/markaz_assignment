package models

import "time"

type User struct {
	Id                    int64     `json:"id" gorm:"column:id;primary_key"`
	Name                  string    `json:"name" gorm:"column:name"`
	Email                 string    `json:"email" gorm:"column:email"`
	PhoneNumber           string    `json:"phone_number" gorm:"column:phone_number"`
	Password              string    `json:"password" gorm:"column:password"`
	ProfilePhotoPath      string    `json:"profile_photo_path" gorm:"column:profile_photo_path"`
	DateOfBirth           string    `json:"date_of_birth" gorm:"column:date_of_birth"`
	Gender                string    `json:"gender" gorm:"column:gender"`
	Status                string    `json:"status" gorm:"column:status"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	DeletedAt             time.Time `json:"deleted_at"`
}

func (b *User) TableName() string {
	return "users"
}
