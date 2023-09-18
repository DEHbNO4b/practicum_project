package user

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Balance  int    `json:"balance"`
}
