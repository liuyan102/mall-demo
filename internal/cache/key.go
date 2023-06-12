package cache

import (
	"fmt"
	"strconv"
)

const (
	Rank = "rank"
)

// ProductViewKey 商品点击量
func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
