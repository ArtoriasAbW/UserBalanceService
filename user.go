package balance

type User struct {
	Id      uint64 `json:"id" uri:"id" binding:"required"`
	Balance uint32 `json:"balance"`
}