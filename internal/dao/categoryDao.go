package dao

import (
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"
)

type CategoryDao struct {
}

// ListCategory 获取商品分类
func (*CategoryDao) ListCategory() (categoryList []model.Category, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Category{}).Find(&categoryList).Error
	return categoryList, err
}
