package dto

import "mall-demo/internal/model"

type CreateFavoriteRequest struct {
	ProductID  uint `json:"productID"`
	BossID     uint `json:"bossID"`
	FavoriteID uint `json:"favoriteID"`
	model.BasePage
}
