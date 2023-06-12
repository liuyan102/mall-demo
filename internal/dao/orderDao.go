package dao

import (
	"errors"
	"gorm.io/gorm"
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDaoWithDB() *OrderDao {
	return &OrderDao{initialize.GetDB()}
}

func (*OrderDao) CreateOrder(order *model.Order) error {
	db := initialize.GetDB()
	return db.Model(&model.Order{}).Create(&order).Error
}

func (*OrderDao) GetOrderByID(orderID uint, userID uint) (order model.Order, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Order{}).Where("id=? and user_id=?", orderID, userID).
		Preload("Product").Preload("Boss").Preload("Address").First(&order).Error
	return order, err
}

func (*OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (orderList []model.Order, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Order{}).Where(condition).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).
		Preload("Product").Preload("Boss").Preload("Address").Find(&orderList).Error
	if err != nil {
		return nil, err
	}
	return orderList, nil
}

func (*OrderDao) DeleteOrderByID(orderID uint, userID uint) error {
	db := initialize.GetDB()
	result := db.Model(&model.Order{}).Where("id=? and user_id=?", orderID, userID).Delete(&model.Order{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to delete record")
	}
	return nil
}

func (*OrderDao) UpdateOrderTypeByID(userID uint, orderID uint, orderType uint) error {
	db := initialize.GetDB()
	result := db.Model(&model.Order{}).Where("id=? and user_id=?", orderID, userID).Update("type", orderType)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update record")
	}
	return nil
}
