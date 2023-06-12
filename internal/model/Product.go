package model

import (
	"context"
	"mall-demo/internal/cache"
	"strconv"

	"gorm.io/gorm"
)

// Product 商品信息
type Product struct {
	gorm.Model
	Name          string `gorm:"type:varchar(255) not null"` // 商品名
	CategoryID    uint   `gorm:"not null"`                   // 商品分类
	Title         string `gorm:"type:varchar(255) not null"` // 商品标题
	Info          string `gorm:"type:varchar(255) not null"` // 商品信息
	ImgPath       string `gorm:"type:varchar(255) not null"` // 商品图片路径
	Price         string `gorm:"type:varchar(255) not null"` // 商品价格
	DiscountPrice string `gorm:"type:varchar(255) not null"` // 商品折扣价格
	OnSale        bool   `gorm:"default:false"`              // 是否打折
	Num           int    `gorm:"not null"`                   // 商品数量
	BossID        uint   `gorm:"not null"`                   // 厂家ID
	BossName      string `gorm:"type:varchar(255) not null"` // 厂家名称
	BossAvatar    string `gorm:"type:varchar(255) not null"` // 厂家头像
}

// View 商品点击量
func (product *Product) View(ctx context.Context) uint64 {
	countStr, _ := cache.RedisDB.Get(ctx, cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 增加商品点击量
func (product *Product) AddView(ctx context.Context) {
	// 自增
	cache.RedisDB.Incr(ctx, cache.ProductViewKey(product.ID))
	// 对有序集合中的指定成员member的分数上增量increment
	cache.RedisDB.ZIncrBy(ctx, cache.Rank, 1, strconv.Itoa(int(product.ID)))
}
