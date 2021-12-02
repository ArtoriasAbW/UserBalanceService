package balance

type Operation struct {
	UserId      uint64 `json:"id" binding:"required"`
	Value float64 `json:"value" binding:"required"`
}

type TransferOperation struct {
	SenderId   uint64 `json:"sender_id" binding:"required"`
	ReceiverId uint64 `json:"receiver_id" binding:"required"`
	Value float64 `json:"value" binding:"required"`
}