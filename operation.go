package balance

type Operation struct {
	UserId      uint64 `json:"id" binding:"required"`
	Value uint32 `json:"value" binding:"required"`
}