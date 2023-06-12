package vo

import (
	"github.com/spf13/viper"
	"mall-demo/internal/model"
)

type FavoriteResponse struct {
	UserID        uint   `json:"userID"`
	ProductID     uint   `json:"productID"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"categoryID"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"imgPath"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discountPrice"`
	BossID        uint   `json:"bossID"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"onSale"`
}

func BuildFavoriteResponse(favorite *model.Favorite) FavoriteResponse {
	return FavoriteResponse{
		UserID:        favorite.UserID,
		ProductID:     favorite.ProductID,
		Name:          favorite.Product.Name,
		CategoryID:    favorite.Product.CategoryID,
		Title:         favorite.Product.Title,
		Info:          favorite.Product.Info,
		ImgPath:       viper.GetString("path.hostPath") + viper.GetString("path.productPath") + favorite.Product.ImgPath,
		Price:         favorite.Product.Price,
		DiscountPrice: favorite.Product.DiscountPrice,
		BossID:        favorite.BossID,
		Num:           favorite.Product.Num,
		OnSale:        favorite.Product.OnSale,
	}
}

func BuildFavoriteResponseList(favoriteList []model.Favorite) (favoriteResponseList []FavoriteResponse) {
	for _, favorite := range favoriteList {
		favoriteResponse := BuildFavoriteResponse(&favorite)
		favoriteResponseList = append(favoriteResponseList, favoriteResponse)
	}
	return favoriteResponseList
}
