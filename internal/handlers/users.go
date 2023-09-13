package handlers

import "net/http"

type UserRegRepo interface {
	New()
}
type UserRegister struct {
	userRepo UserRegRepo
}

func NewRegister(repo UserRegRepo) *UserRegister {
	r := UserRegister{userRepo: repo}
	return &r
}
func (u *UserRegister) Register(w http.ResponseWriter, r *http.Request) {

}
