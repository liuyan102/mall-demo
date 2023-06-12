package dao

import (
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"
)

type NoticeDao struct {
}

// GetNoticeByID 获取通知
func (*NoticeDao) GetNoticeByID(id uint) (notice *model.Notice, err error) {
	db := initialize.GetDB()
	err = db.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	if err != nil {
		return nil, err
	}
	return notice, nil
}
