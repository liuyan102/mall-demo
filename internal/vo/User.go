package vo

import (
	"github.com/spf13/viper"
	"mall-demo/internal/model"
)

// UserInfoResponse 用户登录返回用户信息
type UserInfoResponse struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Money    string `json:"money"`
}

type UserInfoWithTokenResponse struct {
	UserInfo UserInfoResponse `json:"userInfo"`
	Token    string           `json:"token"`
}

func BuildUserInfoResponse(user *model.User) UserInfoResponse {
	return UserInfoResponse{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Avatar:   viper.GetString("path.hostPath") + viper.GetString("path.avatarPath") + user.Avatar,
		Money:    user.Money,
	}
}
