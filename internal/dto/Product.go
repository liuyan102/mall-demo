package dto

import "mall-demo/internal/model"

// CreateProductRequest 创建商品请求数据
type CreateProductRequest struct {
	ID            uint   `json:"id"`            // 商品ID
	Name          string `json:"name"`          // 商品名
	CategoryID    uint   `json:"categoryID"`    // 商品分类
	Title         string `json:"title"`         // 商品标题
	Info          string `json:"info"`          // 商品信息
	ImgPath       string `json:"imgPath"`       // 商品图片路径
	Price         string `json:"price"`         // 商品价格
	DiscountPrice string `json:"discountPrice"` // 商品折扣价格
	OnSale        bool   `json:"onSale"`        // 是否打折
	Num           int    `json:"num"`           // 商品数量
}

// ListProductRequest 展示商品请求
type ListProductRequest struct {
	CategoryID uint `json:"categoryID"`
	model.BasePage
}

// SearchProductRequest 搜索商品请求
type SearchProductRequest struct {
	Info string `json:"info"`
	model.BasePage
}
