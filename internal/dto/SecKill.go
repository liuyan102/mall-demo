package dto

// SecKillProductRequest 创建秒杀商品请求数据
type SecKillProductRequest struct {
	Name         string `json:"name"`          // 商品名
	CategoryID   uint   `json:"categoryID"`    // 商品分类
	Title        string `json:"title"`         // 商品标题
	Info         string `json:"info"`          // 商品信息
	ImgPath      string `json:"imgPath"`       // 商品图片路径
	Price        string `json:"price"`         // 商品价格
	SecKillPrice string `json:"discountPrice"` // 商品折扣价格
	Num          int    `json:"num"`           // 商品数量
}

// SecKill 秒杀商品
type SecKill struct {
	SecKillProductID uint    `json:"secKillProductID"`
	BossID           uint    `json:"bossID"`
	AddressID        uint    `json:"addressID"`
	Num              int     `json:"num"`
	Money            float64 `json:"money"`
}
