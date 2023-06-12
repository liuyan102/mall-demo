package dto

import "mall-demo/internal/model"

type OrderRequest struct {
	UserID    uint    `json:"userID"`
	BossID    uint    `json:"bossID"`
	ProductID uint    `json:"productID"`
	AddressID uint    `json:"addressID"`
	Num       int     `json:"num"`
	OrderNum  uint64  `json:"orderNum"`
	Type      uint    `json:"type"`
	Money     float64 `json:"money"`
	model.BasePage
}
