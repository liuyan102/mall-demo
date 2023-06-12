package vo

import (
	"mall-demo/internal/model"

	"github.com/spf13/viper"
)

type SecKillProductResponse struct {
	ID           uint   `json:"id"`           // 商品ID
	Name         string `json:"name"`         // 商品名
	CategoryID   uint   `json:"categoryID"`   // 商品分类
	Title        string `json:"title"`        // 商品标题
	Info         string `json:"info"`         // 商品信息
	ImgPath      string `json:"imgPath"`      // 商品图片路径
	Price        string `json:"price"`        // 商品价格
	SecKillPrice string `json:"secKillPrice"` // 秒杀价格
	Num          int    `json:"num"`          // 商品数量
	BossID       uint   `json:"bossID"`
	BossName     string `json:"bossName"`
	BossAvatar   string `json:"bossAvatar"`
}

func BuildSecKillProductResponse(product *model.SecKillProduct) SecKillProductResponse {
	return SecKillProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		CategoryID: product.CategoryID,
		Title:      product.Title,
		Info:       product.Info,
		ImgPath:    viper.GetString("path.hostPath") + viper.GetString("path.productPath") + product.ImgPath,
		Price:      product.Price,
		Num:        product.Num,
		BossID:     product.BossID,
		BossName:   product.BossName,
		BossAvatar: viper.GetString("path.hostPath") + viper.GetString("path.avatarPath") + product.BossAvatar,
	}
}

func BuildSecKillProductResponseList(secKillProductList []model.SecKillProduct) (secKillProductResponseList []SecKillProductResponse) {
	for _, secKillProduct := range secKillProductList {
		secKillProductResponse := BuildSecKillProductResponse(&secKillProduct)
		secKillProductResponseList = append(secKillProductResponseList, secKillProductResponse)
	}
	return secKillProductResponseList
}
