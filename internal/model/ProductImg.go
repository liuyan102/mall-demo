package model

import "gorm.io/gorm"

// ProductImg 商品图片
type ProductImg struct {
	gorm.Model
	ProductID uint   `gorm:"not null"`
	ImgPath   string `gorm:"type:varchar(255) not null"`
}
