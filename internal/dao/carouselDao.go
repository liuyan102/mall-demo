package dao

import (
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"
)

type CarouselDao struct {
}

func (*CarouselDao) List() (carouselList []model.Carousel, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Carousel{}).Find(&carouselList).Error
	if err != nil {
		return nil, err
	}
	return carouselList, nil
}
