package model

import "gorm.io/gorm"

// User 用户信息
type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20)"`  // 用户名
	Email    string `gorm:"type:varchar(50)"`  // 邮箱
	Password string `gorm:"type:varchar(255)"` // 密码
	NickName string `gorm:"type:varchar(20)"`  // 昵称
	Avatar   string `gorm:"type:varchar(255)"` // 头像
	Status   string `gorm:"type:varchar(255)"` // 状态
	Money    string `gorm:"type:varchar(255)"` // 余额
}

const (
	Active = "active"
)
