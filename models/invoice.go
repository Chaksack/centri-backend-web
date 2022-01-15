package models

import "time"

type Invoice struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time
	CategoryRefer int      `json:"category_id"`
	Category      Category `gorm:"foreignKey:CategoryRefer"`
	ProductRefer  int      `json:"product_id"`
	Product       Product  `gorm:"foreignKey:ProductRefer"`
	UserRefer     int      `json:"user_id"`
	User          User     `json:"foreignKey:UserRefer"`
}
