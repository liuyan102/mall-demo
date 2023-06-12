package dao

import (
	"errors"
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"

	"gorm.io/gorm"
)

type SecKillProductDao struct {
	*gorm.DB
}

func NewSecKillProductDaoWithDB() *SecKillProductDao {
	return &SecKillProductDao{initialize.GetDB()}
}

func (*SecKillProductDao) CreateSecKillProduct(product *model.SecKillProduct) error {
	return initialize.DB.Model(&model.SecKillProduct{}).Create(&product).Error
}

func (*SecKillProductDao) GetSecKillProductNum(secKillProductID uint) (num int, err error) {
	err = initialize.DB.Model(&model.SecKillProduct{}).Select("num").Where("id=?", secKillProductID).First(&num).Error
	return num, err
}

func (*SecKillProductDao) UpdateSecKillProductNum(secKillProductID uint, num int) error {
	result := initialize.DB.Model(&model.SecKillProduct{}).Where("id=?", secKillProductID).Update("num", num)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update record")
	}
	return nil
}
