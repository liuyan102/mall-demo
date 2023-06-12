package dao

import (
	"errors"
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"
)

type CartDao struct {
}

func (*CartDao) CreateCart(cart *model.Cart) error {
	db := initialize.GetDB()
	return db.Model(&model.Cart{}).Create(&cart).Error
}

func (*CartDao) GetCartByID(cartID uint, userID uint) (cart *model.Cart, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Cart{}).Where("id=? and user_id=?", cartID, userID).
		Preload("Product").Preload("Boss").First(&cart).Error
	return cart, err
}

func (*CartDao) LisCartByUserID(userID uint) (cartList []model.Cart, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Cart{}).Where("user_id=?", userID).
		Preload("Product").Preload("Boss").Find(&cartList).Error
	return cartList, err
}

func (*CartDao) UpdateCartNumByID(cartID uint, userID uint, num uint) error {
	db := initialize.GetDB()
	result := db.Model(&model.Cart{}).Where("id=? and user_id=?", cartID, userID).Update("num", num)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update record")
	}
	return nil
}

func (*CartDao) DeleteCartByID(cartID uint, userID uint) error {
	db := initialize.GetDB()
	result := db.Model(&model.Cart{}).Where("id=? and user_id=?", cartID, userID).Delete(&model.Cart{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to delete record")
	}
	return nil
}
