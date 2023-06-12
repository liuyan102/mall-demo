package dto

type PayRequest struct {
	OrderID  uint   `json:"orderID"`
	OrderNum uint64 `json:"orderNum"`
	UserKey  string `json:"userKey"`
	BossKey  string `json:"bossKey"`
}
