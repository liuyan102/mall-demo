package dao

import (
	"errors"
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"
)

type ProductDao struct {
}

// CreateProduct 创建商品
func (*ProductDao) CreateProduct(product *model.Product) error {
	db := initialize.GetDB()
	return db.Create(&product).Error
}

// CountProductByCondition 分类查询商品数
func (*ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Product{}).Where(condition).Count(&total).Error
	return total, err
}

// ListProductByCondition 分类查询商品
func (*ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) (productList []model.Product, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Product{}).Where(condition).Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&productList).Error
	return productList, err
}

// SearchProduct 搜索商品
func (*ProductDao) SearchProduct(info string, page model.BasePage) (productList []model.Product, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Product{}).Where("name like ? or title like ? or info like ?", "%"+info+"%", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&productList).Error
	return productList, err
}

// GetProductByID 通过id获取商品信息
func (*ProductDao) GetProductByID(id uint) (product *model.Product, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Product{}).Where("id = ?", id).First(&product).Error
	return product, err
}

func (*ProductDao) UpdateProductNumByID(productID uint, num int) error {
	db := initialize.GetDB()
	result := db.Model(&model.Product{}).Where("id=?", productID).Update("num", num)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update record")
	}
	return nil
}
