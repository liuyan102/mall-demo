package model

import (
	"gorm.io/gorm"
)

// SecKillProduct 秒杀商品信息
type SecKillProduct struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255) not null"` // 秒杀商品名
	CategoryID   uint   `gorm:"not null"`                   // 商品分类
	Title        string `gorm:"type:varchar(255) not null"` // 商品标题
	Info         string `gorm:"type:varchar(255) not null"` // 商品信息
	ImgPath      string `gorm:"type:varchar(255) not null"` // 商品图片
	Price        string `gorm:"type:varchar(255) not null"` // 商品原价
	SecKillPrice string `gorm:"type:varchar(255) not null"` // 秒杀价格
	Num          int    `gorm:"not null"`                   // 商品数量
	//SecKillTime  time.Time `gorm:"not null"`                   // 秒杀时间
	BossID     uint   `gorm:"not null"`                   // 厂家ID
	BossName   string `gorm:"type:varchar(255) not null"` // 厂家名称
	BossAvatar string `gorm:"type:varchar(255) not null"` // 厂家头像
}
