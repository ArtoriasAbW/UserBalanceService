package balance


type User struct {
	Id       uint64 `json:"id" uri:"id" binding:"required"`
	Balance  float64     `json:"balance"`
	Currency string `query:"currency"`
}