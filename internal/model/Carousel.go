package model

import "gorm.io/gorm"

// Carousel 轮播图
type Carousel struct {
	gorm.Model
	ImgPath   string `gorm:"type:varchar(255) not null"` // 图片路径
	ProductID uint   `gorm:"not null"`                   // 商品ID
}
