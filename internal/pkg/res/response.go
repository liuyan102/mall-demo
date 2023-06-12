package res

// Response 响应数据结构
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"` // 空接口类型，空接口没有方法，所有类型都默认实现了空接口，可以接收任意类型的变量
	Msg  string      `json:"msg"`
}

// DataList 数据切片
type DataList struct {
	Items interface{} `json:"item"`
	Total int64       `json:"total"`
}

func BuildDataList(items interface{}, total int64) DataList {
	return DataList{
		Items: items,
		Total: total,
	}
}
