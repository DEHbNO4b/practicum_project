package balance

type Balance struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
	User_id   int `json:"-"`
}
