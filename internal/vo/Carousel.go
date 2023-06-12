package vo

import "mall-demo/internal/model"

// CarouselResponse 轮播图响应数据
type CarouselResponse struct {
	ID        uint   `json:"id"`
	ImgPath   string `json:"imgPath"`   // 图片路径
	ProductID uint   `json:"productID"` // 商品ID
}

func BuildCarouselResponse(carousel *model.Carousel) CarouselResponse {
	return CarouselResponse{
		ID:        carousel.ID,
		ImgPath:   carousel.ImgPath,
		ProductID: carousel.ProductID,
	}
}

func BuildCarouselListResponse(carouselList []model.Carousel) (carouselResponseList []CarouselResponse) {
	for _, carousel := range carouselList {
		carouselResponse := BuildCarouselResponse(&carousel)
		carouselResponseList = append(carouselResponseList, carouselResponse)
	}
	return carouselResponseList
}
