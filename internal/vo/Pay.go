package vo

import (
	"mall-demo/internal/model"
	"time"
)

type PayResponse struct {
	OrderID   uint    `json:"orderID"`
	OrderNum  uint64  `json:"orderNum"`
	ProductID uint    `json:"productID"`
	BossID    uint    `json:"bossID"`
	BossName  string  `json:"bossName"`
	Num       int     `json:"num"`
	Money     float64 `json:"money"`
	Type      uint    `json:"type"`
	PayTime   string  `json:"payTime"`
}

func BuildPayResponse(order *model.Order) PayResponse {
	return PayResponse{
		OrderID:   order.ID,
		OrderNum:  order.OrderNum,
		ProductID: order.ProductID,
		BossID:    order.BossID,
		BossName:  order.Boss.NickName,
		Num:       order.Num,
		Money:     order.Money * float64(order.Num),
		Type:      order.Type,
		PayTime:   time.Now().String(),
	}
}
