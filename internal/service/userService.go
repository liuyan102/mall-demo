package service

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"mall-demo/internal/dao"
	"mall-demo/internal/dto"
	"mall-demo/internal/model"
	"mall-demo/internal/pkg/e"
	"mall-demo/internal/pkg/res"
	"mall-demo/internal/pkg/util"
	"mall-demo/internal/vo"
	"mime/multipart"
	"strconv"
	"strings"
)

type UserService struct {
}

// UserRegister 用户注册
func (*UserService) UserRegister(ctx *gin.Context, request dto.UserRegisterRequest) res.Response {
	// 验证密钥
	if request.Key == "" || len(request.Key) != 16 {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "密钥长度不足",
		}
	}

	// 验证用户名是否存在
	var userDao dao.UserDao
	exist := userDao.IsUserNameExist(request.UserName)
	if exist {
		return res.Response{
			Code: e.InvalidParam,
			Data: nil,
			Msg:  "用户名已存在",
		}
	}
	// 密码加密
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "密码加密错误",
		}
	}
	// 初始金额10000,金额加密
	money, err2 := util.AesEncrypt(strconv.Itoa(10000), request.Key)
	if err2 != nil {
		util.Loggers.Errorln(err2)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "金额加密错误",
		}
	}
	newUser := &model.User{
		UserName: request.UserName,
		NickName: request.NickName,
		Password: string(password),
		Status:   model.Active,
		Avatar:   "avatar.jpg",
		Money:    money,
	}
	// 创建用户
	err3 := userDao.CreateUser(newUser)
	if err3 != nil {
		util.Loggers.Errorln(err3)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "注册失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "注册成功",
	}
}

// UserLogin 用户登录
func (*UserService) UserLogin(ctx *gin.Context, request dto.UserLoginRequest) res.Response {
	var userDao dao.UserDao
	// 查找用户名是否存在
	user, err := userDao.GetUserInfoByUserName(request.UserName)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "用户名不存在",
		}
	}

	// 判断密码是否正确
	err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err2 != nil {
		util.Loggers.Errorln(err2)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "密码错误",
		}
	}

	// 发放token
	token, err3 := util.GenerateToken(user.ID, user.UserName)
	if err3 != nil {
		util.Loggers.Errorln(err3)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "token签发失败",
		}
	}

	// 登录成功
	return res.Response{
		Code: e.Success,
		Data: vo.UserInfoWithTokenResponse{
			UserInfo: vo.BuildUserInfoResponse(user),
			Token:    token,
		},
		Msg: "登录成功",
	}
}

// NickNameUpdate 修改用户昵称
func (*UserService) NickNameUpdate(ctx *gin.Context, newNickName string) res.Response {
	var userDao dao.UserDao
	user, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	// nickName为空或nickName长度>20
	if len(newNickName) == 0 || len(newNickName) > 20 {
		return res.Response{
			Code: e.InvalidParam,
			Data: nil,
			Msg:  "参数错误",
		}
	}
	// 类型断言
	newUser := user.(*model.User)
	newUser.NickName = newNickName
	err := userDao.UpdateUser(newUser)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "修改失败",
		}
	}
	ctx.Set("user", newUser)
	return res.Response{
		Code: e.Success,
		Data: vo.BuildUserInfoResponse(newUser),
		Msg:  "修改成功",
	}

}

// AvatarUpload 头像上传
func (*UserService) AvatarUpload(ctx *gin.Context, file multipart.File, fileSize int64) res.Response {
	var userDao dao.UserDao
	user, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	// 上传文件是否为空
	if file == nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "上传文件为空",
		}
	}
	newUser := user.(*model.User)
	// 保存图片到本地
	avatarPath, err := util.UploadAvatar(file, newUser.ID, newUser.UserName)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "头像上传错误",
		}
	}
	newUser.Avatar = avatarPath
	err2 := userDao.UpdateUser(newUser)
	if err2 != nil {
		util.Loggers.Errorln(err2)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "头像上传失败",
		}
	}
	ctx.Set("user", newUser)
	return res.Response{
		Code: e.Success,
		Data: vo.BuildUserInfoResponse(newUser),
		Msg:  "头像上传成功",
	}
}

// SendEmail 发送邮箱
func (*UserService) SendEmail(ctx *gin.Context, request dto.SendEmailRequest) res.Response {
	user, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	newUser := user.(*model.User)
	// 签发email token
	emailToken, err := util.GenerateEmailToken(newUser.ID, request.Email, request.Password, request.OperationType)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "签发token失败",
		}
	}
	var noticeDao dao.NoticeDao
	notice, err := noticeDao.GetNoticeByID(request.OperationType)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "获取通知失败",
		}
	}

	mailAddr := viper.GetString("email.validEmail") + emailToken
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", mailAddr, -1)
	// 新建mail
	mail := gomail.NewMessage()
	// 设置邮件头
	mail.SetHeader("From", viper.GetString("email.smtpEmail"))
	mail.SetHeader("To", request.Email)
	mail.SetHeader("Subject", "mall-demo")
	mail.SetBody("text/html", mailText)
	// 新建
	dialer := gomail.NewDialer(viper.GetString("email.smtpHost"), 465, viper.GetString("email.smtpEmail"), viper.GetString("email.smtpPass"))
	// 关闭SSL协议认证
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err2 := dialer.DialAndSend(mail)
	if err2 != nil {
		util.Loggers.Errorln(err2)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "邮件发送失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "邮件发送成功",
	}
}

// ValidEmail 验证邮箱
func (*UserService) ValidEmail(ctx *gin.Context) res.Response {
	// 验证token
	emailToken := ctx.GetHeader("Authorization")
	emailToken = emailToken[7:]
	emailClaims, err := util.ParseEmailToken(emailToken)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.UnAuthorized,
			Data: nil,
			Msg:  "权限不足",
		}
	}
	user, exist := ctx.Get("user")
	if !exist {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	newUser := user.(*model.User)
	if emailClaims.OperationType == 1 {
		// 绑定邮箱
		newUser.Email = emailClaims.Email
	} else if emailClaims.OperationType == 2 {
		// 解绑邮箱
		newUser.Email = ""
	} else if emailClaims.OperationType == 3 {
		// 修改密码
		password, err2 := bcrypt.GenerateFromPassword([]byte(emailClaims.Password), bcrypt.DefaultCost)
		if err2 != nil {
			util.Loggers.Errorln(err2)
			return res.Response{
				Code: e.Error,
				Data: nil,
				Msg:  "加密错误",
			}
		}
		newUser.Password = string(password)
	}
	var userDao dao.UserDao
	err3 := userDao.UpdateUser(newUser)
	if err3 != nil {
		util.Loggers.Errorln(err3)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "修改失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: newUser,
		Msg:  "修改成功",
	}
}

// ShowMoney 显示金额
func (*UserService) ShowMoney(ctx *gin.Context, key string) res.Response {
	user, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	money, err := util.AesDecrypt(user.(*model.User).Money, key)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "密钥错误",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: money,
		Msg:  "显示成功",
	}
}
