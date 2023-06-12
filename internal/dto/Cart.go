package dto

type CartRequest struct {
	ID        uint `json:"id"`
	BossID    uint `json:"bossID"`
	ProductID uint `json:"productID"`
	Num       uint `json:"num"`
}
