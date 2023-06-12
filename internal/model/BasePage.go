package model

// BasePage 分页信息
type BasePage struct {
	PageNum  int `json:"pageNum"`  // 页数
	PageSize int `json:"pageSize"` // 每页容量
}
