package model

import "gorm.io/gorm"

// Address 地址信息
type Address struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`                   // 用户ID
	Name    string `gorm:"type:varchar(20) not null"`  // 用户名称
	Phone   string `gorm:"type:varchar(11) not null"`  // 用户手机号
	Address string `gorm:"type:varchar(100) not null"` // 用户地址
}
