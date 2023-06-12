package model

import "gorm.io/gorm"

// Order 订单信息
type Order struct {
	gorm.Model
	UserID    uint    `gorm:"not null"` // 用户ID
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `gorm:"not null"` // 商品ID
	BossID    uint    `gorm:"not null"` // 厂家ID
	Boss      User    `gorm:"ForeignKey:BossID"`
	AddressID uint    `gorm:"not null"` // 地址ID
	Address   Address `gorm:"ForeignKey:AddressID"`
	Num       int     `gorm:"not null"` // 订单数
	OrderNum  uint64  `gorm:"not null"` // 订单号
	Type      uint    `gorm:"not null"` // 是否支付 0未支付 1已支付
	Money     float64 `gorm:"not null"` // 金额
}
