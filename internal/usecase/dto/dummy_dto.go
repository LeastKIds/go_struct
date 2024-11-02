package dto

import "github.com/LeastKIds/go_struct/internal/domain/vo"

type Dummy struct {
	ID   vo.ID   `json:"id"`
	Name vo.Name `json:"name"`
	Age  vo.Age  `json:"age"`
}
