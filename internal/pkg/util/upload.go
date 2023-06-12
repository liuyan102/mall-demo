package util

import (
	"github.com/spf13/viper"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

// UploadAvatar 上传头像
func UploadAvatar(file multipart.File, userID uint, userName string) (filePath string, err error) {
	id := strconv.Itoa(int(userID))
	// 每个用户一个文件夹
	basePath := "." + viper.GetString("path.avatarPath") + "user" + id + "/"
	// 文件夹是否存在
	if !IsDirExist(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, bytes, os.ModePerm)
	if err != nil {
		return "", err
	}
	return "user" + id + "/" + userName + ".jpg", nil
}

// UploadProduct 上传商品图片
func UploadProduct(file multipart.File, userID uint, productName string) (filePath string, err error) {
	id := strconv.Itoa(int(userID))
	// 每个用户一个文件夹
	basePath := "." + viper.GetString("path.productPath") + "boss" + id + "/"
	// 文件夹是否存在
	if !IsDirExist(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + productName + ".jpg"
	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(productPath, bytes, os.ModePerm)
	if err != nil {
		return "", err
	}
	return "boss" + id + "/" + productName + ".jpg", nil
}

// IsDirExist 判断文件夹是否存在
func IsDirExist(fileAddr string) bool {
	stat, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// CreateDir 创建文件夹
func CreateDir(dirName string) bool {
	// mkdirAll 创建多级目录，第一个参数是文件夹路径，第二个参数是文件夹权限
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}
