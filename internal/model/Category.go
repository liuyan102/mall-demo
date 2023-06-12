package model

import "gorm.io/gorm"

// Category 商品分类
type Category struct {
	gorm.Model
	CategoryName string `gorm:"type:varchar(50) not null"` // 商品分类名
}
