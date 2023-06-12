package dao

import (
	"errors"
	"fmt"
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"
)

type UserDao struct {
}

// IsUserNameExist 判断用户名是否存在
func (*UserDao) IsUserNameExist(username string) bool {
	var count int64
	db := initialize.GetDB()
	db.Model(&model.User{}).Where("user_name=?", username).Count(&count)
	fmt.Println(count)
	if count != 0 {
		return true
	}
	return false
}

// CreateUser 创建用户
func (*UserDao) CreateUser(user *model.User) error {
	db := initialize.GetDB()
	err := db.Create(&user).Error
	if err != nil {
		return errors.New("user create err")
	}
	return nil
}

// GetUserInfoByUserName 根据用户名获取用户信息
func (*UserDao) GetUserInfoByUserName(username string) (user *model.User, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.User{}).Where("user_name=?", username).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetUserInfoByID 根据用户id获取用户名信息
func (*UserDao) GetUserInfoByID(id uint) (user *model.User, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.User{}).Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// UpdateUser 修改用户信息
func (*UserDao) UpdateUser(user *model.User) error {
	db := initialize.GetDB()
	result := db.Model(&model.User{}).Where("id=?", user.ID).Updates(&user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update record")
	}
	return nil
}

// UpdateUserMoneyByID 修改用户金额
func (*UserDao) UpdateUserMoneyByID(userID uint, money string) error {
	db := initialize.GetDB()
	result := db.Model(&model.User{}).Where("id=?", userID).Update("money", money)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update record")
	}
	return nil

}
