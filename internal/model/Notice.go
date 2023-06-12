package model

import "gorm.io/gorm"

// Notice 通知
type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"` // 通知信息
}
