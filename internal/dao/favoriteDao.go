package dao

import (
	"errors"
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"
)

type FavoriteDao struct {
}

func (*FavoriteDao) Create(favorite *model.Favorite) error {
	db := initialize.GetDB()
	return db.Model(&model.Favorite{}).Create(&favorite).Error
}

func (*FavoriteDao) Exist(userID uint, productID uint) (exist bool, err error) {
	db := initialize.GetDB()
	var count int64
	err = db.Model(&model.Favorite{}).Where("user_id=? and product_id=?", userID, productID).Count(&count).Error
	if err != nil || count == 0 {
		return false, err
	}
	return true, nil
}

func (*FavoriteDao) List(userID uint) (favoriteList []model.Favorite, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Favorite{}).Where("user_id=?", userID).
		Preload("User").Preload("Product").Preload("Boss").Find(&favoriteList).Error
	return favoriteList, err
}

func (*FavoriteDao) Delete(userID uint, favoriteID uint) error {
	db := initialize.GetDB()
	result := db.Model(&model.Favorite{}).Where("user_id=? and id=?", userID, favoriteID).Delete(&model.Favorite{})
	if result.RowsAffected == 0 {
		return errors.New("failed to delete record")
	}
	return nil
}
