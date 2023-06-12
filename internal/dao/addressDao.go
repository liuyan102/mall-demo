package dao

import (
	"errors"
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"
)

type AddressDao struct {
}

func (*AddressDao) CreateAddress(address *model.Address) error {
	db := initialize.GetDB()
	return db.Model(&model.Address{}).Create(&address).Error
}

func (*AddressDao) GetAddressByID(id uint, userID uint) (address *model.Address, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Address{}).Where("id=? and user_id=?", id, userID).First(&address).Error
	return address, err
}

func (*AddressDao) ListAddressByUserID(userID uint) (addressList []model.Address, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Address{}).Where("user_id=?", userID).Find(&addressList).Error
	return addressList, err
}

func (*AddressDao) UpdateAddressByID(addressID uint, userID uint, address *model.Address) error {
	db := initialize.GetDB()
	result := db.Model(&model.Address{}).Where("id=? and user_id=?", addressID, userID).Updates(&address)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update record")
	}
	return nil
}

func (*AddressDao) DeleteAddressByID(addressID uint, userID uint) error {
	db := initialize.GetDB()
	result := db.Model(&model.Address{}).Where("id=? and user_id=?", addressID, userID).Delete(&model.Address{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to delete record")
	}
	return nil
}
