package vo

import (
	"github.com/spf13/viper"
	"mall-demo/internal/model"
)

type OrderResponse struct {
	ID           uint    `json:"id"`
	OrderNum     uint64  `json:"orderNum"`
	UserID       uint    `json:"userID"`
	ProductID    uint    `json:"productID"`
	BossID       uint    `json:"bossID"`
	Num          int     `json:"num"`
	AddressName  string  `json:"addressName"`
	AddressPhone string  `json:"addressPhone"`
	Address      string  `json:"address"`
	Type         uint    `json:"type"`
	ProductName  string  `json:"productName"`
	ImgPath      string  `json:"imgPath"`
	Money        float64 `json:"money"`
}

func BuildOrderResponse(order *model.Order) OrderResponse {
	return OrderResponse{
		ID:           order.ID,
		OrderNum:     order.OrderNum,
		UserID:       order.UserID,
		ProductID:    order.ProductID,
		BossID:       order.BossID,
		Num:          order.Num,
		AddressName:  order.Address.Name,
		AddressPhone: order.Address.Phone,
		Address:      order.Address.Address,
		Type:         order.Type,
		ProductName:  order.Product.Name,
		ImgPath:      viper.GetString("path.hostPath") + viper.GetString("path.productPath") + order.Product.ImgPath,
		Money:        order.Money,
	}
}

func BuildOrderResponseList(orderList []model.Order) (orderResponseList []OrderResponse) {
	for _, order := range orderList {
		orderResponse := BuildOrderResponse(&order)
		orderResponseList = append(orderResponseList, orderResponse)
	}
	return orderResponseList
}
