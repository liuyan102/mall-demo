package dao

import (
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"
)

type ProductImgDao struct {
}

// CreateProductImg 创建商品图片
func (*ProductImgDao) CreateProductImg(productImg *model.ProductImg) error {
	db := initialize.GetDB()
	return db.Model(&model.ProductImg{}).Create(&productImg).Error
}

// GetProductImgByID 通过id获取商品图片
func (*ProductImgDao) GetProductImgByID(id uint) (productImgList []model.ProductImg, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.ProductImg{}).Where("product_id=?", id).Find(&productImgList).Error
	return productImgList, err
}
