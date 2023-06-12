package model

import "gorm.io/gorm"

// Admin 管理员信息
type Admin struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20) not null"`  // 管理员用户名
	Password string `gorm:"type:varchar(20) not null"`  // 管理员密码(加密)
	Avatar   string `gorm:"type:varchar(255) not null"` // 管理员头像
}
