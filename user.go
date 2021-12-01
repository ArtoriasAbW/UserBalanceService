package balance

type User struct {
	Id int `json:"-"`
	Balance int `json:"balance"`
}