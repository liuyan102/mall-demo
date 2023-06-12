package vo

import (
	"github.com/spf13/viper"
	"mall-demo/internal/model"
)

type ProductImgResponse struct {
	ProductID uint   `json:"productID"`
	ImgPath   string `json:"imgPath"`
}

func BuildProductImgResponse(productImg *model.ProductImg) ProductImgResponse {
	return ProductImgResponse{
		ProductID: productImg.ProductID,
		ImgPath:   viper.GetString("path.hostPath") + viper.GetString("path.productPath") + productImg.ImgPath,
	}
}

func BuildProductImgListResponse(productImgs []model.ProductImg) (productImgResponseList []ProductImgResponse) {
	for _, productImg := range productImgs {
		productImgResponse := BuildProductImgResponse(&productImg)
		productImgResponseList = append(productImgResponseList, productImgResponse)
	}
	return productImgResponseList
}
