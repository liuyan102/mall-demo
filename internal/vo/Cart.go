package vo

import (
	"github.com/spf13/viper"
	"mall-demo/internal/model"
)

type CartResponse struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"userID"`
	ProductID     uint   `json:"productID"`
	Num           uint   `json:"num"`
	MaxNum        uint   `json:"maxNum"`
	ImgPath       string `json:"imgPath"`
	Check         bool   `json:"check"`
	DiscountPrice string `json:"discountPrice"`
	BossID        uint   `json:"bossID"`
	BossName      string `json:"bossName"`
}

func BuildCartResponse(cart *model.Cart) CartResponse {
	return CartResponse{
		ID:            cart.ID,
		UserID:        cart.UserID,
		ProductID:     cart.ProductID,
		Num:           cart.Num,
		MaxNum:        cart.MaxNum,
		ImgPath:       viper.GetString("path.hostPath") + viper.GetString("path.productPath") + cart.Product.ImgPath,
		Check:         cart.Check,
		DiscountPrice: cart.Product.DiscountPrice,
		BossID:        cart.Boss.ID,
		BossName:      cart.Boss.NickName,
	}
}

func BuildCartResponseList(cartList []model.Cart) (cartResponseList []CartResponse) {
	for _, cart := range cartList {
		cartResponse := BuildCartResponse(&cart)
		cartResponseList = append(cartResponseList, cartResponse)
	}
	return cartResponseList
}
