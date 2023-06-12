package model

import "gorm.io/gorm"

// Cart 购物车信息
type Cart struct {
	gorm.Model
	UserID    uint    `gorm:"not null"` // 用户ID
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `gorm:"not null"` // 商品ID
	Boss      User    `gorm:"ForeignKey:BossID"`
	BossID    uint    `gorm:"not null"` // 厂家ID
	Num       uint    `gorm:"not null"` // 商品数目
	MaxNum    uint    `gorm:"not null"` // 商品最大数目
	Check     bool    // 是否支付
}
