package dao

import (
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"

	"gorm.io/gorm"
)

type SecKillOrderDao struct {
	*gorm.DB
}

func NewSecKillOrderDaoWithDB() *SecKillOrderDao {
	return &SecKillOrderDao{initialize.GetDB()}
}

func (*SecKillOrderDao) CreateSecKillOrder(order *model.SecKillOrder) error {
	return initialize.DB.Model(&model.SecKillOrder{}).Create(&order).Error
}
