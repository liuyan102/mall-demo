package model

import "gorm.io/gorm"

// Favorite 收藏夹
type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserID"`    // 用户
	UserID    uint    `gorm:"not null"`             // 用户ID
	Product   Product `gorm:"ForeignKey:ProductID"` // 商品
	ProductID uint    `gorm:"not null"`             // 商品ID
	Boss      User    `gorm:"ForeignKey:BossID"`    // 厂家
	BossID    uint    `gorm:"not null"`             // 厂家ID
}
