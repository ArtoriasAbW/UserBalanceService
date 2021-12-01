package balance

type User struct {
	Id      int `json:"id" uri:"id" binding:"required"`
	Balance int `json:"balance"`
}