package vo

import (
	"context"
	"mall-demo/internal/model"

	"github.com/spf13/viper"
)

type ProductResponse struct {
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
	View          uint64 `json:"view"`          // 浏览量
	BossID        uint   `json:"bossID"`
	BossName      string `json:"bossName"`
	BossAvatar    string `json:"bossAvatar"`
}

func BuildProductResponse(product *model.Product) ProductResponse {
	return ProductResponse{
		ID:            product.ID,
		Name:          product.Name,
		CategoryID:    product.CategoryID,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       viper.GetString("path.hostPath") + viper.GetString("path.productPath") + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		OnSale:        product.OnSale,
		Num:           product.Num,
		View:          product.View(context.Background()),
		BossID:        product.BossID,
		BossName:      product.BossName,
		BossAvatar:    viper.GetString("path.hostPath") + viper.GetString("path.avatarPath") + product.BossAvatar,
	}
}

func BuildProductResponseList(productList []model.Product) (productResponseList []ProductResponse) {
	for _, product := range productList {
		productResponse := BuildProductResponse(&product)
		productResponseList = append(productResponseList, productResponse)
	}
	return productResponseList
}
