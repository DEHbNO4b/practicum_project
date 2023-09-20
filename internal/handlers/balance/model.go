package balance

type Balance struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
	UserID    int `json:"-"`
}
